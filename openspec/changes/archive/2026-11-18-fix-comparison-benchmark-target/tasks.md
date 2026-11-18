# Tasks: fix the comparison benchmark target

Group 1 edits only `Makefile`; Group 2 edits `AGENTS.md:17` and
`_example/03-benchmarks-with-others/README.md:4-8,66-70` and documents
targets Group 1 creates, so Group 1 lands first. After Group 1 alone, `make check`
is expected green: no gate reads Markdown or the Makefile itself (see
`design.md` "Files to touch" and Context). Exact recipe bytes are in
`design.md` Decision 4.

## 1. Makefile — fix bench-compare, add the lock, wire it into check

- [x] 1.1 Replace the `bench-compare` recipe at `Makefile:51` with
  `cd $(BENCH_DIR) && go test -run '^$$' -bench . -benchmem ./bench` (the
  literal bytes in the Makefile are `'^$$'`) and update the help
  description at `Makefile:50` to name the `bench` package of module 03.
  Verify: `make bench-compare` output contains `BenchmarkSimpletable`,
  `BenchmarkTermtables`, `BenchmarkUITable`, and `PASS`; the command
  exits 0; the `make help` row for `bench-compare` names the `bench`
  package.
  Map to scenarios `bench-compare runs the three named benchmarks` and `Documentation surfaces agree on the bench package`
- [x] 1.2 Add `.PHONY: bench-compare-check` and the `bench-compare-check`
  target immediately after the `bench-compare` block, using the exact
  recipe bytes from `design.md` Decision 4: capture
  `go test -list 'Benchmark.*' ./bench` inside `$(BENCH_DIR)` and fail,
  printing `bench-compare-check: missing benchmark <name>` to stderr,
  unless each of `BenchmarkSimpletable`, `BenchmarkTermtables`,
  `BenchmarkUITable` appears as a whole line. Verify:
  `make bench-compare-check` exits 0 on the current tree.
  Map to scenario `bench-compare-check passes on the current tree`
- [x] 1.3 Prove the lock really fails: temporarily change the recipe's
  `-list` pattern from `'Benchmark.*'` to `'BenchmarkSimpletable.*'`, run
  `make bench-compare-check`, and expect a non-zero exit reporting
  `missing benchmark BenchmarkTermtables`; restore the pattern to
  `'Benchmark.*'`, rerun, and expect exit 0. Verify: after the temporary
  `-list` narrowing is restored, `git diff Makefile` shows no leftover
  change from this task's temporary edit (only the changes made by the
  other tasks in this group remain).
  Map to scenario `bench-compare-check fails when a name is missing`
- [x] 1.4 Add `bench-compare-check` to the `check` prerequisites at
  `Makefile:62` (after `fmt-check`) and extend that line's help comment to
  mention the benchmark lock. Verify: `make check` exits 0.
  Map to scenario `make check includes the lock and stays green`
- [x] 1.5 Verify: `make help` exits 0, still lists every existing target
  (`help`, `test`, `build`, `lint`, `fmt`, `fmt-check`, `cover`, `bench`,
  `bench-compare`, `openspec-validate`, `check`, `skills`, `skills-update`),
  and now also lists `bench-compare-check` with its description; only that
  row overruns the `%-18s` column (accepted, see `proposal.md`
  REJECTED-RISK).
  Map to scenario `make help lists the new target`

## 2. Docs — align AGENTS.md and the harness README with the bench package and module mode

- [x] 2.1 Update `AGENTS.md:17` so the `make bench-compare` sentence names
  the `bench` package of `_example/03-benchmarks-with-others`; leave the
  `make bench` guidance in the same line unchanged. Verify:
  `make lint && make fmt-check` exit 0 (no enabled linter reads Markdown).
  Map to scenario `Documentation surfaces agree on the bench package`
- [x] 2.2 Edit only the `## Run benchmarks` section
  (`_example/03-benchmarks-with-others/README.md:66-70`): keep the heading
  byte-identical, add `make bench-compare` as the repo-root invocation, and
  keep the manual `cd .../bench && go test -bench=. -benchmem` lines
  pointing at the `bench` package with a repo-relative path (the pre-module
  `$GOPATH/src` prefix remains only in the excluded `## Show` section,
  `:12`). Do not touch `Show`, `Results`, or root `README.md`;
  `Installation` is task 2.3. Verify:
  `grep -c '^## Run benchmarks$' _example/03-benchmarks-with-others/README.md`
  prints `1`, and `git diff --name-only` lists neither `README.md` nor any
  workflow file.
  Map to scenario `Documentation surfaces agree on the bench package`
- [x] 2.3 Replace the harness README `## Installation` section
  (`_example/03-benchmarks-with-others/README.md:4-8`): drop the
  `$ go get github.com/apcera/termtables` and
  `$ go get github.com/gosuri/uitable` lines and state that both
  comparators are pinned as module requirements in
  `_example/03-benchmarks-with-others/go.mod:7-8` (recorded in its
  `go.sum`; the `go.mod:7-8` line range is an implementer-side citation
  for this task only — the delta's scenario requires the README text to
  say only that the comparators are pinned in the module's `go.mod`, and
  it must not embed line numbers), are fetched automatically from those
  requirements, and that `go mod download` run from the module directory
  is the explicit retrieval step. Do not touch the `## Show` section
  (`:10-14`). Verify:
  `grep -c 'go get' _example/03-benchmarks-with-others/README.md` prints
  `0` (`grep -c` exits 1 when it prints `0` because no line matches:
  exit status 1 alongside the count `0` is the expected passing outcome
  here, and the pass condition is the printed count being `0` — do not
  chain this grep with `&&`), and the read-only
  `cd _example/03-benchmarks-with-others && GOPROXY=off go mod download`
  exits 0 (equivalently, `GOPROXY=off go test -list 'Benchmark.*' ./bench`
  lists the three benchmarks and exits 0).
  Map to scenario `Installation section describes module-aware retrieval`
- [x] 2.4 Land the group's gates on the final tree. Verify: `make check`
  (which now includes `bench-compare-check` and `openspec-validate`) exits
  0, and `openspec validate --all --strict --no-interactive` exits 0. If
  either fails, fix the failure — never weaken or skip a check.
  Map to scenario `make check includes the lock and stays green`
