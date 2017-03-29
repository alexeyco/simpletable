package simpletable

import (
	"strings"
	"unicode/utf8"
	"fmt"
	"log"
)

const (
	AlignLeft   = -1
	AlignCenter = 0
	AlignRight  = 1
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

	c.SetWidth(s + (len(c.children) * 3))

	// TODO: resize columns if it needed
}

func (c *TextCell) String() string {
	l := c.width - c.Len()
	if l == 0 {
		return c.Content
	}

	if l < 0 {
		log.Fatalln(c.width, c.Len())
	}

	return fmt.Sprintf("%s%s", c.Content, strings.Repeat(" ", l))
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
	return strings.Repeat("=", d.width)
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
