package simpletable_test

import (
	"strings"
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCellRendering(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{
		{{Text: "12345", Align: simpletable.AlignCenter}, {Text: "short"}},
		{{Text: "123456789"}, {Text: "row"}},
	}

	output := table.String()
	assert.Contains(t, output, "12345")
	assert.Contains(t, output, "123456789")
	assert.Contains(t, output, "short")
}

func TestSpannedCellRendering(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{
		{{Span: 2, Text: "spanned cell"}},
		{{Text: "left"}, {Text: "right"}},
	}

	output := table.String()
	assert.Contains(t, output, "spanned cell")
	assertEqualLineWidths(t, output)
}

func TestDividerRendering(t *testing.T) {
	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{{Text: "AAA"}, {Text: "BBB"}}
	table.Body.Cells = [][]*simpletable.Cell{{{Text: "CCC"}, {Text: "DDD"}}}

	output := table.String()
	assert.Contains(t, output, "+-----+-----+")
}

func TestMultilineCellHeight(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{{
		{Text: "first\nsecond"},
		{Text: "single"},
	}}

	lines := strings.Split(table.String(), "\n")
	require.Len(t, lines, 4)
	assert.Contains(t, lines[2], "second")
}

func assertEqualLineWidths(t *testing.T, output string) {
	t.Helper()

	lines := strings.Split(output, "\n")
	width := len(lines[0])
	for i, line := range lines[1:] {
		assert.Lenf(t, line, width, "line %d has unexpected width: %q", i+2, line)
	}
}
