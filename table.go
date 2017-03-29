package simpletable

import (
	"errors"
	"fmt"
	"strings"
)

const(
	AlignLeft = "left"
	AlignCenter = "right"
	AlignRight = "center"
)

type Table struct {
	style  *Style
	align  []string
	header *Row
	rows   []*Row
}

func (t *Table) SetStyle(style *Style) {
	t.style = style
}

func (t *Table) SetAlign(align ...string) {
	t.align = align
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

	align := []string{}
	for i := 0; i < t.header.Len(); i++ {
		align = append(align, AlignCenter)
	}

	c = append(c, t.header.SetAlign(align...).String(widths...))
	c = append(c, t.line("╪", widths...))

	for _, col := range t.rows {
		c = append(c, col.SetAlign(t.align...).String(widths...))
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

func (t *Table) line(sep string, widths ...int) string {
	s := []string{}

	for _, w := range widths {
		s = append(s, strings.Repeat("═", w))
	}

	return fmt.Sprintf("═%s═", strings.Join(s, fmt.Sprintf("═%s═", sep)))
}

func New(headers ...string) *Table {
	header := &Row{}
	align := []string{}

	for _, h := range headers {
		header.AddColumn(&Column{
			content: newContent(h),
		})

		align = append(align, AlignLeft)
	}

	header.Capitalize()

	return &Table{
		style:  StyleSimpleMinimalistic,
		align:  align,
		header: header,
		rows:   []*Row{},
	}
}
