package simpletable

var (
	// StyleDefault - MySql-like table style:
	//
	// +---+------------------+------------+
	// | # |       NAME       |    TAX     |
	// +---+------------------+------------+
	// | 1 | Newton G. Goetz  |   $ 532.70 |
	// | 2 | Rebecca R. Edney |  $ 1423.25 |
	// | 3 | John R. Jackson  |  $ 7526.12 |
	// | 4 | Ron J. Gomes     |   $ 123.84 |
	// | 5 | Penny R. Lewis   |  $ 3221.11 |
	// +---+------------------+------------+
	// |   |         Subtotal | $ 12827.02 |
	// +---+------------------+------------+
	StyleDefault = Style{
		Border: Border{
			TopLeft:            '+',
			Top:                '-',
			TopRight:           '+',
			Right:              '|',
			BottomRight:        '+',
			Bottom:             '-',
			BottomLeft:         '+',
			Left:               '|',
			TopIntersection:    '+',
			BottomIntersection: '+',
		},
		Divider: Divider{
			Left:         '+',
			Center:       '-',
			Right:        '+',
			Intersection: '+',
		},
		Grid: '|',
	}
)

type Border struct {
	TopLeft            rune
	Top                rune
	TopRight           rune
	Right              rune
	BottomRight        rune
	Bottom             rune
	BottomLeft         rune
	Left               rune
	TopIntersection    rune
	BottomIntersection rune
}

type Divider struct {
	Left         rune
	Center       rune
	Right        rune
	Intersection rune
}

// Style is a table style (borders, dividers etc)
type Style struct {
	Border  Border
	Divider Divider
	Grid    rune
}
