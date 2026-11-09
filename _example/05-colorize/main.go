// Package main demonstrates colored table content.
package main

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

const (
	colorDefault = "\x1b[39m"

	colorRed   = "\x1b[91m"
	colorGreen = "\x1b[32m"
	colorBlue  = "\x1b[94m"
	colorGray  = "\x1b[90m"
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
	return fmt.Sprintf("%s%s%s", colorRed, s, colorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, s, colorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, s, colorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", colorGray, s, colorDefault)
}
