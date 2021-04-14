package simpletable

type row struct {
	columns []Col
	length  uint
}

func (r row) pad(pad uint) {
	if r.length >= pad || len(r.columns) == 0 {
		return
	}

	r.columns[len(r.columns)-1].options.Span = pad - r.length
	r.length = pad
}

type divider struct{}
