package main

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
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
				&simpletable.Cell{Text: red("I am red")},
				&simpletable.Cell{Text: green("I am green")},
				&simpletable.Cell{Text: blue("I am blue")},
			},
		},
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 3, Text: gray("And all columns is well sized")},
		},
	}

	table.Println()
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
