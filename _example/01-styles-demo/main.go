package main

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

var (
	data = [][]interface{}{
		{1, "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com", 10},
		{2, "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com", 12},
		{3, "John R. Jackson", "810-325-1417", "JohnRJackson@armyspy.com", 15},
		{4, "Ron J. Gomes", "217-450-8568", "RonJGomes@rhyta.com", 25},
		{5, "Penny R. Lewis", "870-794-1666", "PennyRLewis@rhyta.com", 5},
		{6, "Sofia J. Smith", "770-333-7379", "SofiaJSmith@armyspy.com", 3},
		{7, "Karlene D. Owen", "231-242-4157", "KarleneDOwen@jourrapide.com", 12},
		{8, "Daniel L. Love", "978-210-4178", "DanielLLove@rhyta.com", 44},
		{9, "Julie T. Dial", "719-966-5354", "JulieTDial@jourrapide.com", 8},
		{10, "Juan J. Kennedy", "908-910-8893", "JuanJKennedy@dayrep.com", 16},
	}

	styles = map[string]*simpletable.Style{
		"Default style":         simpletable.StyleDefault,
		"Compact style":         simpletable.StyleCompact,
		"Compact Lite style":    simpletable.StyleCompactLite,
		"Compact Classic style": simpletable.StyleCompactClassic,
		"Markdown style":        simpletable.StyleMarkdown,
		"Rounded style":         simpletable.StyleRounded,
		"Unicode style":         simpletable.StyleUnicode,
	}
)

func main() {
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

	for n, s := range styles {
		fmt.Println(n)

		table.SetStyle(s)
		table.Println()

		fmt.Println()
	}
}
