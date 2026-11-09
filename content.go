package simpletable

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

// content is cell content.
type content struct {
	c  []string // content lines (trimmed)
	mw int      // maximal content width (f.e, for multilines columns)
	w  int      // meta content width
}

// width returns content width.
func (c *content) width() int {
	m := c.mw
	if m > c.w {
		return m
	}

	return c.w
}

// setWidth sets content block width.
func (c *content) setWidth(w int) {
	c.w = w
}

// height returns content height.
func (c *content) height() int {
	return len(c.c)
}

// setHeigth sets content height.
func (c *content) setHeight(h int) {
	l := c.height()
	if h <= l {
		return
	}

	for i := 0; i < h-l; i++ {
		c.c = append(c.c, "")
	}
}

// lines returns content as a string slice.
func (c *content) lines(a int) []string {
	r := []string{}

	for _, l := range c.c {
		r = append(r, c.line(l, a))
	}

	return r
}

// line formats a content line.
func (c *content) line(l string, a int) string {
	padding := c.width() - realLength(l)
	if padding <= 0 {
		return l
	}

	switch a {
	case AlignLeft:
		l = fmt.Sprintf("%s%s", l, strings.Repeat(" ", padding))

	case AlignCenter:
		lft := padding / 2
		rgt := padding - lft

		left := strings.Repeat(" ", lft)
		right := strings.Repeat(" ", rgt)

		l = fmt.Sprintf("%s%s%s", left, l, right)

	case AlignRight:
		l = fmt.Sprintf("%s%s", strings.Repeat(" ", padding), l)
	}

	return l
}

// newContent returns a new content object.
func newContent(s string) *content {
	c := strings.Split(s, "\n")
	mw := 0

	for i, v := range c {
		c[i] = strings.TrimSpace(v)
		w := realLength(c[i])

		if w > mw {
			mw = w
		}
	}

	return &content{
		c:  c,
		mw: mw,
	}
}

// stripAnsiEscape returns a string without ANSI escape sequences, such as color codes.
func stripAnsiEscape(s string) string {
	if strings.IndexByte(s, '\x1b') < 0 && strings.IndexByte(s, '\x9b') < 0 {
		return s
	}

	var result strings.Builder
	for i := 0; i < len(s); {
		start := i
		switch {
		case s[i] == '\x1b' && i+1 < len(s) && s[i+1] == '[':
			i += 2
		case s[i] == '\x9b':
			i++
		default:
			_ = result.WriteByte(s[i])
			i++
			continue
		}

		if end, ok := csiSequenceEnd(s, i); ok {
			i = end
			continue
		}

		_, _ = result.WriteString(s[start:i])
	}

	return result.String()
}

func csiSequenceEnd(s string, i int) (int, bool) {
	for i < len(s) && s[i] >= 0x30 && s[i] <= 0x3f {
		i++
	}
	for i < len(s) && s[i] >= 0x20 && s[i] <= 0x2f {
		i++
	}
	if i < len(s) && s[i] >= 0x40 && s[i] <= 0x7e {
		return i + 1, true
	}

	return 0, false
}

// realLength returns the visible string width without ANSI escape sequences.
func realLength(s string) int {
	return runewidth.StringWidth(stripAnsiEscape(s))
}
