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
			column: st.Column("Lorem ipsum dolor sit amet"),
			width:  26,
		},
		{
			name:   "SingleLineTrim",
			column: st.Column(" Lorem ipsum dolor sit amet "),
			width:  26,
		},
		{
			name:   "DoubleLine",
			column: st.Column("Lorem\nipsum dolor sit amet"),
			width:  20,
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column(" Lorem \n ipsum dolor sit amet "),
			width:  20,
		},
		{
			name:   "MultiLine",
			column: st.Column("Lorem\nipsum\ndolor\nsit\namet"),
			width:  5,
		},
		{
			name:   "MultiLineTrim",
			column: st.Column(" Lorem \n ipsum \n dolor \n sit \n amet "),
			width:  5,
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
			column: st.Column("Lorem ipsum dolor sit amet"),
			height: 1,
		},
		{
			name:   "SingleLineTrim",
			column: st.Column(" Lorem ipsum dolor sit amet "),
			height: 1,
		},
		{
			name:   "DoubleLine",
			column: st.Column("Lorem\nipsum dolor sit amet"),
			height: 2,
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column(" Lorem \n ipsum dolor sit amet "),
			height: 2,
		},
		{
			name:   "MultiLine",
			column: st.Column("Lorem\nipsum\ndolor\nsit\namet"),
			height: 5,
		},
		{
			name:   "MultiLineTrim",
			column: st.Column(" Lorem \n ipsum \n dolor \n sit \n amet "),
			height: 5,
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
			column: st.Column("Lorem ipsum dolor sit amet"),
			lines:  []string{"Lorem ipsum dolor sit amet"},
		},
		{
			name:   "SingleLineTrim",
			column: st.Column(" Lorem ipsum dolor sit amet "),
			lines:  []string{"Lorem ipsum dolor sit amet"},
		},
		{
			name:   "SingleLineColored",
			column: st.Column("\x1b[38;5;140mLorem ipsum dolor sit amet\x1b[0m"),
			lines:  []string{"Lorem ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLine",
			column: st.Column("Lorem\nipsum dolor sit amet"),
			lines:  []string{"Lorem               ", "ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLineTrim",
			column: st.Column(" Lorem \n ipsum dolor sit amet "),
			lines:  []string{"Lorem               ", "ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLineAlignRight",
			column: st.Column("Lorem\nipsum dolor sit amet", st.Align(st.Center)),
			lines:  []string{"       Lorem        ", "ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLineTrimAlignCenter",
			column: st.Column(" Lorem \n ipsum dolor sit amet ", st.Align(st.Center)),
			lines:  []string{"       Lorem        ", "ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLineAlignRight",
			column: st.Column("Lorem\nipsum dolor sit amet", st.Align(st.Right)),
			lines:  []string{"               Lorem", "ipsum dolor sit amet"},
		},
		{
			name:   "DoubleLineTrimAlignRight",
			column: st.Column(" Lorem \n ipsum dolor sit amet ", st.Align(st.Right)),
			lines:  []string{"               Lorem", "ipsum dolor sit amet"},
		},
		{
			name:   "MultiLine",
			column: st.Column("Lorem\nipsum\ndolor\nsit\namet"),
			lines:  []string{"Lorem", "ipsum", "dolor", "sit  ", "amet "},
		},
		{
			name:   "MultiLineTrim",
			column: st.Column(" Lorem \n ipsum \n dolor \n sit \n amet "),
			lines:  []string{"Lorem", "ipsum", "dolor", "sit  ", "amet "},
		},
		{
			name:   "MultiLineAlignCenter",
			column: st.Column("Lorem\nipsum\ndolor\nsit\namet", st.Align(st.Center)),
			lines:  []string{"Lorem", "ipsum", "dolor", " sit ", "amet "},
		},
		{
			name:   "MultiLineTrimAlignCenter",
			column: st.Column(" Lorem \n ipsum \n dolor \n sit \n amet ", st.Align(st.Center)),
			lines:  []string{"Lorem", "ipsum", "dolor", " sit ", "amet "},
		},
		{
			name:   "MultiLineAlignRight",
			column: st.Column("Lorem\nipsum\ndolor\nsit\namet", st.Align(st.Right)),
			lines:  []string{"Lorem", "ipsum", "dolor", "  sit", " amet"},
		},
		{
			name:   "MultiLineTrimAlignRight",
			column: st.Column(" Lorem \n ipsum \n dolor \n sit \n amet ", st.Align(st.Right)),
			lines:  []string{"Lorem", "ipsum", "dolor", "  sit", " amet"},
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
		name         string
		column       *st.Col
		resizeWidth  int
		resizeHeight int
		lines        []string
		width        int
		height       int
	}{
		{
			name:        "SingleLineResizeWidth26",
			column:      st.Column("Lorem ipsum dolor sit amet"),
			resizeWidth: 26,
			lines:       []string{"Lorem ipsum dolor sit amet"},
			width:       26,
			height:      1,
		},
		{
			name:        "SingleLineResizeWidth14",
			column:      st.Column("Lorem ipsum dolor sit amet"),
			resizeWidth: 14,
			lines:       []string{"Lorem ipsum   ", "dolor sit amet"},
			width:       14,
			height:      2,
		},
		{
			name:        "SingleLineResizeWidth14AlignCenter",
			column:      st.Column("Lorem ipsum dolor sit amet", st.Align(st.Center)),
			resizeWidth: 14,
			lines:       []string{" Lorem ipsum  ", "dolor sit amet"},
			width:       14,
			height:      2,
		},
		{
			name:        "SingleLineResizeWidth14AlignRight",
			column:      st.Column("Lorem ipsum dolor sit amet", st.Align(st.Right)),
			resizeWidth: 14,
			lines:       []string{"   Lorem ipsum", "dolor sit amet"},
			width:       14,
			height:      2,
		},
		{
			name:         "SingleLineResizeWidth14Height3",
			column:       st.Column("Lorem ipsum dolor sit amet"),
			resizeWidth:  14,
			resizeHeight: 3,
			lines:        []string{"Lorem ipsum   ", "dolor sit amet", "              "},
			width:        14,
			height:       3,
		},
		{
			name:        "SingleLineResizeWidth28",
			column:      st.Column("Lorem ipsum dolor sit amet  "),
			resizeWidth: 28,
			lines:       []string{"Lorem ipsum dolor sit amet  "},
			width:       28,
			height:      1,
		},
		{
			name:        "SingleLineResizeWidth1",
			column:      st.Column("Lorem ipsum dolor sit amet  "),
			resizeWidth: 1,
			lines:       []string{"Lorem", "ipsum", "dolor", "sit  ", "amet "},
			width:       5,
			height:      5,
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			testDatum.column.Resize(testDatum.resizeWidth, testDatum.resizeHeight)

			assert.Equal(t, testDatum.lines, testDatum.column.Lines())
			assert.Equal(t, testDatum.width, testDatum.column.Width())
			assert.Equal(t, testDatum.height, testDatum.column.Height())
		})
	}
}
