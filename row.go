package simpletable

import "strings"

type Row struct {
	Cells []Cell
}

func (r *Row) String() string {
	s := []string{}

	for _, c := range r.Cells {
		s = append(s, c.String())
	}

	return strings.Join(s, " | ")
}
