package simpletable

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

// stripAnsiEscapeExp is a regular expression to clean ANSI Control sequences
// feat https://stackoverflow.com/questions/14693701/how-can-i-remove-the-ansi-escape-sequences-from-a-string-in-python#33925425
var stripAnsiEscapeExp = regexp.MustCompile(`(\x9B|\x1B\[)[0-?]*[ -/]*[@-~]`)

type Col struct {
	text    string
	options Options
	width   uint
	height  uint
	lines   []string
}

func (c Col) Apply(options ...Option) {
	for _, option := range options {
		option(&c.options)
	}

	c.width = 0
	c.height = 0
	c.lines = []string{}

	lines := strings.Split(c.text, "\n")
	for _, line := range lines {
		resizedLines, w := resizeLine(line, c.options.Width)
		if w > c.width {
			c.width = w
		}

		c.lines = append(c.lines, resizedLines...)
	}

	c.height = uint(len(c.lines))
}

func (c Col) Width() uint {
	return c.width
}

func (c Col) Height() uint {
	return c.height
}

func (c Col) Lines() []string {
	return c.lines
}

func Column(text string, options ...Option) Col {
	c := Col{
		text: text,
		options: Options{
			Span: 1,
		},
	}

	c.Apply(options...)

	return c
}

func resizeLine(line string, width uint) ([]string, uint) {
	if width == 0 {
		return []string{line}, uint(runewidth.StringWidth(line))
	}

	w := uint(runewidth.StringWidth(line))
	if w <= width {
		return []string{line}, width
	}

	return []string{line}, 0
}

func alignLines(lines []string, align uint8) []string {
	return lines
}
