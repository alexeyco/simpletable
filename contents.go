package simpletable

type Header struct {
	Cells []Cell
}

type Body struct {
	Cells [][]Cell
}

type Footer struct {
	Cells []Cell
}
