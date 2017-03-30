package simpletable

var (
	StyleDefault = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "+",
			Center:       "-",
			Right:        "+",
			Intersection: "+",
		},
		Cell: "|",
	}

	StyleCompact = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "",
			Center:       "=",
			Right:        "",
			Intersection: " ",
		},
		Cell: " ",
	}

	StyleCompactLite = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "",
			Center:       "-",
			Right:        "",
			Intersection: " ",
		},
		Cell: " ",
	}

	StyleMarkdown = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "|",
			Center:       "-",
			Right:        "|",
			Intersection: "|",
		},
		Cell: "|",
	}

	StyleRounded = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "+",
			Center:       "-",
			Right:        "+",
			Intersection: "+",
		},
		Cell: "|",
	}

	StyleUnicode = &Style{
		Border: &StyleBorder{
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
		Divider: &StyleDivider{
			Left:         "╟",
			Center:       "━",
			Right:        "╢",
			Intersection: "┼",
		},
		Cell: "│",
	}
)

type StyleBorder struct {
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

type StyleDivider struct {
	Left         string
	Center       string
	Right        string
	Intersection string
}

type Style struct {
	Border  *StyleBorder
	Divider *StyleDivider
	Cell    string
}
