package simpletable

import (
	"strings"
)

type Row struct {
	Columns []*Column
}

func (r *Row) String() string {
	s := []string{}

	for _, c := range r.Columns {
		s = append(s, c.String())
	}

	return strings.Join(s, " | ")
}
