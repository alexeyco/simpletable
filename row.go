package simpletable

import (
	"fmt"
	"strings"
)

type Row struct {
	Cells []Cell
}

func (r *Row) String() string {
	s := []string{}

	for _, c := range r.Cells {
		s = append(s, c.String())
	}

	c := strings.Join(s, " | ")

	switch r.Cells[0].(type) {
	case *Divider:
		return c
	}

	return fmt.Sprintf(" %s ", c)
}
