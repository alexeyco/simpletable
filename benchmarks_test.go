package simpletable

import (
	"fmt"
	"testing"
)

var (
	benchStyleData = [][]interface{}{
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

	benchStyleDefaultTable        = benchStyleTable(StyleDefault)
	benchStyleCompactTable        = benchStyleTable(StyleCompact)
	benchStyleCompactLiteTable    = benchStyleTable(StyleCompactLite)
	benchStyleCompactClassicTable = benchStyleTable(StyleCompactClassic)
	benchStyleMarkdownTable       = benchStyleTable(StyleMarkdown)
	benchStyleRoundedTable        = benchStyleTable(StyleRounded)
	benchStyleUnicodeTable        = benchStyleTable(StyleUnicode)
)

func BenchmarkStyleDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleDefaultTable.String()
	}
}

func BenchmarkStyleCompact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleCompactTable.String()
	}
}

func BenchmarkStyleCompactLite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleCompactLiteTable.String()
	}
}

func BenchmarkStyleCompactClassic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleCompactClassicTable.String()
	}
}

func BenchmarkStyleMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleMarkdownTable.String()
	}
}

func BenchmarkStyleRounded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleRoundedTable.String()
	}
}

func BenchmarkStyleUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchStyleUnicodeTable.String()
	}
}

func benchStyleTable(s *Style) *Table {
	t := New()
	t.SetStyle(s)

	t.Header = &Header{
		Cells: []*Cell{
			{Align: AlignCenter, Text: "#"},
			{Align: AlignCenter, Text: "NAME"},
			{Align: AlignCenter, Text: "PHONE"},
			{Align: AlignCenter, Text: "EMAIL"},
			{Align: AlignCenter, Text: "QTTY"},
		},
	}

	subtotal := 0
	for _, row := range benchStyleData {
		r := []*Cell{
			{Align: AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
			{Text: row[3].(string)},
			{Align: AlignRight, Text: fmt.Sprintf("%d", row[4])},
		}

		t.Body.Cells = append(t.Body.Cells, r)
		subtotal += row[4].(int)
	}

	t.Footer = &Footer{
		Cells: []*Cell{
			{Align: AlignRight, Span: 4, Text: "Subtotal"},
			{Align: AlignRight, Text: fmt.Sprintf("%d", subtotal)},
		},
	}

	return t
}
