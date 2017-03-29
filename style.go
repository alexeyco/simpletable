package simpletable

type OuterBorderIntersection struct {
	Top    string
	Bottom string
}

type OuterBorder struct {
	TopLeft      string
	Top          string
	TopRight     string
	Right        string
	BottomRight  string
	Bottom       string
	BottomLeft   string
	Left         string
	Intersection *OuterBorderIntersection
}

type HeaderBorder struct {
	Line         string
	Left         string
	Right        string
	Intersection string
}

type Style struct {
	Outer  *OuterBorder
	Inner  string
	Header *HeaderBorder
}

var (
	//
	// +===========+==========+
	// | FIRSTNAME | LASTNAME |
	// +===========+==========+
	// | John      | Doe      |
	// | Ivan      | Ivanov   |
	// +===========+==========+
	//
	StyleSimpleBordered = &Style{
		Outer: &OuterBorder{
			TopLeft:     "+",
			Top:         "=",
			TopRight:    "+",
			Right:       "|",
			BottomRight: "+",
			Bottom:      "=",
			BottomLeft:  "+",
			Left:        "|",
			Intersection: &OuterBorderIntersection{
				Top:    "+",
				Bottom: "+",
			},
		},
		Inner: "|",
		Header: &HeaderBorder{
			Line:         "=",
			Left:         "+",
			Right:        "+",
			Intersection: "+",
		},
	}

	//
	//  FIRSTNAME | LASTNAME
	// ======================
	//  John      | Doe
	//  Ivan      | Ivanov
	//
	StyleSimpleMinimalistic = &Style{
		Inner: "|",
		Header: &HeaderBorder{
			Line:         "=",
			Intersection: "=",
		},
	}

	//
	// ╔═══════════╤══════════╗
	// ║ FIRSTNAME │ LASTNAME ║
	// ╠═══════════╪══════════╣
	// ║ John      │ Doe      ║
	// ║ Ivan      │ Ivanov   ║
	// ╚═══════════╧══════════╝
	//
	StylePrettyBordered = &Style{
		Outer: &OuterBorder{
			TopLeft:     "╔",
			Top:         "═",
			TopRight:    "╗",
			Right:       "║",
			BottomRight: "╝",
			Bottom:      "═",
			BottomLeft:  "╚",
			Left:        "║",
			Intersection: &OuterBorderIntersection{
				Top:    "╤",
				Bottom: "╧",
			},
		},
		Inner: "│",
		Header: &HeaderBorder{
			Line:         "═",
			Left:         "╠",
			Right:        "╣",
			Intersection: "╪",
		},
	}

	//
	//  FIRSTNAME │ LASTNAME
	// ═══════════╪══════════
	//  John      │ Doe
	//  Ivan      │ Ivanov
	//
	StylePrettyMinimalistic = &Style{
		Inner: "│",
		Header: &HeaderBorder{
			Line:         "═",
			Intersection: "╪",
		},
	}
)
