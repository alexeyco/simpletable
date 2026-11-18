# Fix the bench-compare target to run the comparison benchmarks

## Why now

`make bench-compare` does not run anything. The defect is reproduced
read-only at f21f69b: `GOPROXY=off make bench-compare →
?   github.com/alexeyco/simpletable/_example/03-benchmarks-with-others  [no test files]`
with exit status 0 — zero benchmarks executed. The recipe (Makefile:51)
omits the package argument, so `go test` sees only the module's root
package, which contains just `main.go`; the three benchmarks live in the
module's `bench` package. `AGENTS.md:17` already asserts the target "runs
the third-party comparison benchmarks", so the guide documents behavior
that does not exist, and the v1.2.1 README figures had to be captured with
an ad-hoc equivalent command instead of the documented target.

## What changes

- `Makefile:51`: the `bench-compare` recipe becomes
  `cd $(BENCH_DIR) && go test -run '^$$' -bench . -benchmem ./bench` —
  explicit `./bench`, and the `-run '^$$'` guard aligned with the `bench`
  target (Makefile:47); the literal bytes in the Makefile are `'^$$'`.
- New `.PHONY: bench-compare-check` target built on
  `go test -list 'Benchmark.*' ./bench` inside `$(BENCH_DIR)`: it fails,
  naming the culprit, unless all three concrete benchmark names
  (`BenchmarkSimpletable`, `BenchmarkTermtables`, `BenchmarkUITable`)
  appear; wired as a prerequisite of `check`.
- Doc-surface alignment: the `Makefile:50` help description, `AGENTS.md:17`,
  and the `_example/03-benchmarks-with-others/README.md` `## Run benchmarks`
  section (lines 66-70 only) point consistently at the `bench` package; the
  README names `make bench-compare` as the repo-root invocation.
- The harness README `## Installation` section (lines 4-8) replaces its
  obsolete `go get github.com/apcera/termtables` and
  `go get github.com/gosuri/uitable` lines with module-aware guidance: both
  comparators are already pinned as module requirements in
  `_example/03-benchmarks-with-others/go.mod:7-8` (recorded in its
  `go.sum:1-2,7-8`), so Go fetches them automatically and `go mod download`
  run from the module directory is the explicit retrieval step. `go get`
  stopped installing packages in module mode as of Go 1.17+ and the module
  declares `go 1.23` (`go.mod:3`); offline resolution is proven read-only:
  `GOPROXY=off go test -list 'Benchmark.*' ./bench` from the module root
  lists the three benchmarks, exit 0.
- Acceptance gate, all six observable:
  1. `make bench-compare` output reports the three named benchmarks
     `BenchmarkSimpletable`, `BenchmarkTermtables`, `BenchmarkUITable`
     followed by `PASS`.
  2. `make bench-compare-check` FAILS when any of the three names is
     missing and PASSES now.
  3. `make check` stays green with `bench-compare-check` added as a
     prerequisite.
  4. `make help` still lists all targets; the new target appears.
  5. `openspec validate --all --strict --no-interactive` stays green.
  6. The harness README `## Installation` section contains no `go get`
     line, and `cd _example/03-benchmarks-with-others && GOPROXY=off
     go mod download` exits 0.
  CI workflows are not modified.

Out of scope: running the lock check in CI; `CONTRIBUTING.md` (AGENTS.md:3
already routes contribution guidance there); the harness README's `Show`
section, which keeps its pre-module `$GOPATH/src` path
(`_example/03-benchmarks-with-others/README.md:12`); `make cover`; changing
the benchmarks, fixtures, comparators, or timing values; root `README.md`;
any Go source, dependency, or workflow change; `table-output`/
`package-readme` deltas.

## Impact on capabilities

- `repository-quality-gates`: two ADDED requirements — `Comparison
  benchmark gate` (five scenarios: the fixed target output, the lock's
  pass and fail sides, the `check` aggregate, and the `make help` listing)
  and `Comparison benchmark documentation` (two scenarios: the aligned
  documentation surfaces and the module-aware `Installation` guidance),
  split so each requirement description stays under the 500-character
  strict limit. No existing requirement is modified — none of the five
  mentions `make bench-compare`, `make bench`, or `make check`.
- `table-output`, `package-readme`: no delta. The root README keeps its
  `#run-benchmarks` link, which keeps resolving because the harness README
  heading stays verbatim.
- The baseline spec file `openspec/specs/repository-quality-gates/spec.md`
  is synced by archiving this change, not hand-edited during it.
- SYNC-TIME INSTRUCTION (applied at archive/sync only, not a delta
  operation): the main spec's `## Purpose` sentence
  (`openspec/specs/repository-quality-gates/spec.md:5`, today "Defines the
  Go test and lint checks run by the repository's GitHub Actions
  workflow.") SHALL be widened to cover the local `make check` aggregate in
  addition to the workflow; proposed sentence: "Defines the Go test, lint,
  and benchmark checks run by the repository's GitHub Actions workflow and
  by the local `make check` aggregate." Every `### Requirement:` and
  `#### Scenario:` heading stays verbatim. A Purpose edit is not
  expressible as a delta in this repo's convention — all four archived
  deltas carry only `## ADDED`/`## MODIFIED Requirements`, never
  `## Purpose` (verified by grep) — so it is a deliberate archive-time
  edit, and `AGENTS.md:9` routes wording about existing behavior to the
  main spec directly.

**DECISION:** The delta lives in the EXISTING capability
`repository-quality-gates` as `## ADDED Requirements`; a MODIFIED of `Pull
request and tag lint gate` was rejected — it has 331/500 chars headroom but
is semantically about lint, not benchmark selection.

**DECISION:** The target formula is fixed as
`cd $(BENCH_DIR) && go test -run '^$$' -bench . -benchmem ./bench`
(explicit `./bench`; `-run '^$$'` guard aligned with the `bench` target at
Makefile:47; in the Makefile the literal bytes are `'^$$'`).

**DECISION:** The lock asserts PRESENCE of the three concrete benchmark
names, not set equality: it does not fail when an additional fourth
benchmark appears (set equality was considered and declined). Renaming or
adding a benchmark therefore requires updating the lock only when one of
the three locked names disappears.

**DECISION:** The harness README heading `## Run benchmarks` stays verbatim,
because root `README.md:100` links to the anchor `#run-benchmarks` and
`openspec/specs/package-readme/spec.md:63-66` requires the benchmark section
to link to the benchmark instructions; root `README.md` is NOT edited.

**ASSUMPTION:** local enforcement only (no CI job) is acceptable for this
tooling target, matching how `make bench` is already excluded from CI —
tested by one read-only local run:
`grep -rn "make " .github/workflows/ → quality-gates.yml:27,30,54,57 run
make test/build/lint/fmt-check; coveralls.yml:27 runs make cover`, with no
workflow invoking `make check`, `make bench`, or `make bench-compare`.

**ASSUMPTION:** only the harness README `## Show` section keeps its
pre-module `$GOPATH/src/...` path (verified at
`_example/03-benchmarks-with-others/README.md:12`) as a documented,
deliberate exclusion; the `Installation` and `Run benchmarks` guidance are
aligned in this change.

**REJECTED-RISK:** the check does not run in CI — the chosen lock lives in
`make check`, while CI calls only `make test`, `make build`, `make lint`,
`make fmt-check` (`.github/workflows/quality-gates.yml:27,30,54,57`); a
regression surfaces locally, not in gates.

**REJECTED-RISK:** the 19-character name `bench-compare-check` overruns the
`%-18s` help column in `Makefile:12`, shifting only its own `make help` row
— accepted instead of renaming the target.
