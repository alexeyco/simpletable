package simpletable

// Line table line.
type Line interface {
}

type row struct {
	columns []*Col
}

// Row table row.
func Row(columns ...*Col) Line {
	return row{
		columns: columns,
	}
}

type divider struct {
}

// Divider between rows.
func Divider() Line {
	return divider{}
}
