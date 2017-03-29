package simpletable

import (
	"strings"
	"fmt"
)

type Row struct {
	columns []*Column
}

func (r *Row) AddColumn(c *Column) {
	r.columns = append(r.columns, c)
}

func (r *Row) Columns() []*Column {
	return r.columns
}

func (r *Row) Len() int {
	return len(r.columns)
}

func (r *Row) Height() int {
	height := 0
	for _, c := range r.columns {
		h := c.Content().Height()

		if height < h {
			height = h
		}
	}

	return height
}

func (r *Row) String(widths ...int) string {
	d := [][]string{}

	for n, c := range r.columns {
		d = append(d, c.Content().StringSlice(widths[n]))
	}

	s := []string{}
	for l := 0; l < r.Height(); l++ {
		t := []string{}

		for _, col := range d {
			t = append(t, col[l])
		}

		s = append(s, strings.Join(t, " │ "))
	}

	return fmt.Sprintf(" %s ", strings.Join(s, "\n"))
}

func newRow() *Row {
	return &Row{
		columns: []*Column{},
	}
}
