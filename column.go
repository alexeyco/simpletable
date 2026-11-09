package simpletable

// tblColumn is a meta table column.
type tblColumn struct {
	Cells []cellInterface
	Table *Table
	width int
}

// resize resizes the column to the maximum nested cell length.
func (c *tblColumn) resize() {
	c.width = 0

	for _, cell := range c.Cells {
		if !cell.isSpanned() {
			w := cell.len()
			if w > c.width {
				c.width = w
			}
		}
	}

	c.incrementWidth(0)
}

// incrementWidth increments nested cell widths by i.
func (c *tblColumn) incrementWidth(i int) {
	c.width += i
	for _, cell := range c.Cells {
		cell.setWidth(c.width)
	}
}

// getWidth returns the column width.
func (c *tblColumn) getWidth() int {
	return c.width
}
