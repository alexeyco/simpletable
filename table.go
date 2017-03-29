package simpletable

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Table struct {
	Header  *Header
	Body    *Body
	Footer  *Footer
	style   *Style
	rows    []*Row
	columns []*Column
}

func (t *Table) SetStyle(style *Style) {
	t.style = style
}

func (t *Table) String() string {
	t.prepare()

	return ""
}

func (t *Table) Print() {
	fmt.Print(t.String())
}

func (t *Table) Println() {
	fmt.Println(t.String())
}

func (t *Table) prepare() {
	t.prepareRows()
	t.prepareColumns()
}

func (t *Table) prepareRows() {
	hlen := len(t.Header.Cells)
	if hlen > 0 {
		t.rows = append(t.rows, &Row{
			Cells: t.Header.Cells,
		})

		t.rows = append(t.rows, &Row{
			Cells: []Cell{
				&Divider{
					Span: hlen,
				},
			},
		})
	}

	for _, r := range t.Body.Cells {
		t.rows = append(t.rows, &Row{
			Cells: r,
		})
	}

	flen := len(t.Footer.Cells)
	if flen > 0 {
		t.rows = append(t.rows, &Row{
			Cells: []Cell{
				&Divider{
					Span: flen,
				},
			},
		})

		t.rows = append(t.rows, &Row{
			Cells: t.Footer.Cells,
		})
	}
}

func (t *Table) prepareColumns() {
	m := [][]Cell{}

	for _, r := range t.rows {
		row := []Cell{}

		for _, c := range r.Cells {
			row = append(row, c)
			span := 0
			var p Parent

			switch v := c.(type) {
			case *TextCell:
				span = v.Span
				p = v
			case *Divider:
				span = v.Span
				p = v
			}

			if span > 1 {
				for i := 1; i < span; i++ {
					row = append(row, &EmptyCell{
						parent: p,
					})
				}
			}
		}

		m = append(m, row)
	}

	m = t.transposeCells(m)
	spew.Dump(m)
}

func (t *Table) transposeCells(i [][]Cell) [][]Cell {
	r := [][]Cell{}

	for x := 0; x < len(i[0]); x++ {
		r = append(r, make([]Cell, len(i)))
	}

	for x, row := range i {
		for y, c := range row {
			r[y][x] = c
		}
	}

	return r
}

func New() *Table {
	return &Table{
		style: StyleDefault,
		Header: &Header{
			Cells: []Cell{},
		},
		Body: &Body{
			Cells: [][]Cell{},
		},
		Footer: &Footer{
			Cells: []Cell{},
		},
		rows:    []*Row{},
		columns: []*Column{},
	}
}
