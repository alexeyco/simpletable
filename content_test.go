package simpletable_test

import (
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/assert"
)

const testContent = "1\n22\n333"

func TestContentWidthAndHeight(t *testing.T) {
	table := simpletable.New()
	table.Body.Cells = [][]*simpletable.Cell{{{Text: testContent}}, {{Text: "1234"}}}

	output := table.String()
	assert.Contains(t, output, "+------+")
	assert.Contains(t, output, "333")
}

func TestContentAlignment(t *testing.T) {
	tests := []struct {
		name  string
		align int
		want  string
	}{
		{name: "left", align: simpletable.AlignLeft, want: "| a    |"},
		{name: "center", align: simpletable.AlignCenter, want: "|  a   |"},
		{name: "right", align: simpletable.AlignRight, want: "|    a |"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			table := simpletable.New()
			table.Body.Cells = [][]*simpletable.Cell{
				{{Text: "a\nbb", Align: test.align}},
				{{Text: "xxxx"}},
			}

			assert.Contains(t, table.String(), test.want)
		})
	}
}

func TestContentStripsANSIForWidth(t *testing.T) {
	tests := []struct {
		name    string
		colored string
	}{
		{name: "7-bit CSI", colored: "\x1b[91mabc\x1b[0m"},
		{name: "8-bit CSI", colored: "\x9b91mabc\x9b0m"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			table := simpletable.New()
			table.Body.Cells = [][]*simpletable.Cell{{{Text: test.colored}}, {{Text: "1234"}}}

			output := table.String()
			assert.Contains(t, output, "+------+")
			assert.Contains(t, output, test.colored)
		})
	}
}
