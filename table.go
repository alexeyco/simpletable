package simpletable

import (
	"fmt"
	"strings"
)

type Header struct {
	Row
	Columns []*Column
}

type Body struct {
	Rows []*Row
}

func (b *Body) String() string {
	s := []string{}

	for _, r := range b.Rows {
		s = append(s, r.String())
	}

	return strings.Join(s, "\n")
}

type Footer struct {
	Row
	Columns []*Column
}

type Table struct {
	style  *Style
	widths []int
	Header *Header
	Body   *Body
	Footer *Footer
}

func (t *Table) SetStyle(style *Style) {
	t.style = style
}

func (t *Table) String() string {
	t.resize()

	s := []string{}

	h := t.Header.String()
	if h != "" {
		s = append(s, h)
	}

	b := t.Body.String()
	if b != "" {
		s = append(s, b)
	}

	f := t.Footer.String()
	if f != "" {
		s = append(s, f)
	}

	return strings.Join(s, "\n")
}

func (t *Table) Print() {
	fmt.Print(t.String())
}

func (t *Table) Println() {
	fmt.Println(t.String())
}

func (t *Table) rows() [][]*Column {
	rows := [][]*Column{}

	if len(t.Header.Columns) > 0 {
		rows = append(rows, t.Header.Columns)
	}

	for _, r := range t.Body.Rows {
		rows = append(rows, r.Columns)
	}

	if len(t.Footer.Columns) > 0 {
		rows = append(rows, t.Footer.Columns)
	}

	return rows
}

func (t *Table) resize() {
	rows := t.rows()
	widths := []int{}

	for _, r := range rows {
		for n, c := range r {
			if c.Span < 1 {
				c.Span = 1
			}

			w := c.Width()
			if len(widths) >= n + 1 {
				if widths[n] < w {
					widths[n] = w
				}
			} else {
				widths = append(widths, w)
			}
		}
	}

	t.widths = widths
}

func New() *Table {
	return &Table{
		style: StyleDefault,
		widths: []int{},
		Header: &Header{
			Columns: []*Column{},
		},
		Body: &Body{
			Rows: []*Row{},
		},
		Footer: &Footer{
			Columns: []*Column{},
		},
	}
}
