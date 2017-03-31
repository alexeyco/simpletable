package main

import (
	"github.com/alexeyco/simpletable"
)

func main() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []simpletable.Cell{
			&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "FOO"},
			&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "BAR"},
			&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "BAZ"},
		},
	}

	table.Body = &simpletable.Body{
		Cells: [][]simpletable.Cell{
			{
				&simpletable.TextCell{Align: simpletable.AlignCenter, Span: 2, Content: "11111"},
				&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "11111"},
			},
			{
				&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "22222"},
				&simpletable.TextCell{Align: simpletable.AlignCenter, Span: 2, Content: "22222"},
			},
			{
				&simpletable.TextCell{Align: simpletable.AlignCenter, Span: 3, Content: "33333"},
			},
			{
				&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "44444"},
				&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "44444"},
				&simpletable.TextCell{Align: simpletable.AlignCenter, Content: "44444"},
			},
		},
	}

	table.Println()
}
