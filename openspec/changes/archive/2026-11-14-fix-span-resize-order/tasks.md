## 1. Make spanning-cell resize order-independent

- [x] 1.1 Replace `Cell.resize` in `cell.go` with `Cell.spanWidth`, `Cell.growColumns`, and `Cell.syncColumns`; keep the `Span > 1` guard and the covered-width sum including inter-column spacing. Map to scenario `Resize a spanning cell`.
- [x] 1.2 In `Table.resizeColumns` in `table.go`, run `growColumns` over all spanning cells, then `syncColumns` over all spanning cells, before divider widths are set. Map to scenario `Later spanning cell widens shared columns`.

## 2. Cover the regression

- [x] 2.1 Add `TestSpannedCellsResizeAfterColumnGrowth` to `cell_test.go` reproducing issue #23 row order; assert the widened spanning-cell line and equal line widths.
- [x] 2.2 Run `make check`; expect tests, lint, format, and OpenSpec validation to pass and example rendering to be unchanged.
