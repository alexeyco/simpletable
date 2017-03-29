package simpletable

type Header struct {
	Cells []Cell
}

func (h *Header) String() string {
	return ""
}

type Body struct {
	Cells [][]Cell
}

func (b *Body) String() string {
	return ""
}

type Footer struct {
	Cells []Cell
}

func (h *Footer) String() string {
	return ""
}
