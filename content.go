package simpletable

import (
	"strings"
	"unicode/utf8"
	"fmt"
)

type content struct {
	c     []string
}

func (c *content) width() int {
	w := 0

	for _, r := range c.c {
		l := utf8.RuneCountInString(r)
		if l > w {
			w = l
		}
	}

	return w
}

func (c *content) height() int {
	return len(c.c)
}

func (c *content) setHeight(h int) {
	l := c.height()
	if h <= l {
		return
	}

	for i := 0; i < h-l; i++ {
		c.c = append(c.c, "")
	}
}

func (c *content) lines(a int) []string {
	r := []string{}

	for _, l := range c.c {
		r = append(r, c.line(l, a))
	}

	return r
}

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

func newContent(s string) *content {
	c := strings.Split(s, "\n")

	for i, v := range c {
		c[i] = strings.TrimSpace(v)
	}

	return &content{
		c: c,
	}
}
