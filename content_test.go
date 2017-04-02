package simpletable

import (
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

func testContentEqual(o, s []string) bool {
	for i, v := range o {
		if s[i] != v {
			return false
		}
	}

	return true
}
