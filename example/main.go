package main

import (
	"github.com/alexeyco/simpletable"
	"log"
)

var data = [][]string{
	{"Newton G. Goetz", "252-585-5166", "NewtonGGoetz@dayrep.com"},
	{"Rebecca R. Edney", "865-475-4171", "RebeccaREdney@armyspy.com"},
	{"John R. Jackson", "810-325-1417", "JohnRJackson@armyspy.com"},
	{"Ron J. Gomes", "217-450-8568", "RonJGomes@rhyta.com"},
	{"Penny R. Lewis", "870-794-1666", "PennyRLewis@rhyta.com"},
	{"Sofia J. Smith", "770-333-7379", "SofiaJSmith@armyspy.com"},
	{"Karlene D. Owen", "231-242-4157", "KarleneDOwen@jourrapide.com"},
	{"Daniel L. Love", "978-210-4178", "DanielLLove@rhyta.com"},
	{"Julie T. Dial", "719-966-5354", "JulieTDial@jourrapide.com"},
	{"Juan J. Kennedy", "908-910-8893", "JuanJKennedy@dayrep.com"},
}

func main() {
	table := simpletable.New("Name", "Phone", "Email")

	for _, row := range data {
		if err := table.AddRow(row...); err != nil {
			log.Fatalln(err)
		}
	}

	table.Print()
}
