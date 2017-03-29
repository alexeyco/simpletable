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

func (c *Content) String() string {
	s := []string{}

	for _, l := range c.lines {
		w := utf8.RuneCountInString(l)
		s = append(s, fmt.Sprintf("%s%s", l, strings.Repeat(" ", c.maxLineWidth-w)))
	}

	return strings.Join(s, "\n")
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
