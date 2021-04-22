package simpletable

import (
	"fmt"
	"math"
	"strings"

	"github.com/alexeyco/simpletable/xterm"
	"github.com/mattn/go-runewidth"
)

const (
	newLine = "\n"
	space   = " "
)

type Col struct {
	options       Options
	minWidth      int
	originalLines []string
	resizedLines  []string
	width         int
	height        int
}

func (c *Col) Width() int {
	return c.width
}

func (c *Col) Height() int {
	return c.height
}

func (c *Col) Lines() []string {
	return c.resizedLines
}

func (c *Col) Resize(width, height int) {
	c.resizeWidth(width)
	c.align()
	c.resizeHeight(height)
}

func (c *Col) resizeWidth(width int) {
	if width < c.minWidth {
		width = c.minWidth
	}

	if width == c.width {
		return
	}

	var lines []string
	for _, originalLine := range c.originalLines {
		words := strings.Split(originalLine, space)

		var (
			line      []string
			lineWidth int
		)

		for _, word := range words {
			wordWidth := runewidth.StringWidth(word)
			overallLineWidth := lineWidth + wordWidth + len(line)

			switch {
			case overallLineWidth <= width:
				line = append(line, word)
				lineWidth += wordWidth
			case overallLineWidth > width:
				lines = append(lines, strings.Join(line, space))

				line = []string{word}
				lineWidth = wordWidth
			}
		}

		if len(line) > 0 {
			lines = append(lines, strings.Join(line, space))
		}
	}

	c.resizedLines = lines
	c.width = width
	c.height = len(c.resizedLines)
}

func (c *Col) align() {
	for n := range c.resizedLines {
		var (
			spanLeft  int
			spanRight int
		)

		lineWidth := runewidth.StringWidth(c.resizedLines[n])

		switch c.options.Align {
		case Left:
			spanRight = c.width - lineWidth
		case Center:
			span := c.width - lineWidth

			spanLeft = int(math.Floor(float64(span) / float64(2)))
			spanRight = span - spanLeft
		case Right:
			spanLeft = c.width - lineWidth
		}

		c.resizedLines[n] = fmt.Sprintf("%s%s%s",
			strings.Repeat(space, spanLeft),
			c.resizedLines[n],
			strings.Repeat(space, spanRight))
	}
}

func (c *Col) resizeHeight(height int) {
	if height <= c.height {
		return
	}

	span := height - c.height
	for i := 0; i < span; i++ {
		c.resizedLines = append(c.resizedLines, strings.Repeat(space, c.width))
	}

	c.height = height
}

func Column(text string, options ...Option) *Col {
	o := Options{
		Span: 1,
	}

	for _, option := range options {
		option(&o)
	}

	var (
		minWidth int
		width    int
	)

	if text != "" {
		text = xterm.Strip(text)
	}

	lines := strings.Split(text, newLine)
	for n := range lines {
		lines[n] = strings.TrimSpace(lines[n])

		l := runewidth.StringWidth(lines[n])
		if l > width {
			width = l
		}

		words := strings.Split(lines[n], space)
		for _, word := range words {
			w := runewidth.StringWidth(word)
			if w > minWidth {
				minWidth = w
			}
		}
	}

	col := &Col{
		options:       o,
		minWidth:      minWidth,
		originalLines: lines,
		resizedLines:  lines,
		width:         width,
		height:        len(lines),
	}

	col.align()

	return col
}
