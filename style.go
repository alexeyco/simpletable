package simpletable

var (
	// Default (MySql-like) table style:
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
	StyleDefault = &Style{
		Border: &BorderStyle{
			TopLeft:            "+",
			Top:                "-",
			TopRight:           "+",
			Right:              "|",
			BottomRight:        "+",
			Bottom:             "-",
			BottomLeft:         "+",
			Left:               "|",
			TopIntersection:    "+",
			BottomIntersection: "+",
		},
		Divider: &DividerStyle{
			Left:         "+",
			Center:       "-",
			Right:        "+",
			Intersection: "+",
		},
		Cell: "|",
	}

	// Compact table style:
	//
	//  #         NAME            TAX
	// === ================== ============
	//  1   Newton G. Goetz      $ 532.70
	//  2   Rebecca R. Edney    $ 1423.25
	//  3   John R. Jackson     $ 7526.12
	//  4   Ron J. Gomes         $ 123.84
	//  5   Penny R. Lewis      $ 3221.11
	// === ================== ============
	//              Subtotal   $ 12827.02
	StyleCompact = &Style{
		Border: &BorderStyle{
			TopLeft:            "",
			Top:                "",
			TopRight:           "",
			Right:              "",
			BottomRight:        "",
			Bottom:             "",
			BottomLeft:         "",
			Left:               "",
			TopIntersection:    "",
			BottomIntersection: "",
		},
		Divider: &DividerStyle{
			Left:         "",
			Center:       "=",
			Right:        "",
			Intersection: " ",
		},
		Cell: " ",
	}

	// Compact lite table style:
	//
	//  #         NAME            TAX
	// --- ------------------ ------------
	//  1   Newton G. Goetz      $ 532.70
	//  2   Rebecca R. Edney    $ 1423.25
	//  3   John R. Jackson     $ 7526.12
	//  4   Ron J. Gomes         $ 123.84
	//  5   Penny R. Lewis      $ 3221.11
	// --- ------------------ ------------
	//              Subtotal   $ 12827.02
	StyleCompactLite = &Style{
		Border: &BorderStyle{
			TopLeft:            "",
			Top:                "",
			TopRight:           "",
			Right:              "",
			BottomRight:        "",
			Bottom:             "",
			BottomLeft:         "",
			Left:               "",
			TopIntersection:    "",
			BottomIntersection: "",
		},
		Divider: &DividerStyle{
			Left:         "",
			Center:       "-",
			Right:        "",
			Intersection: " ",
		},
		Cell: " ",
	}

	// Markdown table style:
	//
	// | # |       NAME       |    TAX     |
	// |---|------------------|------------|
	// | 1 | Newton G. Goetz  |   $ 532.70 |
	// | 2 | Rebecca R. Edney |  $ 1423.25 |
	// | 3 | John R. Jackson  |  $ 7526.12 |
	// | 4 | Ron J. Gomes     |   $ 123.84 |
	// | 5 | Penny R. Lewis   |  $ 3221.11 |
	// |---|------------------|------------|
	// |   |         Subtotal | $ 12827.02 |
	StyleMarkdown = &Style{
		Border: &BorderStyle{
			TopLeft:            "",
			Top:                "",
			TopRight:           "",
			Right:              "|",
			BottomRight:        "",
			Bottom:             "",
			BottomLeft:         "",
			Left:               "|",
			TopIntersection:    "",
			BottomIntersection: "",
		},
		Divider: &DividerStyle{
			Left:         "|",
			Center:       "-",
			Right:        "|",
			Intersection: "|",
		},
		Cell: "|",
	}

	// Rounded table style:
	//
	// .---.------------------.------------.
	// | # |       NAME       |    TAX     |
	// +---+------------------+------------+
	// | 1 | Newton G. Goetz  |   $ 532.70 |
	// | 2 | Rebecca R. Edney |  $ 1423.25 |
	// | 3 | John R. Jackson  |  $ 7526.12 |
	// | 4 | Ron J. Gomes     |   $ 123.84 |
	// | 5 | Penny R. Lewis   |  $ 3221.11 |
	// +---+------------------+------------+
	// |   |         Subtotal | $ 12827.02 |
	// '---'------------------'------------'
	StyleRounded = &Style{
		Border: &BorderStyle{
			TopLeft:            ".",
			Top:                "-",
			TopRight:           ".",
			Right:              "|",
			BottomRight:        "'",
			Bottom:             "-",
			BottomLeft:         "'",
			Left:               "|",
			TopIntersection:    ".",
			BottomIntersection: "'",
		},
		Divider: &DividerStyle{
			Left:         "+",
			Center:       "-",
			Right:        "+",
			Intersection: "+",
		},
		Cell: "|",
	}

	// Unicode (awesome!!!) table style:
	StyleUnicode = &Style{
		Border: &BorderStyle{
			TopLeft:            "╔",
			Top:                "═",
			TopRight:           "╗",
			Right:              "║",
			BottomRight:        "╝",
			Bottom:             "═",
			BottomLeft:         "╚",
			Left:               "║",
			TopIntersection:    "╤",
			BottomIntersection: "╧",
		},
		Divider: &DividerStyle{
			Left:         "╟",
			Center:       "━",
			Right:        "╢",
			Intersection: "┼",
		},
		Cell: "│",
	}
)

// BorderStyle defines table border style
type BorderStyle struct {
	TopLeft            string
	Top                string
	TopRight           string
	Right              string
	BottomRight        string
	Bottom             string
	BottomLeft         string
	Left               string
	TopIntersection    string
	BottomIntersection string
}

// DividerStyle defines table divider style
type DividerStyle struct {
	Left         string
	Center       string
	Right        string
	Intersection string
}

// Style is a table style (borders, dividers etc)
type Style struct {
	Border  *BorderStyle
	Divider *DividerStyle
	Cell    string // Symbol between cells
}
