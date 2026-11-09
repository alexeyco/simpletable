package simpletable

// Header is a table header.
type Header struct {
	Cells []*Cell
}

// Body is a table body.
type Body struct {
	Cells [][]*Cell
}

// Footer is a table footer.
type Footer struct {
	Cells []*Cell
}
