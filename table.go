package simpletable

import (
	"strings"

	"github.com/alexeyco/simpletable/grid"
)

type Table struct {
	style Style
	sizer *grid.grid
	rows  []interface{}
}

func (t *Table) Style(s Style) *Table {
	t.style = s

	return t
}

func (t *Table) Row(columns ...*Col) *Table {
	var (
		length int
		height int
	)

	for n, column := range columns {
		length += column.options.Span

		if column.Height() > height {
			height = column.Height()
		}
	}

	t.rows = append(t.rows, row{
		columns: columns,
		length:  length,
		height:  height,
	})

	return t
}

func (t *Table) Divider() *Table {
	t.rows = append(t.rows, divider{})

	return t
}

func (t *Table) String() string {
	var lines []string

	return strings.Join(lines, "\n")
}

func New() *Table {
	return &Table{
		style: StyleDefault,
	}
}
