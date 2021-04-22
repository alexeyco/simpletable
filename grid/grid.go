package grid

type Grid struct {
	columns map[int]int
	rows    map[int]int
}

func (g *Grid) Width(idx int) int {
	w, ok := g.columns[idx]
	if !ok {
		return 0
	}

	return w
}

func (g *Grid) SetWidth(idx int, width int) {
	w := g.Width(idx)
	if w < width {
		g.columns[idx] = width
	}
}
