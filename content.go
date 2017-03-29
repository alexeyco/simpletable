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

func (c *Content) StringSlice(width int) []string {
	s := []string{}
	if c.maxLineWidth > width {
		width = c.maxLineWidth
	}

	for _, l := range c.lines {
		w := utf8.RuneCountInString(l)
		s = append(s, fmt.Sprintf("%s%s", l, strings.Repeat(" ", width-w)))
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
