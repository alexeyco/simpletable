package main

import (
	"log"

	"github.com/alexeyco/simpletable"
)

var data = [][]string{
	{"1", "Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com"},
	{"2", "Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com"},
	{"3", "John R. Jackson", "810-325-1417", "JohnRJackson@armyspy.com"},
	{"4", "Ron J. Gomes", "217-450-8568", "RonJGomes@rhyta.com"},
	{"5", "Penny R. Lewis", "870-794-1666", "PennyRLewis@rhyta.com"},
	{"6", "Sofia J. Smith", "770-333-7379", "SofiaJSmith@armyspy.com"},
	{"7", "Karlene D. Owen", "231-242-4157", "KarleneDOwen@jourrapide.com"},
	{"8", "Daniel L. Love", "978-210-4178", "DanielLLove@rhyta.com"},
	{"9", "Julie T. Dial", "719-966-5354", "JulieTDial@jourrapide.com"},
	{"10", "Juan J. Kennedy", "908-910-8893", "JuanJKennedy@dayrep.com"},
}

func main() {
	table := simpletable.New("#", "Name", "Phone", "Email")
	table.SetStyle(simpletable.StylePrettyBordered)
	table.SetAlign(simpletable.AlignRight, simpletable.AlignLeft, simpletable.AlignLeft, simpletable.AlignLeft)

	for _, row := range data {
		if err := table.AddRow(row...); err != nil {
			log.Fatalln(err)
		}
	}

	table.Print()
}
