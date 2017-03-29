package simpletable

import (
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	header *Row
	rows   []*Row
}

func (t *Table) AddRow(columns ...string) error {
	if len(columns) != t.header.Len() {
		err := fmt.Sprintf("Row [%s] must contain %d columns", strings.Join(columns, ", "), t.header.Len())
		return errors.New(err)
	}

	r := newRow()
	for _, c := range columns {
		r.AddColumn(&Column{
			content: newContent(c),
		})
	}

	t.rows = append(t.rows, r)
	return nil
}

func (t *Table) String() string {
	c := []string{}
	widths := t.widths()

	c = append(c, t.header.String(widths...))
	c = append(c, t.line(widths...))

	for _, col := range t.rows {
		c = append(c, col.String(widths...))
	}

	return strings.Join(c, "\n")
}

func (t *Table) Print() {
	fmt.Println(t.String())
}

func (t *Table) widths() []int {
	w := []int{}
	for _, c := range t.header.Columns() {
		w = append(w, c.Content().Width())
	}

	for _, r := range t.rows {
		for i, c := range r.Columns() {
			l := c.Content().Width()
			if w[i] < l {
				w[i] = l
			}
		}
	}

	return w
}

func (t *Table) line(widths ...int) string {
	s := []string{}

	for _, w := range widths {
		s = append(s, strings.Repeat("═", w))
	}

	return fmt.Sprintf("═%s═", strings.Join(s, "═╪═"))
}

func New(headers ...string) *Table {
	header := &Row{}

	for _, h := range headers {
		header.AddColumn(&Column{
			content: newContent(h),
		})
	}

	return &Table{
		header: header,
		rows:   []*Row{},
	}
}
