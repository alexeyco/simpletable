package main

import (
	"github.com/alexeyco/simpletable"
)

var (
	leftAligned = `Multiline content
left aligned`

	centerAligned = `Multiline content
center aligned`

	rightAligned = `Multiline content
right aligned`

	singleLine = "Single line"
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
				&simpletable.Cell{Align: simpletable.AlignLeft, Span: 2, Text: leftAligned},
				&simpletable.Cell{Text: singleLine},
			},
			{
				&simpletable.Cell{Align: simpletable.AlignCenter, Span: 3, Text: centerAligned},
			},
			{
				&simpletable.Cell{Text: singleLine},
				&simpletable.Cell{Align: simpletable.AlignRight, Span: 2, Text: rightAligned},
			},
		},
	}

	table.Println()
}
