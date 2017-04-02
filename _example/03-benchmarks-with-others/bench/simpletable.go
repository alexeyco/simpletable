package bench

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

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
	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
			{Text: row[3].(string)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[4])},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row[4].(int)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 4, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", subtotal)},
		},
	}

	return table.String()
}
