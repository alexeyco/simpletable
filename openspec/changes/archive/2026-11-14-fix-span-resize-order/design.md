# Design

## Context

`Table.resizeColumns` sizes columns from non-spanned cells, then adjusts cells that span several columns, and finally sets divider widths from the resulting table size. Spanned cells share columns with each other and with plain cells, so widening one span's columns changes the total width of every other span covering them. See proposal.md - Why for the failure this ordering produced (issue #23).

Constraints: the exported `Cell` fields and the rendered output of existing valid tables are compatibility contracts, and the resize path is on the hot path of every `Table.String` call, exercised by `make bench`.

## Goals / Non-Goals

**Goals:**

- Rendered geometry must not depend on the order in which spanned cells appear.
- Keep the public API and the output of tables that already rendered correctly unchanged.
- No additional allocations per rendered table.

**Non-Goals:**

- Do not compute globally minimal column widths when several spans compete for the same columns; each span only guarantees its own content fits, as before.
- Do not change divider, header, footer, or style handling.

## Decisions

Split the old single method `Cell.resize` into three unexported pieces and run them as two passes in `table.go`: `growColumns` over all spanned cells first, then `syncColumns`.

- Two ordered passes instead of repeating `Cell.resize` until widths stop changing. Column growth is monotone: after a cell is grown, later growth only increases the spans that already fit it, and every remaining mismatch is one-directional (a cell is narrower than its final span). A fixed-point loop would converge to the same result with extra iterations and a hard stop condition to maintain; the two-pass form makes the invariant "growth finished before any width is fixed" structural.
- Comparisons use `c.len()` (content width including meta width), not `content.mw` directly. During the growth pass a spanned cell's meta width never exceeds its first column width, which is at most its span total, so `c.len() > span` is equivalent to "content does not fit" without reaching into `content` internals.
- `spanWidth` sums column widths without materializing a column slice; the slice is built only inside the growth branch that `incrementColumns` needs. An earlier variant returned the slice from `span()` and cost +4 allocs/op because `syncColumns` never uses it.

Order inside `resizeColumns` is preserved: columns sized from non-spanned cells, span growth pass, span sync pass, divider widths. Dividers must read the final table size, so they stay last.

## Risks / Trade-offs

- A spanned cell processed early in the growth pass may cause columns to grow that a later span would already have covered, slightly overshooting the tightest layout. Pre-existing behavior of `incrementColumns`/`carve`; the fix keeps growth order-neutral rather than optimizing width.
- Two passes iterate spanned cells twice instead of once. The pass is O(spanned cells × span), benchmark-verified to keep `sec/op`, `B/op`, and `allocs/op` at baseline.
