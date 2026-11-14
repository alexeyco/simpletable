# Fix spanned cell resize order

## Why now

Rendering a table where a multi-line cell spans columns next to a later row whose spanning content widens the shared columns produces broken output: the earlier spanning cell keeps the stale narrower width captured before the growth, so its rendered lines no longer reach the row separator (issue #23). The result depends on body row order, which makes valid tables render incorrectly.

## What changes

Split the single-pass spanning-cell resize in `Table.resizeColumns` into two passes over the spanning cells: first widen columns so every spanning cell's content fits its covered span, then set each spanning cell's width to the final total of its covered columns. Column growth never revises an already-processed spanning cell again, so rendered rows stay aligned regardless of row order. No exported API changes.

## Impact on capabilities

Modify the `cell-layout` capability: the existing `Size columns and spanning cells` requirement gains the order-independence guarantee. Rendering, styles, and other capabilities do not change.
