package simpletable

import (
	"strings"

	"github.com/alexeyco/simpletable/style"
)

// Table structure.
type Table struct {
	style style.Style
	lines []Line
}

// SetStyle table style setter.
func (t *Table) SetStyle(s style.Style) *Table {
	t.style = s

	return t
}

// Append line to table.
func (t *Table) Append(line Line) *Table {
	t.lines = append(t.lines, line)

	return t
}

// String returns table as a string.
func (t *Table) String() string {
	var lines []string

	return strings.Join(lines, "\n")
}

// New returns new table instance.
func New() *Table {
	return &Table{
		style: style.StyleDefault,
	}
}
