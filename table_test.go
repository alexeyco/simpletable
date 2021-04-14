package simpletable_test

import (
	"fmt"
	"testing"

	st "github.com/alexeyco/simpletable"
)

func TestTable(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		name     string
		phone    string
		email    string
		quantity int
	}{
		{
			name:     "Newton G. Goetz",
			phone:    "252-585-5166",
			email:    "NewtonGGoetz@dayrep.com",
			quantity: 10,
		},
		{
			name:     "Rebecca R. Edney",
			phone:    "865-475-4171",
			email:    "RebeccaREdney@armyspy.com",
			quantity: 12,
		},
		{
			name:     "John R. Jackson",
			phone:    "810-325-1417",
			email:    "JohnRJackson@armyspy.com",
			quantity: 15,
		},
		{
			name:     "Ron J. Gomes",
			phone:    "217-450-8568",
			email:    "RonJGomes@rhyta.com",
			quantity: 25,
		},
		{
			name:     "Penny R. Lewis",
			phone:    "870-794-1666",
			email:    "PennyRLewis@rhyta.com",
			quantity: 5,
		},
		{
			name:     "Sofia J. Smith",
			phone:    "770-333-7379",
			email:    "SofiaJSmith@armyspy.com",
			quantity: 3,
		},
		{
			name:     "Karlene D. Owen",
			phone:    "231-242-4157",
			email:    "KarleneDOwen@jourrapide.com",
			quantity: 12,
		},
		{
			name:     "Daniel L. Love",
			phone:    "978-210-4178",
			email:    "DanielLLove@rhyta.com",
			quantity: 44,
		},
		{
			name:     "Julie T. Dial",
			phone:    "719-966-5354",
			email:    "JulieTDial@jourrapide.com",
			quantity: 8,
		},
		{
			name:     "Juan J. Kennedy",
			phone:    "908-910-8893",
			email:    "JuanJKennedy@dayrep.com",
			quantity: 16,
		},
	}

	t.Run("Basic", func(t *testing.T) {
		t.Parallel()

		table := st.New().
			Row(
				st.Column("#", st.Align(st.Center)),
				st.Column("NAME", st.Align(st.Center)),
				st.Column("PHONE", st.Align(st.Center)),
				st.Column("EMAIL", st.Align(st.Center)),
				st.Column("QTTY", st.Align(st.Center)),
			).
			Divider()

		var qtty int
		for n, testDatum := range testData {
			table.Row(
				st.Column(fmt.Sprintf("%d", n+1), st.Align(st.Right)),
				st.Column(testDatum.name),
				st.Column(testDatum.phone),
				st.Column(testDatum.email),
				st.Column(fmt.Sprintf("%d", testDatum.quantity), st.Align(st.Right)),
			)

			qtty += testDatum.quantity
		}

		table.Divider().
			Row(
				st.Column("Subtotal", st.Span(4), st.Align(st.Right)),
				st.Column(fmt.Sprintf("%d", qtty), st.Align(st.Right)),
			)

		t.Log(fmt.Sprintf("\n%s", table.String()))
	})
}
