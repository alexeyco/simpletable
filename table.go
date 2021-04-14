package simpletable

import "strings"

type info struct {
	length uint
}

type Table struct {
	style Style
	info  info
	rows  []interface{}
}

func (t *Table) Style(s Style) *Table {
	t.style = s

	return t
}

func (t *Table) Row(columns ...Col) *Table {
	var l uint
	for _, c := range columns {
		l += c.options.Span
	}

	r := row{
		columns: columns,
		length:  l,
	}

	if l > t.info.length {
		t.info.length = l
		t.pad(l)
	} else if l < t.info.length {
		r.pad(t.info.length)
	}

	t.rows = append(t.rows, r)

	return t
}

func (t *Table) Divider() *Table {
	t.rows = append(t.rows, divider{})

	return t
}

func (t *Table) pad(pad uint) {
	t.iterateRows(func(_ int, r row) {
		r.pad(pad)
	})
}

func (t *Table) iterateRows(f func(n int, r row)) {
	t.iterate(func(n int, v interface{}) {
		if r, ok := v.(row); ok {
			f(n, r)
		}
	})
}

func (t *Table) iterate(f func(n int, v interface{})) {
	for n, v := range t.rows {
		f(n, v)
	}
}

func (t *Table) line(left, intersection, right rune) string {
	return ""
}

func (t *Table) String() string {
	var lines []string

	//lines = append(lines, t.line())

	t.iterate(func(_ int, v interface{}) {
		switch v.(type) {
		case row:
			//
		case divider:
			//
		}
	})

	//lines = append(lines, t.line())

	return strings.Join(lines, "\n")
}

func New() *Table {
	return &Table{
		style: StyleDefault,
		info:  info{},
	}
}
