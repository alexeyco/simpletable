# Reduce per-render allocations of table output

## Why now

A profile of master (1c796ce) shows `Table.String()` costs ~671 allocs/op on
`BenchmarkStyleDefault`. The largest producers are pure GC churn with no
behavior change: `content.line` ~14%, `newContent` + `genSplit` ~18%,
`fmt.Sprintf` call sites ~8%. Every render reallocates the space runs, the
per-cell `*content` objects, and the metadata slices that could be reused or
built in place. The root benchmark suite (`benchmarks_test.go`, `make bench`)
measures only `t.String()` on a fixed table shape, so the cost is directly
attributable and exactly reproducible via `allocs/op`.

## What changes

- Tier 1 — byte-identical string-building rewrites, no shared state:
  `strings.Builder` in `content.line`, `Table.line`, and `Table.String`;
  plain concatenation instead of `fmt.Sprintf` in `borderLeftRight` and
  `toStringSlice`; a package-level reusable space run instead of
  `strings.Repeat`; pre-sized slices in `content.lines`, `toStringSlice`, and
  `String`.
- Tier 2 — a `sync.Pool` for the internal `*content` object: `newContent`
  Gets and fully resets a pooled object; `initContent` Puts the cell's
  previous `*content` before replacing it. Rendered output and the exported
  API are unchanged; no new dependencies (stdlib `sync`/`strings` only).
- A new golden-output regression test with committed `testdata/*.golden`
  fixtures (all 7 built-in styles plus multiline, ANSI-colored, CJK-wide, and
  spanned inputs) is added BEFORE any source change and freezes the current
  byte-exact output as the acceptance target.
- Acceptance gate: on the root `make bench` suite, EACH of the 7
  `BenchmarkStyle*` benchmarks must report >=30% lower `allocs/op` AND >=15%
  lower `B/op` versus the baseline recorded in `design.md` (per benchmark, not
  aggregate), plus byte-identical golden output and `make check` green.
- README benchmark block refresh at apply time (Group D): verbatim new stdout
  from `make bench-compare` replaces the fenced stdout block at
  README.md:102-108 inside the existing `<details><summary>(MacBook)</summary>`
  section. No numbers are fabricated if the harness cannot run on the target
  machine — the task is blocked instead.

Out of scope: caching metadata between `String()` calls; restructuring the
render pipeline to stream rows; the `_example/03-benchmarks-with-others`
harness code, fixtures, or comparators; any `package-readme` spec delta; the
README Quick start left-trim; pooling `Cell.children` or `column`.

## Impact on capabilities

- `table-output`: one ADDED requirement "Bound per-render allocations" with
  four scenarios (byte-identical output, per-benchmark allocation gate,
  recycled content re-render stability, distinct-table concurrency under
  `-race`). Existing requirements are untouched; `Print`/`Println`
  signatures and behavior are not modified.
- `package-readme`: no delta in this change. The existing requirement
  "Disclose current benchmark results in a collapsible section" is satisfied
  by the Group D README update at apply time.
- Compatibility contracts (AGENTS.md): rendered output stays byte-identical
  and the exported API is unchanged. Concurrent `String()` on the SAME table
  is unsafe today (`String` mutates `t.rows`/`content`) and remains so — this
  change adds no same-table concurrency guarantee, but the pool must not make
  distinct-table concurrent rendering regress.

**DECISION:** Accept Tier 1 + Tier 2 together as one change; the content pool
is retained on top of Tier 1 for its alloc/GC win despite its small marginal
wall-clock contribution (see the rejected risks below).

**DECISION:** The delta is `## ADDED Requirements`, not a MODIFIED of "Return
newline-separated table text", so the existing requirement headers and the
strict 500-char description merge are left untouched.

**REJECTED-RISK:** ns/op regression does not block acceptance — reason:
+/-20% noise, the deterministic allocs/op and B/op gate the change.

**REJECTED-RISK:** the content pool adds only ~5-6% wall-clock beyond Tier 1 —
kept for ~-50% allocs / -28% B/op and lower GC frequency.

**ASSUMPTION:** the >=30%/>=15% gate is evaluated against the exact baseline
table recorded on this machine at 1c796ce (in `design.md`); a different CPU
can shift ns/op but not allocs/op — tested by one read-only local run:
`go test -run '^$' -bench . -benchmem . → BenchmarkStyleDefault-12 33909
35371 ns/op 28785 B/op 671 allocs/op`, with all seven `allocs/op` values
reproduced exactly and `B/op` within ±1.
