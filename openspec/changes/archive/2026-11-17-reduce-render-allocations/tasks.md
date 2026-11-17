# Tasks: reduce per-render allocations of table output

Group A must complete before any source edit (it freezes the byte-exact
target). Groups B and C both edit `content.go`, `table.go`, and `row.go`, so
B runs before C; D measures the final binary after B+C. See `design.md`
"Files to touch" for exact paths.

## 1. Group A — Freeze current output with golden tests (no source change)

- [x] 1.1 Create `golden_test.go` in the existing external test package
  `simpletable_test`: fixtures render the `benchStyleTable` shape
  (benchmarks_test.go:68) with each of the 7 built-in styles, plus multiline,
  ANSI-colored, CJK-wide, and spanned-cell inputs built via the exported API
  and rendered with `StyleUnicode`; each case compares `Table.String()`
  byte-for-byte against `testdata/<case>.golden`, with an `-update-golden`
  flag to regenerate. Map to scenario `Output stays byte identical`
- [x] 1.2 On the unchanged tree, generate the 11 `testdata/*.golden` files
  (`go test -run TestGolden -update-golden .`), rerun `go test -run TestGolden
  .` with the flag off until it passes, then run `make test` and confirm
  green with zero non-test file diffs — this freezes the byte-exact target
  for Groups B/C. Map to scenario `Output stays byte identical`

## 2. Group B — Tier 1 string-building rewrites

- [x] 2.1 Rewrite `content.line` (content.go:61-85) to build each line into
  one `strings.Builder` (pre-Grown to `len(l) + padding`) and append padding
  from a package-level reusable space run via a write-n-spaces helper,
  replacing `fmt.Sprintf` + `strings.Repeat`; keep the three alignment
  branches' exact padding math, the center split `lft := padding / 2`, and
  the `padding <= 0` early return. Verify: `go test -run TestGolden .` shows
  no golden diff. Map to scenario `Output stays byte identical`
- [x] 2.2 Rewrite `content.lines` (content.go:50-58) to
  `make([]string, len(c.c))` with index assignment. Verify: golden diff
  empty. Map to scenario `Output stays byte identical`
- [x] 2.3 Rewrite `tblRow.toStringSlice` (row.go:26-46): pre-size `l` and
  `ret`, hoist the `" "+style.Cell+" "` separator and the `" "+row+" "` wrap
  out of the per-line `fmt.Sprintf` calls into concatenation. Verify: golden
  diff empty. Map to scenario `Output stays byte identical`
- [x] 2.4 Rewrite `Table.borderLeftRight` (table.go:89-99) to
  `style.Border.Left + v + style.Border.Right`. Verify: golden diff empty.
  Map to scenario `Output stays byte identical`
- [x] 2.5 Rewrite `Table.line` (table.go:102-109) into one pre-Grown
  `strings.Builder` writing the border glyphs, per-column repeated center
  glyph, and intersections; drop the outer `fmt.Sprintf`. Verify: golden
  diff empty. Map to scenario `Output stays byte identical`
- [x] 2.6 Rewrite `Table.String` (table.go:39-55) to accumulate into one
  `strings.Builder` via a small newline-writing helper, preserving the
  skip-empty-border behavior and per-row line output. Verify: golden diff
  empty, then run `make bench` and record per-benchmark `allocs/op` and
  `B/op` against the design.md baseline table.
  Map to scenario `Allocation gate met per benchmark`
- [x] 2.7 Run `make check` with only Tier 1 applied and fix any failure
  introduced by this group. Map to scenario `Output stays byte identical`

## 3. Group C — Tier 2 content pool

- [x] 3.1 Add the package-level `sync.Pool` of `*content` in content.go and
  make `newContent` (content.go:88-105) Get a pooled object, FULLY reset it
  (reuse the `c.c` backing when capacity suffices, else reallocate; zero
  `mw`/`w`), and fill it so `'\n'` splitting with `strings.TrimSpace` and the
  `realLength` max matches `strings.Split` semantics for empty, leading- and
  trailing-newline inputs. Verify: golden diff empty.
  Map to scenario `Output stays byte identical`
- [x] 3.2 Make `Table.initContent` (table.go:136-140) Put the cell's
  previous `*content` into the pool before replacing it at table.go:138;
  leave `dividerCell.lines` (cell.go:152-161) and `emptyCell` unpooled per
  the design's pool-scope decision. Verify: golden diff empty and
  `make test` green. Map to scenario `Output stays byte identical`
- [x] 3.3 Add the re-render-twice stability test in `concurrency_test.go`
  (`simpletable_test`): render each golden fixture's table twice
  consecutively and assert both outputs equal the fixture, proving pooled
  content is fully reset between calls.
  Map to scenario `Recycled cell content renders identically`
- [x] 3.4 Add the distinct-tables concurrency test in `concurrency_test.go`:
  separate goroutines each render their own table and compare against the
  expected golden output; run it under the race detector
  (`go test -race -run TestConcurrent .`) with no race reported.
  Map to scenario `Distinct tables render concurrently`
- [x] 3.5 Run `make bench` and verify EACH of the 7 `BenchmarkStyle*`
  results meets both per-benchmark ceilings in the design.md table
  (>=30% lower `allocs/op` AND >=15% lower `B/op`); record the observed
  numbers in the task report. If a benchmark misses a ceiling, stop and
  report — do not relax the gate.
  Map to scenario `Allocation gate met per benchmark`
- [x] 3.6 Run `make check` on the final tree (its `test` target executes the
  golden, stability, and concurrency tests; race detection itself was run in
  3.4) and fix any failure introduced by this group.
  Map to scenario `Output stays byte identical`

## 4. Group D — Refresh published comparison numbers (apply time)

- [x] 4.1 Run `make bench-compare` (equivalently
  `cd _example/03-benchmarks-with-others/bench && go test -bench . -benchmem`)
  on the MacBook benchmark machine and replace the fenced stdout block at
  README.md:102-108 (three `Benchmark*-N` lines plus the `PASS` and `ok ...`
  lines) inside the existing `<details><summary>(MacBook)</summary>` section
  with the VERBATIM new stdout; expect only the `BenchmarkSimpletable` line
  to change. Do NOT touch harness code, fixtures, or comparators, and do NOT
  add a `package-readme` spec delta. If the harness cannot be run on the
  target machine, leave README unchanged and block this task reporting why —
  never fabricate numbers.
  Map to scenario `Expand benchmark results`
