package main

import (
	"github.com/alexeyco/simpletable"
)

func main() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "FOO"},
			{Align: simpletable.AlignCenter, Text: "BAR"},
			{Align: simpletable.AlignCenter, Text: "BAZ"},
		},
	}

	table.Body = &simpletable.Body{
		Cells: [][]*simpletable.Cell{
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 2, Text: "11111"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "11111"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "22222"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 2, Text: "22222"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 3, Text: "33333"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "44444"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "44444"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Text: "44444"},
			},
		},
	}

	table.Println()
}
