# Design: reduce per-render allocations of table output

Motivation, the gate contract, and rejected risks are in `proposal.md`; this
file records how the change is implemented and measured.

## Context

`Table.String()` (table.go:39-55) rebuilds all render metadata on every call:
`refresh` (table.go:69-74) discards rows/columns/spanned/dividers,
`initCellContent` → `initContent` (table.go:136-140) assigns a fresh
`*content` per cell via `newContent` (content.go:88-105), and every spacing
step goes through `fmt.Sprintf` + `strings.Repeat`. The root benchmark suite
(`benchmarks_test.go`, `BenchmarkStyle*` at lines 31-57, table shape built
once by `benchStyleTable` at lines 68-104, only `t.String()` measured at
lines 59-66) quantifies the churn: ~671 allocs/op for `BenchmarkStyleDefault`
at master 1c796ce.

`make bench` runs `go test -run '^$' -bench . -benchmem .` on the root
module. `make bench-compare` runs the SEPARATE module
`_example/03-benchmarks-with-others` (its go.mod replaces
`github.com/alexeyco/simpletable => ../..`), so the README comparison numbers
(`BenchmarkSimpletable`/`Termtables`/`UITable`, README.md:95-108, verbatim
harness stdout inside `<details><summary>(MacBook)</summary>`) come from that
harness, not from `make bench`; a library change moves only
`BenchmarkSimpletable`.

## Goals / Non-Goals

Goals:

- Each of the 7 root `BenchmarkStyle*` benchmarks reports >=30% lower
  `allocs/op` AND >=15% lower `B/op` versus the baseline below (per
  benchmark, not aggregate).
- Rendered output stays byte-identical, proven by a new golden-output test
  with committed `testdata/*.golden` fixtures frozen BEFORE any source edit.
- `make check` green, plus a `-race` distinct-tables test run in task 3.4.

Non-Goals:

- No `ns/op` gate (+/-20% noise; `allocs/op` is exactly reproducible).
- No exported API change; `Print`/`Println` signatures/behavior untouched.
- No same-table concurrency support (unsafe today, unsafe after).
- No new dependencies; no caching of metadata between `String()` calls; no
  streaming render pipeline; no pooling of `Cell.children`/`column`; no
  changes to the `_example/03-benchmarks-with-others` harness code,
  fixtures, or comparators.

## Baseline and gate ceilings

Recorded on this machine at 1c796ce via `make bench` (allocs/op exactly
reproducible; B/op within ±1; independently re-verified in this planning run).
Ceilings are floor(0.70 × allocs) and floor(0.85 × B).

| Benchmark        | baseline allocs/op | baseline B/op | gate allocs/op ≤ | gate B/op ≤ |
|------------------|--------------------:|--------------:|-----------------:|------------:|
| StyleDefault     |                 671 |         28785 |              469 |       24467 |
| StyleCompact     |                 635 |         27841 |              444 |       23664 |
| StyleCompactLite |                 635 |         27841 |              444 |       23664 |
| StyleCompactClassic |              587 |         25559 |              410 |       21725 |
| StyleMarkdown    |                 663 |         28289 |              464 |       24045 |
| StyleRounded     |                 671 |         28785 |              469 |       24467 |
| StyleUnicode     |                 691 |         32429 |              483 |       27564 |

Output dependencies preserved by every rewrite (content.go:108-150):
runewidth CJK/emoji widths via `realLength` = `runewidth.StringWidth`, ANSI
stripping for measurement while escapes remain in output (the
`stripAnsiEscape` fast path is already allocation-free when no ESC/0x9B is
present), per-line `strings.TrimSpace`, the center split `lft := padding / 2`,
and `padding <= 0` returning the line unchanged.

## Decisions

Tier 1 — pure string building, byte-identical, no shared state:

1. `content.line` (content.go:61-85, largest object producer, ~14% of
   allocations): build each formatted line into one `strings.Builder`
   (pre-`Grow` to `len(l) + padding`) and append spaces from a package-level
   reusable space run (`const spaceRun` plus a helper writing n spaces by
   looping over slices of it) instead of `fmt.Sprintf` + `strings.Repeat`.
   The three alignment branches keep their exact padding math.
2. `content.lines` (content.go:50-58): `make([]string, len(c.c))` with index
   assignment instead of `[]string{}` + append.
3. `tblRow.toStringSlice` (row.go:26-46): pre-size `l` (`make([][]string,
   len(r.Cells))`) and `ret` (`make([]string, len(l))` after transpose);
   hoist the cell separator `" "+style.Cell+" "` and the `" "+row+" "` wrap
   out of the per-line `fmt.Sprintf` calls into plain concatenation.
4. `Table.borderLeftRight` (table.go:89-99): `style.Border.Left + v +
   style.Border.Right` (compiler `concatstring3`) instead of
   `fmt.Sprintf("%s%s%s", ...)`.
5. `Table.line` (table.go:102-109): one pre-Grown `strings.Builder` writing
   the left glyph, each column's repeated center glyph (`col.getWidth()+2`
   repetitions), and intersections; the outer `fmt.Sprintf` is dropped.
6. `Table.String` (table.go:39-55): one `strings.Builder` fed by a small
   writer helper that appends lines separated by `'\n'`; preserves the exact
   "skip empty border/divider line" behavior (top/bottom border appended only
   when non-empty) and per-row line behavior.

Tier 2 — content pool:

7. A package-level `sync.Pool` of `*content` in content.go. `newContent`
   Gets an object and FULLY resets it before use (reuse the `c.c` backing
   slice when its capacity suffices, else reallocate; zero `mw` and `w`) and
   splits `Cell.Text` on `'\n'` with `strings.TrimSpace` per line and a
   `realLength` max, matching `strings.Split` semantics for empty, leading-
   and trailing-newline edge cases. `Table.initContent` (table.go:136-140,
   the `c.content = newContent(c.Text)` assignment at table.go:138) Puts the
   cell's previous `*content` back into the pool before replacing it.

Pool scope decision: the pool covers ONLY the `Cell.content` assignment
above. `dividerCell.lines` (cell.go:152-161) builds a throwaway inline
`&content{}` and stays as-is (a separate object path; `emptyCell.content` is
never initialized and its `height()`/`lines()` are not exercised in normal
rendering).

Safety reasoning for the pool: all render metadata is discarded and rebuilt
on every `String()` (`refresh` at table.go:69-74; `initContent` reassigns
`c.content`), so a released content object is dead before reuse. Plain
consequence: between two `String()` calls a `*Cell` keeps pointing at a
pooled object; the next call reassigns the field, so no public API observes
the retained field (it is unexported). Distinct tables rendered concurrently
stay correct because each cell points at a live object during its own
render; a `-race` test renders distinct tables from multiple goroutines to
prove it. Concurrent `String()` on one table remains unsupported, unchanged
from today.

Toolchain constraints verified: go1.27.4 toolchain, module `go 1.23`,
golangci-lint 2.15.0 with `unparam`, `gocritic`, `makezero`, `gocognit`,
`nolintlint` enabled; adding `sync.Pool` is not lint-blocked; `errcheck`
excludes only `fmt.Print`/`fmt.Println`, so the new builders must not create
unchecked error returns (they do not).

## Files to touch

| Path | Change |
|------|--------|
| `golden_test.go` (new, package `simpletable_test`) | Golden-output test: renders the `benchStyleTable` shape (benchmarks_test.go:68) in all 7 built-in styles plus multiline, ANSI-colored, CJK-wide, and spanned fixtures (extra fixtures rendered with `StyleUnicode`), compared byte-for-byte to `testdata/<case>.golden`; `-update-golden` flag regenerates. Group A test helpers are reused by Group C tests. |
| `testdata/style-default.golden`, `testdata/style-compact.golden`, `testdata/style-compact-lite.golden`, `testdata/style-compact-classic.golden`, `testdata/style-markdown.golden`, `testdata/style-rounded.golden`, `testdata/style-unicode.golden` (new) | Frozen current output of the `benchStyleTable` shape per style. |
| `testdata/multiline.golden`, `testdata/ansi.golden`, `testdata/cjk.golden`, `testdata/spanned.golden` (new) | Frozen current output of the extra input-shape fixtures. |
| `content.go` | Decisions 1, 2 (Tier 1) and 7 (pool: package-level `sync.Pool`, pooled+reset `newContent`, space-run helper). |
| `table.go` | Decisions 4, 5, 6 and the `initContent` Put (decision 7). |
| `row.go` | Decision 3. |
| `concurrency_test.go` (new, package `simpletable_test`) | `-race` distinct-tables test and the re-render-twice stability test (both compare against goldens). |
| `README.md` | Group D only: replace the verbatim fenced stdout block at README.md:102-108 with fresh `make bench-compare` stdout. No other README edits. |

`benchmarks_test.go` and `_example/03-benchmarks-with-others/**` are NOT
touched by this change.

## Sequencing (for the wave planner)

Groups B (Tier 1) and C (Tier 2) both edit `content.go`, `table.go`, and
`row.go`, so they must serialize: B then C. Group A must land before any
source change (it freezes the byte-exact target). Group D measures the final
binary and depends on B+C.

## Risks / Trade-offs

- Pool misuse risk: a content object released while still referenced would
  corrupt output. Mitigated by the invariant above (only objects replaced in
  `initContent` are Put), byte-exact goldens over multiline/empty/leading-
  and trailing-newline inputs, and the re-render-twice stability test.
- Golden fixtures freeze master output on this machine; if a fixture were
  regenerated after a source change it would silently bless a regression.
  Mitigated by Group A generating and verifying fixtures on the unchanged
  tree before any edit (tasks 1.1-1.2 gate the rest).
- The space-run helper loops and slices a package-level constant for runs
  longer than one iteration instead of allocating; the trade-off is repeated
  appends of a few bytes versus one `strings.Repeat` allocation per padding
  operation, plus one builder string per formatted line instead of two or
  more intermediates.
- Tier 2 adds ~5-6% wall-clock cost beyond Tier 1 (recorded as a rejected
  risk in `proposal.md`); accepted for the alloc/B-op reduction and lower GC
  frequency.
- The README numbers depend on running the comparison harness on MacBook
  hardware at apply time; if unavailable, Group D blocks rather than
  fabricates (recorded in tasks 4.1).
