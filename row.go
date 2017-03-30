package simpletable

import (
	"fmt"
	"strings"
)

type Row struct {
	Cells []Cell
	Table *Table
}

func (r *Row) String() string {
	s := []string{}

	for _, c := range r.Cells {
		s = append(s, c.String())
	}

	c := strings.Join(s, fmt.Sprintf(" %s ", r.Table.style.Cell))

	switch r.Cells[0].(type) {
	case *Divider:
		return c
	}

	return fmt.Sprintf(" %s ", c)
}

func (r *Row) IsDivider() bool {
	switch r.Cells[0].(type) {
	case *Divider:
		return true
	}

	return false
}
