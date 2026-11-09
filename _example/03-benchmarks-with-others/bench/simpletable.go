package bench

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

// Simpletable renders the sample rows with simpletable.
func Simpletable() string {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "PHONE"},
			{Align: simpletable.AlignCenter, Text: "EMAIL"},
			{Align: simpletable.AlignCenter, Text: "QTTY"},
		},
	}

	subtotal := 0
	for _, row := range exampleData {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.id)},
			{Text: row.name},
			{Text: row.phone},
			{Text: row.email},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row.quantity)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row.quantity
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 4, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", subtotal)},
		},
	}

	return table.String()
}
