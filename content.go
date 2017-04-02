package simpletable

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// content is a cell content
type content struct {
	c []string // content lines (trimmed)
	w int      // meta content width
}

// width returns maximum content lines width
func (c *content) maxLinewidth() int {
	w := 0

	for _, r := range c.c {
		l := utf8.RuneCountInString(r)
		if l > w {
			w = l
		}
	}

	return w
}

// width returns content width
func (c *content) width() int {
	m := c.maxLinewidth()
	if m > c.w {
		return m
	}

	return c.w
}

// setWidth sets content block width
func (c *content) setWidth(w int) {
	c.w = w
}

// height returns content height
func (c *content) height() int {
	return len(c.c)
}

// setHeigth sets content height
func (c *content) setHeight(h int) {
	l := c.height()
	if h <= l {
		return
	}

	for i := 0; i < h-l; i++ {
		c.c = append(c.c, "")
	}
}

// lines returns content as string slice
func (c *content) lines(a int) []string {
	r := []string{}

	for _, l := range c.c {
		r = append(r, c.line(l, a))
	}

	return r
}

// line formats content line
func (c *content) line(l string, a int) string {
	len := c.width() - utf8.RuneCountInString(l)
	if len <= 0 {
		return l
	}

	switch a {
	case AlignLeft:
		l = fmt.Sprintf("%s%s", l, strings.Repeat(" ", len))

	case AlignCenter:
		lft := len / 2
		rgt := len - lft

		left := strings.Repeat(" ", lft)
		right := strings.Repeat(" ", rgt)

		l = fmt.Sprintf("%s%s%s", left, l, right)

	case AlignRight:
		l = fmt.Sprintf("%s%s", strings.Repeat(" ", len), l)
	}

	return l
}

// newContent returns new content object
func newContent(s string) *content {
	c := strings.Split(s, "\n")

	for i, v := range c {
		c[i] = strings.TrimSpace(v)
	}

	return &content{
		c: c,
	}
}
