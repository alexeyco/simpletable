package simpletable_test

import (
	"reflect"
	"strings"
	"testing"

	st "github.com/alexeyco/simpletable"
)

func TestColumn(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name    string
		column  st.Col
		options []st.Option
		width   uint
		height  uint
		lines   []string
	}{
		{
			name:   "SingleLine",
			column: st.Column("Foo bar"),
			width:  7,
			height: 1,
			lines:  []string{"Foo bar"},
		},
		{
			name:   "MultiLine",
			column: st.Column("Foo\nbar"),
			width:  3,
			height: 2,
			lines:  []string{"Foo", "bar"},
		},
		{
			name:    "SingleLineResize",
			column:  st.Column("Lorem ipsum dolor"),
			options: []st.Option{st.Width(11)},
			width:   11,
			height:  2,
			lines:   []string{"Lorem ipsum", "dolor"},
		},
		{
			name:    "MultiLineAlignCenter",
			column:  st.Column("Foo bar\nfizz"),
			options: []st.Option{st.Align(st.Right)},
			width:   11,
			height:  2,
			lines:   []string{"Foo bar", "  fizz "},
		},
		{
			name:    "MultiLineAlignRight",
			column:  st.Column("Foo bar\nfizz"),
			options: []st.Option{st.Align(st.Right)},
			width:   11,
			height:  2,
			lines:   []string{"Foo bar", "   fizz"},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			col := testDatum.column
			col.Apply(testDatum.options...)

			expectedWidth := testDatum.width
			actualWidth := col.Width()

			if expectedWidth != actualWidth {
				t.Errorf(`Width should be %d, %d given`, expectedWidth, actualWidth)
			}

			expectedHeight := testDatum.height
			actualHeight := col.Height()

			if expectedHeight != actualHeight {
				t.Errorf(`Height should be %d, %d given`, expectedHeight, actualHeight)
			}

			expectedLines := testDatum.lines
			actualLines := col.Lines()

			if !reflect.DeepEqual(expectedLines, actualLines) {
				t.Errorf(`Lines should be ["%s"], "["%s"]" given`,
					strings.Join(expectedLines, `", "`),
					strings.Join(actualLines, `", "`))
			}
		})
	}
}
