package simpletable

type Column struct {
	Cells []Cell
	width int
}

func (c *Column) Resize() {
	c.width = 0

	for _, cell := range c.Cells {
		if !cell.IsSpanned() {
			w := cell.Len()
			if w > c.width {
				c.width = w
			}
		}
	}

	c.IncrementWidth(0)
}

func (c *Column) IncrementWidth(i int) {
	c.width += i
	for _, cell := range c.Cells {
		cell.SetWidth(c.width)
	}
}

func (c *Column) Width() int {
	return c.width
}
