package simpletable

import "sort"

// Header is table header
type Header struct {
	Cells []*Cell
}

// Body is table body
type Body struct {
	Cells    [][]*Cell
	sortBy   int // index of Cells to sort Body
	sortFunc func(i, j string) bool
}

// Footer is table footer
type Footer struct {
	Cells []*Cell
}

func (b *Body) Len() int {
	return len(b.Cells)
}

func (b *Body) Less(i, j int) bool {
	if b.sortFunc != nil {
		return b.sortFunc(b.Cells[i][b.sortBy].Text, b.Cells[j][b.sortBy].Text)
	}
	return b.Cells[i][b.sortBy].Text < b.Cells[j][b.sortBy].Text
}

func (b *Body) Swap(i, j int) {
	b.Cells[i], b.Cells[j] = b.Cells[j], b.Cells[i]
}

// SortByField sorts table.Body by given Cells field index using lexical order
func (b *Body) SortByField(index int) {
	b.sortBy = index
	b.sortFunc = nil
	sort.Sort(b)
}

// SortByFieldWithFunc sorts table.Body by given Cells field index using Less func
func (b *Body) SortByFieldWithFunc(index int, Less func(i, j string) bool) {
	b.sortBy = index
	b.sortFunc = Less
	sort.Sort(b)
}
