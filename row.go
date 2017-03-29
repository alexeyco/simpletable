package simpletable

type Row struct {
	columns []*Column
}

func (r *Row) AddColumn(c *Column) {
	r.columns = append(r.columns, c)
}

func (r *Row) Columns() []*Column {
	return r.columns
}

func (r *Row) Len() int {
	return len(r.columns)
}

func newRow() *Row {
	return &Row{
		columns: []*Column{},
	}
}
