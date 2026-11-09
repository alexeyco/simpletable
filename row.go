package simpletable

import (
	"fmt"
	"strings"
)

// tblRow is a metadata row used to render the table.
type tblRow struct {
	Cells []cellInterface
	Table *Table
}

// len returns the rendered row length.
func (r *tblRow) len() int {
	l := 0
	for _, c := range r.Cells {
		l += c.width()
	}

	cellCount := len(r.Cells)
	return l + (3 * cellCount) - 1
}

// toStringSlice returns the rendered row contents as a string slice.
func (r *tblRow) toStringSlice() []string {
	l := [][]string{}

	for _, c := range r.Cells {
		l = append(l, c.lines())
	}

	l = r.transpose(l)

	ret := []string{}
	for _, s := range l {
		row := strings.Join(s, fmt.Sprintf(" %s ", r.Table.style.Cell))
		if !r.isDivider() {
			row = fmt.Sprintf(" %s ", row)
		}

		ret = append(ret, row)
	}

	return ret
}

// transpose transposes a slice of string slices.
func (r *tblRow) transpose(s [][]string) [][]string {
	ret := [][]string{}
	l := len(s)

	for x := 0; x < len(s[0]); x++ {
		ret = append(ret, make([]string, l))
	}

	for x, row := range s {
		for y, c := range row {
			ret[y][x] = c
		}
	}

	return ret
}

// resize sets every cell to the height of the tallest cell.
func (r *tblRow) resize() {
	m := 0
	for _, c := range r.Cells {
		h := c.height()
		if h > m {
			m = h
		}
	}

	for _, c := range r.Cells {
		c.setHeight(m)
	}
}

// isDivider reports whether the first cell in the row is a divider.
func (r *tblRow) isDivider() bool {
	_, ok := r.Cells[0].(*dividerCell)
	return ok
}
