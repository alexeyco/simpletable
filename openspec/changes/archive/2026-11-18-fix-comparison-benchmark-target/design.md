# Design: fix the comparison benchmark target

Motivation, the acceptance gate, decisions of record, assumptions, and
rejected risks are in `proposal.md`; this file records how the change is
implemented and verified. All facts below were re-verified read-only at
f21f69b during this planning run.

## Context

- `Makefile:6` defines `BENCH_DIR := _example/03-benchmarks-with-others`.
  The `bench-compare` recipe (`Makefile:51`) is
  `cd $(BENCH_DIR) && go test -bench . -benchmem` — no package argument, so
  `go test` targets only the module's root package, which contains exactly
  one file, `main.go` (`package main`), and no test files.
- Reproduced defect: `GOPROXY=off make bench-compare →
  ?   github.com/alexeyco/simpletable/_example/03-benchmarks-with-others  [no test files]`,
  exit status 0, zero benchmarks executed.
- The three benchmarks live in the module's `bench` package:
  `_example/03-benchmarks-with-others/bench/simpletable_test.go:9`
  (`BenchmarkSimpletable`), `bench/termtables_test.go:9`
  (`BenchmarkTermtables`), `bench/uitable_test.go:9` (`BenchmarkUITable`).
  The module's `go.mod:20` has `replace github.com/alexeyco/simpletable => ../..`.
- The lint/fmt targets already operate inside `$(BENCH_DIR)`: `Makefile:27`
  spells out `./...` (`cd $(BENCH_DIR) && golangci-lint run
  --config=../../.golangci.yml ./...`), while `Makefile:32` (`golangci-lint
  fmt`) and `Makefile:37` (`golangci-lint fmt --diff`) pass no path
  argument and golangci-lint defaults to `./...`. `Makefile:51` is the only
  recipe that omits the package argument, so it selects the module's root
  package, which has no test files.
- Lock mechanism verified: from the module root,
  `go test -list 'Benchmark.*' ./bench` prints exactly the three names then
  `ok .../bench` (~0.6 s warm, `GOPROXY=off`, deps already in the module
  cache). `go test -list 'Benchmark.*' .` on the root package prints
  `[no test files]` and STILL exits 0 — so the assertion must inspect
  names, never the exit status alone.
- Documentation surfaces: `AGENTS.md:17` claims `make bench-compare` "runs
  the third-party comparison benchmarks"; `Makefile:50` carries the help
  text; the harness README has heading `## Run benchmarks` at line 66 and
  the manual `cd .../bench && go test -bench=. -benchmem` lines at 67-69;
  root `README.md:100` links to the anchor `#run-benchmarks`;
  `openspec/specs/package-readme/spec.md:63-66` requires the benchmark
  section to link to the benchmark instructions.
- Harness README `## Installation` section (`:4-8`) predates Go modules:
  `$ go get github.com/apcera/termtables` (`:6`) and
  `$ go get github.com/gosuri/uitable` (`:7`), while both comparators are
  pinned as module requirements (`go.mod:7-8`, recorded at
  `go.sum:1-2,7-8`) and the module declares `go 1.23` (`go.mod:3`);
  `go get` stopped installing packages in module mode as of Go 1.17+.
  Offline resolution re-proven read-only in this run:
  `GOPROXY=off go test -list 'Benchmark.*' ./bench` → the three names plus
  `ok .../bench`, exit 0; `GOPROXY=off go mod download` → exit 0.
- Baseline capability `openspec/specs/repository-quality-gates/spec.md`:
  H1 title (:1), `## Purpose` (:3-5), five requirements (descriptions
  118/130/331/167/177 chars, all under the 500-char strict limit); none
  mentions `make bench-compare`, `make bench`, or `make check`.
- CI: `.github/workflows/quality-gates.yml:27,30,54,57` invoke only
  `make test`, `make build`, `make lint`, `make fmt-check`;
  `coveralls.yml:27` invokes `make cover`. No workflow invokes `make check`,
  `make bench`, or `make bench-compare`.
- Lint surface: `.golangci.yml` enables 34 Go-only linters plus `gofmt`/
  `goimports`; there is NO Makefile linter (no checkmake/mbake) and no
  markdown lint anywhere (no markdownlint/prettier/vale/remark config, no
  actionlint — the claim in
  `openspec/changes/archive/2026-11-09-add-go-quality-gates/design.md:7` is
  stale). `misspell`/`godot`/`dupword` read Go files only, so `.md` and
  `Makefile` edits cannot trip `make lint` or `make fmt-check`.
- `make check` prerequisites today: `test build lint fmt-check
  openspec-validate` (`Makefile:61-62`); `bench`/`bench-compare` are
  excluded from it.

## Goals / Non-Goals

Goals:

- `make bench-compare` really executes the three comparison benchmarks.
- A name-based lock (`make bench-compare-check`) fails when any of the
  three names disappears and is enforced by `make check`.
- `AGENTS.md:17`, the `Makefile:50` help text, and the harness README
  `## Run benchmarks` section consistently point at the `bench` package,
  and the README `## Installation` section describes a working,
  module-aware way to obtain the pinned comparators.
- `make help` and `openspec validate --all --strict --no-interactive` stay
  green.

Non-Goals:

- No CI workflow change; the lock is local-only (see Risks).
- No Go source change, no new dependency, no new tooling.
- No edit to root `README.md`, `CONTRIBUTING.md`, `make cover`, the harness
  `Show`/`Results` sections, or the benchmarks, fixtures, comparators, and
  timing values themselves.
- No `table-output` or `package-readme` spec delta.

## Decisions

1. Explicit `./bench` over `./...`. The alternative
   `cd $(BENCH_DIR) && go test -bench . -benchmem ./...` (consistent with
   `Makefile:27,32,37`) also fixes the target, but it always prints
   `[no test files]` for the root `main` package and would not fail if the
   `bench` package moved or vanished; the user chose explicit `./bench`.
   The `-run '^$$'` guard matches the `bench` target (`Makefile:47`) so no
   unit test in the package can run under the benchmark invocation. In the
   Makefile the literal bytes are `'^$$'`: make expands `$$` to a single
   `$`, so the shell receives `'^$'` and Go receives the empty-test
   pattern.
2. Name-based lock, not exit status. `-list` exits 0 even when no benchmark
   matches (verified on the root package), so an exit-status check cannot
   detect the defect class. The lock captures the output of
   `go test -list 'Benchmark.*' ./bench` and fails, naming the culprit,
   unless each of the three concrete names appears as a whole line.
   Consequence accepted by the user: all three existing benchmarks are
   locked names, so renaming one requires a spec change — set equality
   was declined (Decision 8), which makes this presence lock the
   strictest surviving option.
3. Rejected lock mechanism — asserting the output does NOT contain
   `[no test files]`: invalid, because with `./...` the benchmarks DO run
   and that line is still printed for the root `main` package; and it would
   falsely pass if a trivial `main_test.go` were added at the module root.
4. Makefile target over a Go test in the root module. A Go test would have
   to shell out into a second module (separate `go.mod` with the `replace`
   directive) and add cross-module test plumbing plus a new gate surface;
   the Makefile form needs no new tooling and no dependency and sits beside
   the target it protects. Chosen recipe bytes (the loop is ONE recipe line
   continued with backslashes; each `$` is doubled for make):

   ```make
   .PHONY: bench-compare-check
   bench-compare-check: ## Fail unless module 03's bench package lists all three comparison benchmarks
   	@cd $(BENCH_DIR) && listed=$$(go test -list 'Benchmark.*' ./bench); \
   	for b in BenchmarkSimpletable BenchmarkTermtables BenchmarkUITable; do \
   		echo "$$listed" | grep -qx "$$b" || { echo "bench-compare-check: missing benchmark $$b" >&2; exit 1; }; \
   	done
   ```

   A build failure of the `bench` package also fails the lock: `listed`
   then holds error text without the three lines. `grep -qx` matches whole
   lines, so the trailing `ok .../bench` line cannot satisfy a name.
5. Lock wiring: `bench-compare-check` becomes a prerequisite of `check`
   (`Makefile:62`, inserted after `fmt-check`), and the `check` help
   comment gains the benchmark lock. Enforcement is deliberately
   local-only; the CI trade-off is a recorded rejected risk in
   `proposal.md`.
6. Anchor preservation: the harness README heading `## Run benchmarks`
   (line 66) stays byte-identical because root `README.md:100` links to
   `#run-benchmarks` and `openspec/specs/package-readme/spec.md:63-66`
   requires that link to reach the benchmark instructions. Only lines
   66-70 may change; the manual `cd .../bench && go test -bench=. -benchmem`
   form stays pointing at the `bench` package, and the section gains
   `make bench-compare` as the repo-root invocation. The `cd` line's
   pre-module `$GOPATH/src/...` prefix (`:68`) becomes repo-relative, so
   after this change that form remains only in the excluded `## Show`
   section (`:12`); `Show`/`Results` and root `README.md` are untouched
   and `Installation` is modernized by Decision 9.
7. Help-text alignment: the `Makefile:50` description names the `bench`
   package, e.g.
   `bench-compare: ## Run the third-party comparison benchmarks (module 03, bench package)`;
   `AGENTS.md:17` keeps its sentence shape and names the `bench` package of
   `_example/03-benchmarks-with-others`. Cosmetic consequence accepted: the
   19-character target name overruns the `%-18s` awk column
   (`Makefile:12`), shifting only its own `make help` row (recorded
   rejected risk).
8. Presence lock, no `[no test files]` clause. The scenario
   `bench-compare runs the three named benchmarks` asserts only the
   presence of the three names followed by `PASS`; an earlier draft also
   required the output to contain no `[no test files]` line. That clause
   is dropped (user decision): with the explicit `./bench` package
   argument a single-package run cannot print the line, and if the `bench`
   package lost its test files the three-name lock would already fail —
   the clause adds nothing while superficially resembling the
   absence-based mechanism rejected in Decision 3. The lock asserts
   presence, not set equality: an additional fourth benchmark does not
   fail it (set equality was considered and declined); only one of the
   three locked names disappearing does.
9. Installation modernization (user decision; previously out of scope).
   The harness README `## Installation` section
   (`_example/03-benchmarks-with-others/README.md:4-8`) tells readers to
   `go get github.com/apcera/termtables` and `go get
   github.com/gosuri/uitable`. Both lines are wrong in module mode:
   `go get` stopped installing packages as of Go 1.17+, the module
   declares `go 1.23` (`go.mod:3`), and both comparators are already
   pinned as module requirements (`go.mod:7-8`, recorded at
   `go.sum:1-2,7-8`), so Go fetches them automatically. Replacement
   guidance: the pinned requirements download on demand, with
   `go mod download` run from the module directory as the explicit step.
   Offline resolution proven read-only at f21f69b:
   `GOPROXY=off go test -list 'Benchmark.*' ./bench` → three names,
   exit 0; `GOPROXY=off go mod download` → exit 0. The `## Run benchmarks`
   heading must not change (Decision 6): root `README.md:100` resolves the
   `#run-benchmarks` anchor there and
   `openspec/specs/package-readme/spec.md:59-71` requires the benchmark
   section to link to the benchmark instructions. The `## Show` section
   keeps its pre-module `$GOPATH/src` path (`:12`) as a documented,
   deliberate user exclusion.
10. Sync-time Purpose widening (user decision). The delta adds local
    `make` aggregate behavior to `repository-quality-gates`, whose baseline
    `## Purpose` (`openspec/specs/repository-quality-gates/spec.md:5`) says
    the capability "Defines the Go test and lint checks run by the
    repository's GitHub Actions workflow." At archive/sync that sentence
    SHALL be widened to cover the local `make check` aggregate in addition
    to the workflow — proposed: "Defines the Go test, lint, and benchmark
    checks run by the repository's GitHub Actions workflow and by the local
    `make check` aggregate." — keeping every `### Requirement:` and
    `#### Scenario:` heading verbatim. A Purpose edit is not expressible as
    a delta operation in this repo's convention: all four archived deltas
    carry only `## ADDED`/`## MODIFIED Requirements`, never `## Purpose`
    (verified by grep), so it is a deliberate archive-time edit;
    `AGENTS.md:9` routes wording about existing behavior to the main spec
    directly.

## Files to touch

| Path | Change |
|------|--------|
| `Makefile` | :51 recipe → `cd $(BENCH_DIR) && go test -run '^$$' -bench . -benchmem ./bench`; :50 help text names the `bench` package (Decision 7); new `.PHONY: bench-compare-check` + target with the exact bytes of Decision 4, placed after the `bench-compare` block; :62 `check` gains the `bench-compare-check` prerequisite and its help comment mentions the lock. |
| `AGENTS.md` | :17 only — the `make bench-compare` sentence names the `bench` package of `_example/03-benchmarks-with-others`. |
| `_example/03-benchmarks-with-others/README.md` | :4-8 — `## Installation`: the two `go get` lines become module-aware guidance (comparators pinned at `go.mod:7-8` are fetched automatically; `go mod download` from the module directory is the explicit step), Decision 9. :66-70 — heading `## Run benchmarks` byte-identical; add `make bench-compare` as the repo-root invocation; keep the manual `cd .../bench` lines pointing at the `bench` package with a repo-relative path (Decision 6). |
| `openspec/changes/fix-comparison-benchmark-target/**` | This change's artifacts (written at planning time). |
| `openspec/specs/repository-quality-gates/spec.md` | NOT hand-edited during implementation; updated by syncing/archiving this change, including the `## Purpose` widening of Decision 10. |

Not touched: root `README.md`, any Go source or test, `.github/workflows/**`,
`.golangci.yml`, `CONTRIBUTING.md`, harness code/fixtures/comparators.

Sequencing: task Group 1 (Makefile) lands first — Group 2 documents targets
that must already exist; after Group 1 alone, `make check` is expected green
because no gate reads Markdown or the Makefile itself.

## Risks / Trade-offs

- Local-only enforcement: the lock runs in `make check`, which no workflow
  invokes; a regression surfaces locally, not in CI gates (recorded
  REJECTED-RISK in `proposal.md`).
- Strict name lock: renaming any of the three benchmarks breaks
  `make check` until the spec and the lock are updated together — the
  intended coupling, chosen as the strictest option.
- Help-column cosmetics: the 19-character name shifts only its own
  `make help` row (recorded REJECTED-RISK in `proposal.md`).
- The lock compiles the `bench` package (`-list` builds the test binary);
  on a machine without `termtables`/`uitable` in the module cache and
  without network, `go test -list` fails and the lock fails with it — the
  same precondition `bench-compare` itself has, so the lock adds no new
  environment requirement.
- Markdown/Makefile edits are gate-inert: no enabled linter reads them
  (verified inventory in Context), so Groups 1-2 cannot break `make lint`
  or `make fmt-check`; the residual risk is limited to `make help` output
  and the OpenSpec strict validation, both checked in tasks.
