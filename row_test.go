package simpletable_test

import (
	"strings"
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRowFormatting(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{{{Text: "12345"}}, {{Text: "123456789"}}}

	output := table.String()
	assert.Contains(t, output, "12345")
	assert.Contains(t, output, "123456789")
	assertEqualLineWidths(t, output)
}

func TestRowDividerWithHeader(t *testing.T) {
	table := simpletable.New()
	table.Header.Cells = []*simpletable.Cell{{Text: "a"}}
	table.Body.Cells = [][]*simpletable.Cell{{{Text: "b"}}}

	assert.Contains(t, table.String(), "\n+---+\n")
}

func TestRowWithoutHeaderHasNoDivider(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{{{Text: "a"}}}

	lines := strings.Split(table.String(), "\n")
	require.Len(t, lines, 3)
}
