package simpletable_test

import (
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ interface {
	Print()
	Println()
} = (*simpletable.Table)(nil)

func TestNew(t *testing.T) {
	require.NotNil(t, simpletable.New())
}

func TestTable_Carve(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{
		{{Span: 3, Text: "abcdefghijklmn"}},
		{{Text: "a"}, {Text: "b"}, {Text: "c"}},
	}

	output := table.String()
	assert.Contains(t, output, "+-----+-----+----+")
	assert.Contains(t, output, "| a   | b   | c  |")
}
