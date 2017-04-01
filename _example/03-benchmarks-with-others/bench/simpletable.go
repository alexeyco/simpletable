package bench

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

func Simpletable() string {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Content: "#"},
			{Align: simpletable.AlignCenter, Content: "NAME"},
			{Align: simpletable.AlignCenter, Content: "PHONE"},
			{Align: simpletable.AlignCenter, Content: "EMAIL"},
			{Align: simpletable.AlignCenter, Content: "QTTY"},
		},
	}

	subtotal := 0
	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Content: fmt.Sprintf("%d", row[0].(int))},
			{Content: row[1].(string)},
			{Content: row[2].(string)},
			{Content: row[3].(string)},
			{Align: simpletable.AlignRight, Content: fmt.Sprintf("%d", row[4])},
		}

		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row[4].(int)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 4, Content: "Subtotal"},
			{Align: simpletable.AlignRight, Content: fmt.Sprintf("%d", subtotal)},
		},
	}

	return table.String()
}
