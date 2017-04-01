package main

import (
	"github.com/alexeyco/simpletable"
)

func main() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Content: "FOO"},
			{Align: simpletable.AlignCenter, Content: "BAR"},
			{Align: simpletable.AlignCenter, Content: "BAZ"},
		},
	}

	table.Body = &simpletable.Body{
		Cells: [][]*simpletable.Cell{
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 2, Content: "11111"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Content: "11111"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Content: "22222"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 2, Content: "22222"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 3, Content: "33333"},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Content: "44444"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Content: "44444"},
				&simpletable.Cell{Align: simpletable.AlignCenter, Content: "44444"},
			},
		},
	}
simple
	table.Println()
}
