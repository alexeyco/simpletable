package simpletable

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	AlignLeft   = 0
	AlignCenter = 1
	AlignRight  = 2
)

type Cell interface {
	Len() int
	IsSpanned() bool
	SetWidth(int)
	Column() *Column
	SetColumn(*Column)
	String() string
}

type TextCell struct {
	Align    int
	Span     int
	Content  string
	width    int
	children []*EmptyCell
	column   *Column
}

func (c *TextCell) Len() int {
	return utf8.RuneCountInString(c.Content)
}

func (c *TextCell) IsSpanned() bool {
	return c.Span > 1
}

func (c *TextCell) SetWidth(width int) {
	c.width = width
}

func (c *TextCell) Column() *Column {
	return c.column
}

func (c *TextCell) SetColumn(column *Column) {
	c.column = column
}

func (c *TextCell) Resize() {
	if !c.IsSpanned() {
		return
	}

	s := c.column.Width()
	for _, ch := range c.children {
		s += ch.Column().Width()
	}

	s += len(c.children) * 3

	if s > c.Len() {
		c.SetWidth(s)
	} else {
		cols := []*Column{
			c.column,
		}

		for _, ch := range c.children {
			cols = append(cols, ch.column)
		}

		c.column.Table.incrementColumns(cols, c.Len()-s)
	}
}

func (c *TextCell) String() string {
	l := c.width - c.Len()
	if l == 0 {
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

func (c *TextCell) calcSize(v, p int) []int {
	r := []int{}

	return r
}

type Divider struct {
	Span     int
	width    int
	children []*EmptyCell
	column   *Column
}

func (d *Divider) Len() int {
	return d.width
}

func (d *Divider) IsSpanned() bool {
	return d.Span > 1
}

func (d *Divider) SetWidth(width int) {
	d.width = width
}

func (d *Divider) Column() *Column {
	return d.column
}

func (d *Divider) SetColumn(column *Column) {
	d.column = column
}

func (d *Divider) String() string {
	s := d.column.Table.style.Divider
	return d.column.Table.line(s.Left, s.Center, s.Right, s.Intersection)
}

type EmptyCell struct {
	parent Cell
	width  int
	column *Column
}

func (e *EmptyCell) Len() int {
	return 0
}

func (e *EmptyCell) IsSpanned() bool {
	return false
}

func (e *EmptyCell) SetWidth(width int) {
	e.width = width
}

func (e *EmptyCell) Column() *Column {
	return e.column
}

func (e *EmptyCell) SetColumn(column *Column) {
	e.column = column
}

func (e *EmptyCell) String() string {
	return ""
}
