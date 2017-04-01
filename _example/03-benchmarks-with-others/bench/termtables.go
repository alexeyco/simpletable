package bench

import (
	"github.com/apcera/termtables"
)

func Termtables() string {
	table := termtables.CreateTable()
	table.AddHeaders("#", "NAME", "PHONE", "EMAIL", "QTTY")

	subtotal := 0
	for _, row := range data {
		table.AddRow(row...)
		subtotal += row[4].(int)
	}

	table.AddRow("", "", "", "Subtotal", subtotal)
	return table.Render()
}
