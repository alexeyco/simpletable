// Package main demonstrates the built-in table styles.
package main

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type exampleRow struct {
	id       int
	name     string
	phone    string
	email    string
	quantity int
}

var exampleData = []exampleRow{
	{id: 1, name: "Newton G. Goetz", phone: "252-585-5166", email: "NewtonGGoetz@dayrep.com", quantity: 10},
	{id: 2, name: "Rebecca R. Edney", phone: "865-475-4171", email: "RebeccaREdney@armyspy.com", quantity: 12},
	{id: 3, name: "John R. Jackson", phone: "810-325-1417", email: "JohnRJackson@armyspy.com", quantity: 15},
	{id: 4, name: "Ron J. Gomes", phone: "217-450-8568", email: "RonJGomes@rhyta.com", quantity: 25},
	{id: 5, name: "Penny R. Lewis", phone: "870-794-1666", email: "PennyRLewis@rhyta.com", quantity: 5},
	{id: 6, name: "Sofia J. Smith", phone: "770-333-7379", email: "SofiaJSmith@armyspy.com", quantity: 3},
	{id: 7, name: "Karlene D. Owen", phone: "231-242-4157", email: "KarleneDOwen@jourrapide.com", quantity: 12},
	{id: 8, name: "Daniel L. Love", phone: "978-210-4178", email: "DanielLLove@dayrep.com", quantity: 44},
	{id: 9, name: "Julie T. Dial", phone: "719-966-5354", email: "JulieTDial@jourrapide.com", quantity: 8},
	{id: 10, name: "Juan J. Kennedy", phone: "908-910-8893", email: "JuanJKennedy@dayrep.com", quantity: 16},
}

func main() {
	table := simpletable.New()
	styles := map[string]*simpletable.Style{
		"Default style":         simpletable.StyleDefault,
		"Compact style":         simpletable.StyleCompact,
		"Compact Lite style":    simpletable.StyleCompactLite,
		"Compact Classic style": simpletable.StyleCompactClassic,
		"Markdown style":        simpletable.StyleMarkdown,
		"Rounded style":         simpletable.StyleRounded,
		"Unicode style":         simpletable.StyleUnicode,
	}

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

	for n, s := range styles {
		fmt.Println(n)

		table.SetStyle(s)
		table.Println()

		fmt.Println()
	}
}
