package simpletable

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Content struct {
	lines        []string
	maxLineWidth int
}

func (c *Content) Width() int {
	return c.maxLineWidth
}

func (c *Content) Height() int {
	return len(c.lines)
}

func (c *Content) Capitalize() {
	r := []string{}

	for _, l := range c.lines {
		r = append(r, strings.ToUpper(l))
	}

	c.lines = r
}

func (c *Content) StringSlice(width int, align string) []string {
	s := []string{}
	if c.maxLineWidth > width {
		width = c.maxLineWidth
	}

	for _, l := range c.lines {
		w := utf8.RuneCountInString(l)

		switch align {
		case AlignRight:
			s = append(s, fmt.Sprintf("%s%s", strings.Repeat(" ", width-w), l))
		case AlignCenter:
		default:
			s = append(s, fmt.Sprintf("%s%s", l, strings.Repeat(" ", width-w)))
		}
	}

	return s
}

func newContent(t string) *Content {
	var w int

	lines := strings.Split(t, "\n")
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return &Content{
		lines:        lines,
		maxLineWidth: w,
	}
}
