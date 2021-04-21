package simpletable_test

import (
	"testing"

	st "github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
)

func TestCol_Width(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name   string
		column *st.Col
		width  int
	}{
		{
			name:   "SingleLine",
			column: st.Column("Foo Bar"),
			width:  7,
		},
		{
			name:   "SingleLineTrim",
			column: st.Column("   Foo Bar   "),
			width:  7,
		},
		{
			name:   "SingleLineResize",
			column: st.Column("Foo Bar", st.Width(6)),
			width:  6,
		},
		{
			name:   "SingleLineTrimResize",
			column: st.Column("   Foo Bar   ", st.Width(6)),
			width:  6,
		},
		{
			name:   "SingleLineOverResize",
			column: st.Column("Foo Bar", st.Width(1)),
			width:  3,
		},
		{
			name:   "SingleLineTrimOverResize",
			column: st.Column("   Foo Bar   ", st.Width(1)),
			width:  3,
		},
		{
			name:   "DoubleLine",
			column: st.Column("Foo\nBar"),
			width:  3,
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column("  Foo  \n  Bar  "),
			width:  3,
		},
		{
			name:   "DoubleLineResize",
			column: st.Column("Foo\nBar", st.Width(6)),
			width:  6,
		},
		{
			name:   "DoubleLineTrimResize",
			column: st.Column("  Foo  \n  Bar  ", st.Width(6)),
			width:  6,
		},
		{
			name:   "DoubleLineOverResize",
			column: st.Column("Foo\nBar", st.Width(3)),
			width:  3,
		},
		{
			name:   "DoubleLineTrimOverResize",
			column: st.Column("  Foo  \n  Bar  ", st.Width(1)),
			width:  3,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testDatum.width, testDatum.column.Width())
		})
	}
}

func TestCol_Height(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name   string
		column *st.Col
		height int
	}{
		{
			name:   "SingleLine",
			column: st.Column("Foo Bar"),
			height: 1,
		},
		{
			name:   "SingleLineResize",
			column: st.Column("Foo Bar", st.Width(6)),
			height: 2,
		},
		{
			name:   "SingleLineOverResize",
			column: st.Column("Foo Bar", st.Width(1)),
			height: 2,
		},
		{
			name:   "DoubleLine",
			column: st.Column("Foo Bar\nBaz"),
			height: 2,
		},
		{
			name:   "DoubleLineResize",
			column: st.Column("Foo Bar\nBaz", st.Width(6)),
			height: 3,
		},
		{
			name:   "DoubleLineOverResize",
			column: st.Column("Foo Bar\nBaz", st.Width(1)),
			height: 3,
		},
		{
			name:   "MultiLine",
			column: st.Column("Foo\nBar\nFiz\nBaz"),
			height: 4,
		},
		{
			name:   "MultiLineResize",
			column: st.Column("Foo\nBar\nFiz\nBaz", st.Width(3)),
			height: 4,
		},
		{
			name:   "MultiLineOverResize",
			column: st.Column("Foo\nBar\nFiz\nBaz", st.Width(1)),
			height: 4,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testDatum.height, testDatum.column.Height())
		})
	}
}

func TestCol_Lines(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name   string
		column *st.Col
		lines  []string
	}{
		{
			name:   "SingleLine",
			column: st.Column("Foo Bar"),
			lines:  []string{"Foo Bar"},
		},
		{
			name:   "SingleLineTrim",
			column: st.Column("   Foo Bar   "),
			lines:  []string{"Foo Bar"},
		},
		{
			name:   "DoubleLine",
			column: st.Column("Foo\nBar"),
			lines:  []string{"Foo", "Bar"},
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column("  Foo  \n  Bar  "),
			lines:  []string{"Foo", "Bar"},
		},
		{
			name:   "DoubleLineResized",
			column: st.Column("Foo\nBar", st.Width(12)),
			lines:  []string{"Foo         ", "Bar         "},
		},
		{
			name:   "DoubleLineResizedAlignCenter",
			column: st.Column("Foo\nBar", st.Width(12), st.Align(st.Center)),
			lines:  []string{"    Foo     ", "    Bar     "},
		},
		{
			name:   "DoubleLineResizedAlignRight",
			column: st.Column("Foo\nBar", st.Width(12), st.Align(st.Right)),
			lines:  []string{"         Foo", "         Bar"},
		},
		{
			name:   "MultiLine",
			column: st.Column("Foo\nBar\nFizz\nBuzz"),
			lines:  []string{"Foo ", "Bar ", "Fizz", "Buzz"},
		},
		{
			name:   "MultiLineTrim",
			column: st.Column("  Foo  \n  Bar  \n  Fizz  \n  Buzz  "),
			lines:  []string{"Foo ", "Bar ", "Fizz", "Buzz"},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testDatum.lines, testDatum.column.Lines())
		})
	}
}

func TestCol_Resize(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name   string
		column *st.Col
		resize int
		lines  []string
		width  int
		height int
	}{
		{
			name:   "SingleLine",
			column: st.Column("Foo Bar"),
			lines:  []string{"Foo Bar"},
		},
		{
			name:   "SingleLineTrim",
			column: st.Column("   Foo Bar   "),
			lines:  []string{"Foo Bar"},
		},
		{
			name:   "DoubleLine",
			column: st.Column("Foo\nBar"),
			lines:  []string{"Foo", "Bar"},
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column("  Foo  \n  Bar  "),
			lines:  []string{"Foo", "Bar"},
		},
		{
			name:   "MultiLine",
			column: st.Column("Foo\nBar\nFizz\nBuzz"),
			lines:  []string{"Foo ", "Bar ", "Fizz", "Buzz"},
		},
		{
			name:   "MultiLineTrim",
			column: st.Column("  Foo  \n  Bar  \n  Fizz  \n  Buzz  "),
			lines:  []string{"Foo ", "Bar ", "Fizz", "Buzz"},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testDatum.lines, testDatum.column.Lines())
		})
	}
}
