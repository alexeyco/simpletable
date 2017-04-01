package simpletable

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	// AlignLeft sets cell left alignment (default)
	AlignLeft = 0

	// AlignCenter sets cell center alignment
	AlignCenter = 1

	// AlignRight sets cell right alignment
	AlignRight = 2
)

// cellInterface is a basic cell interface
type cellInterface interface {
	len() int              // Returns cell text length
	isSpanned() bool       // Returns true if cell spanned
	setWidth(int)          // Sets cell width
	getColumn() *tblColumn // Returns parent column
	setColumn(*tblColumn)  // Sets parent column
	toString() string      // Returns cell content as a string
}

// Cell is a table cell
type Cell struct {
	Align    int          // Cell alignment
	Span     int          // span cell to right (1 - default)
	Content  string       // Cell content
	width    int          // Cell width
	children []*emptyCell // Nested empty cells
	column   *tblColumn   // Parent column
}

func (c *Cell) len() int {
	return utf8.RuneCountInString(c.Content)
}

func (c *Cell) isSpanned() bool {
	return c.Span > 1
}

func (c *Cell) setWidth(width int) {
	c.width = width
}

func (c *Cell) getColumn() *tblColumn {
	return c.column
}

func (c *Cell) setColumn(column *tblColumn) {
	c.column = column
}

func (c *Cell) resize() {
	if !c.isSpanned() {
		return
	}

	s := c.column.getWidth()
	for _, ch := range c.children {
		s += ch.getColumn().getWidth()
	}

	s += len(c.children) * 3

	if s > c.len() {
		c.setWidth(s)
	} else {
		cols := []*tblColumn{
			c.column,
		}

		for _, ch := range c.children {
			cols = append(cols, ch.column)
		}

		c.column.Table.incrementColumns(cols, c.len()-s)
	}
}

func (c *Cell) toString() string {
	l := c.width - c.len()
	if l <= 0 {
		return c.Content
	}

	var s string
	switch c.Align {
	case AlignLeft:
		s = fmt.Sprintf("%s%s", c.Content, strings.Repeat(" ", l))

	case AlignCenter:
		lft := l / 2
		rgt := l - lft

		left := strings.Repeat(" ", lft)
		right := strings.Repeat(" ", rgt)

		s = fmt.Sprintf("%s%s%s", left, c.Content, right)

	case AlignRight:
		s = fmt.Sprintf("%s%s", strings.Repeat(" ", l), c.Content)
	}

	return s
}

// dividerCell is table divider cell
type dividerCell struct {
	span     int          // Divider span
	width    int          // Divider width
	children []*emptyCell // Nested empty meta cells
	column   *tblColumn   // Divider parent column
}

func (d *dividerCell) len() int {
	return d.width
}

func (d *dividerCell) isSpanned() bool {
	return d.span > 1
}

func (d *dividerCell) setWidth(width int) {
	d.width = width
}

func (d *dividerCell) getColumn() *tblColumn {
	return d.column
}

func (d *dividerCell) setColumn(column *tblColumn) {
	d.column = column
}

func (d *dividerCell) toString() string {
	s := d.column.Table.style.Divider
	return d.column.Table.line(s.Left, s.Center, s.Right, s.Intersection)
}

// emptyCell is meta cell, used when cell is spanned
type emptyCell struct {
	parent cellInterface // Parent cell
	width  int           // Cell width
	column *tblColumn    // Parent cell column
}

func (e *emptyCell) len() int {
	return 0
}

func (e *emptyCell) isSpanned() bool {
	return false
}

func (e *emptyCell) setWidth(width int) {
	e.width = width
}

func (e *emptyCell) getColumn() *tblColumn {
	return e.column
}

func (e *emptyCell) setColumn(column *tblColumn) {
	e.column = column
}

func (e *emptyCell) toString() string {
	return ""
}
