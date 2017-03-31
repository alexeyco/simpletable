package simpletable

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"math"
)

type Table struct {
	Header   *Header
	Body     *Body
	Footer   *Footer
	style    *Style
	rows     []*Row
	columns  []*Column
	spanned  []*TextCell
	dividers []*Divider
}

func (t *Table) SetStyle(style *Style) {
	t.style = style
}

func (t *Table) String() string {
	t.refresh()

	// TODO: protect against wrong spans

	t.prepareRows()
	t.prepareColumns()
	t.resizeColumns()

	s := []string{}

	b := t.borderTop()
	if b != "" {
		s = append(s, b)
	}

	for _, r := range t.rows {
		s = append(s, t.borderLeftRight(r.String(), r.IsDivider()))
	}

	b = t.borderBottom()
	if b != "" {
		s = append(s, b)
	}

	return strings.Join(s, "\n")
}

func (t *Table) Print() {
	fmt.Print(t.String())
}

func (t *Table) Println() {
	fmt.Println(t.String())
}

func (t *Table) refresh() {
	t.rows = []*Row{}
	t.columns = []*Column{}
	t.spanned = []*TextCell{}
	t.dividers = []*Divider{}
}

func (t *Table) borderTop() string {
	s := t.style.Border
	return t.line(s.TopLeft, s.Top, s.TopRight, s.TopIntersection)
}

func (t *Table) borderBottom() string {
	s := t.style.Border
	return t.line(s.BottomLeft, s.Bottom, s.BottomRight, s.BottomIntersection)
}

func (t *Table) borderLeftRight(s string, d bool) string {
	if d {
		return s
	}

	return fmt.Sprintf("%s%s%s", t.style.Border.Left, s, t.style.Border.Right)
}

func (t *Table) line(l, c, r, i string) string {
	b := []string{}
	for _, col := range t.columns {
		b = append(b, strings.Repeat(c, col.Width()+2))
	}

	return fmt.Sprintf("%s%s%s", l, strings.Join(b, i), r)
}

func (t *Table) prepareRows() {
	hlen := len(t.Header.Cells)
	if hlen > 0 {
		t.rows = append(t.rows, &Row{
			Cells: t.Header.Cells,
			Table: t,
		})

		d := &Divider{
			Span: hlen,
		}

		t.rows = append(t.rows, &Row{
			Cells: []Cell{
				d,
			},
			Table: t,
		})

		t.dividers = append(t.dividers, d)
	}

	for _, r := range t.Body.Cells {
		t.rows = append(t.rows, &Row{
			Cells: r,
			Table: t,
		})
	}

	flen := len(t.Footer.Cells)
	if flen > 0 {
		d := &Divider{
			Span: hlen,
		}

		t.rows = append(t.rows, &Row{
			Cells: []Cell{
				d,
			},
			Table: t,
		})

		t.dividers = append(t.dividers, d)

		t.rows = append(t.rows, &Row{
			Cells: t.Footer.Cells,
			Table: t,
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
			var p Cell
			var tc *TextCell

			switch v := c.(type) {
			case *TextCell:
				span = v.Span
				p = v
				tc = v
			case *Divider:
				span = v.Span
				p = v
			}

			if span > 1 {
				empty := []*EmptyCell{}

				for i := 1; i < span; i++ {
					empty = append(empty, &EmptyCell{
						parent: p,
					})
				}

				for _, c := range empty {
					row = append(row, c)
				}

				if tc != nil {
					t.spanned = append(t.spanned, tc)
				}

				switch v := c.(type) {
				case *TextCell:
					v.children = empty
				case *Divider:
					v.children = empty
				}
			}
		}

		m = append(m, row)
	}

	m = t.transposeCells(m)
	for _, r := range m {
		c := &Column{
			Cells: r,
			Table: t,
		}

		for _, cell := range c.Cells {
			cell.SetColumn(c)
		}

		t.columns = append(t.columns, c)
	}
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

func (t *Table) resizeColumns() {
	for _, c := range t.columns {
		c.Resize()
	}

	for _, c := range t.spanned {
		c.Resize()
	}

	for _, d := range t.dividers {
		s := t.size()
		d.SetWidth(s)
	}
}

func (t *Table) incrementColumns(c []*Column, length int) {
	sizes := t.carve(length, len(c))

	for i, col := range c {
		col.IncrementWidth(sizes[i])
	}
}

func (t *Table) carve(length, parts int) []int {
	r := []int{}
	step := int(math.Floor(float64(length)/float64(parts)))
	if step * parts != length {
		step++
	}

	for i := 0; i < parts; i++ {
		var n int
		if length < step {
			n = length
		} else {
			n = step
		}

		length -= n
		r = append(r, n)
	}

	return r
}

// size returns table content size.
func (t *Table) size() int {
	return utf8.RuneCountInString(t.rows[0].String())
}

// New is a Table constructor. It loads struct data, ready to be manipulated.
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
		rows:     []*Row{},
		columns:  []*Column{},
		spanned:  []*TextCell{},
		dividers: []*Divider{},
	}
}
