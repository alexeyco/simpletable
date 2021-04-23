// Package grid allows to manage table grid
package grid

// Cell grid cell.
type Cell struct {
	Row      int
	Col      int
	HSpan    int
	VSpan    int
	Width    int
	Height   int
	mergedTo *Cell
}

// Grid structure.
type Grid struct {
	rows    int
	columns int
	cells   [][]*Cell
}

func (g *Grid) cell(rowIdx, colIdx int) *Cell {
	if rowIdx < 0 || rowIdx >= g.rows {
		return nil
	}

	if colIdx < 0 || colIdx >= g.columns {
		return nil
	}

	return g.cells[rowIdx][colIdx]
}

// SetWidth changes grid column width.
func (g *Grid) SetWidth(colIdx, width int) {
	c := &Collection{}

	for rowIdx := range g.cells {
		if cell := g.cell(rowIdx, colIdx); cell != nil && cell.Width < width {
			if cell.mergedTo != nil && !c.Exists(cell.mergedTo) {
				cell.mergedTo.Width = cell.mergedTo.Width - cell.Width + width
				c.Append(cell.mergedTo)
			}

			cell.Width = width
		}
	}
}

// SetHeight changes row height.
func (g *Grid) SetHeight(rowIdx, height int) {
	c := &Collection{}

	for colIdx := range g.cells[rowIdx] {
		if cell := g.cell(rowIdx, colIdx); cell != nil && cell.Height < height {
			if cell.mergedTo != nil && !c.Exists(cell.mergedTo) {
				cell.mergedTo.Height = cell.mergedTo.Height - cell.Height + height
				c.Append(cell.mergedTo)
			}

			cell.Height = height
		}
	}
}

// VExpand expand column vertically.
func (g *Grid) VExpand(rowIdx, colIdx, span int) {
	if span < 1 {
		return
	}

	cell := g.cell(rowIdx, colIdx)
	if cell == nil || cell.mergedTo != nil {
		return
	}

	fromIdx := cell.Row + cell.VSpan
	for celRowIdx := fromIdx; celRowIdx <= fromIdx+span; celRowIdx++ {
		var heightIncreased bool

		for cellColIdx := cell.Col; cellColIdx < cell.Col+cell.HSpan; cellColIdx++ {
			if c := g.cell(celRowIdx, cellColIdx); c != nil && c.mergedTo != nil {
				return
			}
		}

		for cellColIdx := cell.Col; cellColIdx < cell.Col+cell.HSpan; cellColIdx++ {
			if c := g.cell(celRowIdx, cellColIdx); c != nil {
				c.mergedTo = cell
				if !heightIncreased {
					cell.Height += c.Height
					heightIncreased = true
				}
			}
		}
	}

	cell.VSpan = span
}

// HExpand expand column vertically.
func (g *Grid) HExpand(rowIdx, colIdx, span int) {
	if span < 1 {
		return
	}

	cell := g.cell(rowIdx, colIdx)
	if cell == nil || cell.mergedTo != nil {
		return
	}

	fromIdx := cell.Col + cell.HSpan
	for cellColIdx := fromIdx; cellColIdx <= fromIdx+span; cellColIdx++ {
		var widthIncreased bool

		for cellRowIdx := cell.Row; cellRowIdx < cell.Row+cell.VSpan; cellRowIdx++ {
			if c := g.cell(cellRowIdx, cellColIdx); c != nil && c.mergedTo != nil {
				return
			}
		}

		for cellRowIdx := cell.Row; cellRowIdx < cell.Row+cell.VSpan; cellRowIdx++ {
			if c := g.cell(cellRowIdx, cellColIdx); c != nil {
				c.mergedTo = cell

				if !widthIncreased {
					cell.Width += c.Width
					widthIncreased = true
				}
			}
		}
	}

	cell.HSpan = span
}

// Iterate over grid cells.
func (g *Grid) Iterate(f func(*Cell)) {
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.mergedTo == nil {
				f(cell)
			}
		}
	}
}

// Cells returns grid cells slice.
func (g *Grid) Cells() []*Cell {
	var cells []*Cell
	g.Iterate(func(cell *Cell) {
		cells = append(cells, cell)
	})

	return cells
}

// New returns new grid instance with specified rows and columns number.
func New(rows, columns int) *Grid {
	cells := make([][]*Cell, rows)
	for rowIdx := 0; rowIdx < rows; rowIdx++ {
		cells[rowIdx] = make([]*Cell, columns)

		for colIdx := 0; colIdx < columns; colIdx++ {
			cells[rowIdx][colIdx] = &Cell{
				Row:   rowIdx,
				Col:   colIdx,
				HSpan: 1,
				VSpan: 1,
			}
		}
	}

	return &Grid{
		rows:    rows,
		columns: columns,
		cells:   cells,
	}
}
