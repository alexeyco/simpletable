package simpletable

import (
	"strings"
	"sync"

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
	r := make([]string, len(c.c))

	for i, l := range c.c {
		r[i] = c.line(l, a)
	}

	return r
}

// spaceRun is a reusable run of 32 spaces written in slices by writeSpaces,
// so padding never allocates a fresh string per formatted line.
const spaceRun = "                                "

// writeSpaces writes n spaces to b, looping over full slices of spaceRun
// for runs longer than one pass and writing the remainder as the last one.
func writeSpaces(b *strings.Builder, n int) {
	for n > len(spaceRun) {
		_, _ = b.WriteString(spaceRun)
		n -= len(spaceRun)
	}

	_, _ = b.WriteString(spaceRun[:n])
}

// line formats a content line.
func (c *content) line(l string, a int) string {
	padding := c.width() - realLength(l)
	if padding <= 0 {
		return l
	}

	var b strings.Builder
	b.Grow(len(l) + padding)

	switch a {
	case AlignLeft:
		_, _ = b.WriteString(l)
		writeSpaces(&b, padding)

	case AlignCenter:
		lft := padding / 2
		rgt := padding - lft

		writeSpaces(&b, lft)
		_, _ = b.WriteString(l)
		writeSpaces(&b, rgt)

	case AlignRight:
		writeSpaces(&b, padding)
		_, _ = b.WriteString(l)

	default:
		return l
	}

	return b.String()
}

// contentPool recycles *content objects between renders: initContent
// returns each cell's previous object to the pool, and newContent takes a
// reset one from it. Only the Cell.content assignment participates;
// dividerCell builds a throwaway content inline and emptyCell content is
// never initialized, so neither uses the pool.
var contentPool = sync.Pool{
	New: func() any {
		return &content{}
	},
}

// newContent returns a pooled content object loaded from s. The object is
// fully reset before it is filled: the c backing slice is reused when its
// capacity fits the line count and reallocated otherwise, every stale line
// is cleared, and mw and w are zeroed, so no state of a previous use can
// leak into the new one.
func newContent(s string) *content {
	// The discarded ok is safe: the pool's New returns *content and the
	// only Put site, initContent, never stores a nil object.
	c, _ := contentPool.Get().(*content)

	// One line per '\n' plus the remainder matches strings.Split(s, "\n"):
	// "" yields one empty line, and leading, trailing, and interior
	// newlines yield empty lines.
	n := strings.Count(s, "\n") + 1

	if cap(c.c) < n {
		c.c = make([]string, n)
	} else {
		// Reuse the backing array and clear it fully, so stale lines from
		// a previous, longer use are released and cannot resurface
		// through a later reuse.
		c.c = c.c[:cap(c.c)]
		clear(c.c)
		c.c = c.c[:n]
	}

	mw := 0

	for i := 0; i < n; i++ {
		l := s
		if idx := strings.IndexByte(s, '\n'); idx >= 0 {
			l, s = s[:idx], s[idx+1:]
		}

		l = strings.TrimSpace(l)
		c.c[i] = l

		if w := realLength(l); w > mw {
			mw = w
		}
	}

	c.mw = mw
	c.w = 0

	return c
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
