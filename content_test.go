package simpletable

import (
	"fmt"
	"testing"
)

var testContent = `1
22
333`

func TestContent_Width(t *testing.T) {
	c := newContent(testContent)

	if c.width() != 3 {
		t.Error("Wrong content width")
	}
}

func TestContent_Height(t *testing.T) {
	c := newContent(testContent)

	if c.height() != 3 {
		t.Error("Wrong content height")
	}
}

func TestContent_SetHeight(t *testing.T) {
	c := newContent(testContent)

	c.setHeight(5)
	if c.height() != 5 {
		t.Error("Wrong content height")
	}
}

func TestContent_Lines(t *testing.T) {
	c := newContent(testContent)
	c.setHeight(5)

	left := []string{
		"1  ",
		"22 ",
		"333",
		"   ",
		"   ",
	}

	center := []string{
		" 1 ",
		"22 ",
		"333",
		"   ",
		"   ",
	}

	right := []string{
		"  1",
		" 22",
		"333",
		"   ",
		"   ",
	}

	if !testContentEqual(c.lines(AlignLeft), left) {
		t.Error("Wrong lines rendering (align: left)")
	}

	if !testContentEqual(c.lines(AlignCenter), center) {
		t.Error("Wrong lines rendering (align: center)")
	}

	if !testContentEqual(c.lines(AlignRight), right) {
		t.Error("Wrong lines rendering (align: right)")
	}
}

func TestContent_StripAnsiEscape(t *testing.T) {
	line := "\t\u001b[0;35mBlabla\u001b[0m                                  \u001b[0;36m172.18.0.2\u001b[0m"

	if stripAnsiEscape(line) != "\tBlabla                                  172.18.0.2" {
		t.Error("Wrong ANSI escape sequences stripping")
	}
}

func TestContent_RealLength(t *testing.T) {
	colorDefault := "\x1b[39m"
	colorRed := "\x1b[91m"

	plainString := "I am string"
	colorizedString := fmt.Sprintf("%s%s%s", colorRed, plainString, colorDefault)

	if realLength(plainString) != realLength(colorizedString) {
		t.Error("Wrong real string length calculation")
	}
}

func testContentEqual(o, s []string) bool {
	for i, v := range o {
		if s[i] != v {
			return false
		}
	}

	return true
}
