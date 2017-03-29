package simpletable

import (
	"strings"
	"unicode/utf8"
)

const (
	AlignLeft   = -1
	AlignCenter = 0
	AlignRight  = 1
)

type Cell interface {
	Len() int
	SetWidth(int)
	String() string
}

type Parent interface {
	Len() int
	ResetWidth()
}

type TextCell struct {
	Align    int
	Span     int
	Content  string
	width    int
	children []*EmptyCell
}

func (c *TextCell) Len() int {
	return utf8.RuneCountInString(c.String())
}

func (c *TextCell) SetWidth(width int) {
	c.width = width
}

func (c *TextCell) ResetWidth() {

}

func (c *TextCell) String() string {
	return c.Content
}

type Divider struct {
	Span     int
	width    int
	children []*EmptyCell
}

func (d *Divider) Len() int {
	return d.width
}

func (d *Divider) SetWidth(width int) {
	d.width = width
}

func (d *Divider) ResetWidth() {

}

func (d *Divider) String() string {
	return strings.Repeat("=", d.width)
}

type EmptyCell struct {
	parent Parent
	width  int
}

func (e *EmptyCell) Len() int {
	return 0
}

func (e *EmptyCell) SetWidth(width int) {
	e.width = width
	e.parent.ResetWidth()
}

func (e *EmptyCell) String() string {
	return ""
}
