package simpletable

import (
	"errors"
	"fmt"
	"strings"
)

type Table struct {
	header *Row
	rows   []*Row
}

func (t *Table) AddRow(columns ...string) error {
	if len(columns) != t.header.Len() {
		err := fmt.Sprintf("Row [%s] must contain %d columns", strings.Join(columns, ", "), t.header.Len())
		return errors.New(err)
	}

	return nil
}

func (t *Table) String() string {
	return ""
}

func (t *Table) Print() {
	fmt.Println(t.String())
}

func New(headers ...string) *Table {
	header := &Row{}

	for _, h := range headers {
		header.AddColumn()
	}

	return &Table{
		header: header,
		rows:   []*Row{},
	}
}
