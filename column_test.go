package simpletable_test

import (
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
)

func TestColumnUsesLongestCell(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{
		{{Text: "123"}},
		{{Text: "12345"}},
	}

	output := table.String()
	assert.Contains(t, output, "| 123   |")
	assert.Contains(t, output, "| 12345 |")
}

func TestSpannedCellExpandsColumns(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{
		{{Span: 2, Text: "1234567890"}},
		{{Text: "12345"}, {Text: "12345"}},
	}

	output := table.String()
	assert.Contains(t, output, "1234567890")
	assertEqualLineWidths(t, output)
}
