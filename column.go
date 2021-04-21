package simpletable

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

const (
	newLine = "\n"
	space   = " "
)

// stripAnsiEscapeExp is a regular expression to clean ANSI Control sequences
// feat https://stackoverflow.com/questions/14693701/how-can-i-remove-the-ansi-escape-sequences-from-a-string-in-python#33925425
var stripAnsiEscapeExp = regexp.MustCompile(`(\x9B|\x1B\[)[0-?]*[ -/]*[@-~]`)

type Col struct {
	options                 Options
	original, resized       []string
	minWidth, width, height int
}

func (c *Col) Width() int {
	return c.width
}

func (c *Col) Height() int {
	return c.height
}

func (c *Col) Lines() []string {
	return c.resized
}

func (c *Col) Resize(width, height int) {
	c.resizeWidth(width)
	c.resizeHeight(height)
}

func (c *Col) resizeWidth(width int) {
	if width < c.minWidth {
		width = c.minWidth
	}
	if width == c.width && len(c.resized) > 0 {
		return
	}

	var lines []string
	for _, originalLine := range c.original {
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
			case overallLineWidth == width:
				lines = append(lines, strings.Join(line, space))

				line = []string{}
				lineWidth = 0
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

	for n := range lines {
		var (
			spanLeft  int
			spanRight int
		)

		lineWidth := runewidth.StringWidth(lines[n])

		switch c.options.Align {
		case Left:
			spanRight = width - lineWidth
		case Center:
			span := width - lineWidth

			spanLeft = int(math.Floor(float64(span) / float64(2)))
			spanRight = span - spanLeft
		case Right:
			spanLeft = width - lineWidth
		}

		lines[n] = fmt.Sprintf("%s%s%s",
			strings.Repeat(space, spanLeft),
			lines[n],
			strings.Repeat(space, spanRight))
	}

	c.resized = lines
	c.width = width
	c.height = len(c.resized)
}

func (c *Col) resizeHeight(height int) {
	if height <= c.height {
		return
	}

	span := height - c.height
	for i := 0; i < span; i++ {
		c.resized = append(c.resized, strings.Repeat(space, c.width))
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
		options:  o,
		original: lines,
		minWidth: minWidth,
		width:    width,
		height:   len(lines),
	}

	if o.Width > 0 {
		width = o.Width
	}

	col.resizeWidth(width)

	return col
}
