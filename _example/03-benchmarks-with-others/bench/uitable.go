package bench

import "github.com/gosuri/uitable"

// UITable renders the sample rows with uitable.
func UITable() string {
	table := uitable.New()
	table.AddRow("#", "NAME", "PHONE", "EMAIL", "QTTY")

	subtotal := 0
	for _, row := range exampleData {
		table.AddRow(row.id, row.name, row.phone, row.email, row.quantity)
		subtotal += row.quantity
	}

	table.AddRow("", "", "", "Subtotal", subtotal)
	return table.String()
}
