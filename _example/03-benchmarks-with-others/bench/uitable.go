package bench

import (
	"github.com/gosuri/uitable"
)

func UITable() string {
	table := uitable.New()
	table.AddRow("#", "NAME", "PHONE", "EMAIL", "QTTY")

	subtotal := 0
	for _, row := range data {
		table.AddRow(row...)
		subtotal += row[4].(int)
	}

	table.AddRow("", "", "", "Subtotal", subtotal)
	return table.String()
}
