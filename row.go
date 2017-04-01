package simpletable

import (
	"fmt"
	"strings"
)

// tblRow is a meta table row
type tblRow struct {
	Cells []cellInterface
	Table *Table
}

// toString returns row contents as a toString
func (r *tblRow) toString() string {
	s := []string{}

	for _, c := range r.Cells {
		s = append(s, c.toString())
	}

	c := strings.Join(s, fmt.Sprintf(" %s ", r.Table.style.Cell))

	switch r.Cells[0].(type) {
	case *dividerCell:
		return c
	}

	return fmt.Sprintf(" %s ", c)
}

// isDivider check if a row is divider (if first row cell is divider - row divider to)
func (r *tblRow) isDivider() bool {
	switch r.Cells[0].(type) {
	case *dividerCell:
		return true
	}

	return false
}
