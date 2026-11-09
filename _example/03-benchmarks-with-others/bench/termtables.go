package bench

import "github.com/apcera/termtables"

// Termtables renders the sample rows with termtables.
func Termtables() string {
	table := termtables.CreateTable()
	table.AddHeaders("#", "NAME", "PHONE", "EMAIL", "QTTY")

	subtotal := 0
	for _, row := range exampleData {
		table.AddRow(row.id, row.name, row.phone, row.email, row.quantity)
		subtotal += row.quantity
	}

	table.AddRow("", "", "", "Subtotal", subtotal)
	return table.Render()
}
