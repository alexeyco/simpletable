package simpletable_test

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexeyco/simpletable"
	"github.com/stretchr/testify/require"
)

// updateGolden rewrites the testdata/*.golden fixtures with the rendered
// output instead of comparing them, via go test -update-golden.
var updateGolden = flag.Bool("update-golden", false, "rewrite golden fixtures with the rendered output")

// goldenCase is one golden fixture: a fixture name paired with a builder
// returning a fresh table whose rendered output is frozen in
// testdata/<name>.golden.
type goldenCase struct {
	name  string
	table func() *simpletable.Table
}

// goldenCases lists every golden fixture: the benchStyleTable shape in all
// 7 built-in styles, plus multiline, ANSI-colored, CJK-wide, and
// spanned-cell fixtures rendered with StyleUnicode.
var goldenCases = []goldenCase{
	{name: "style-default", table: goldenStyleTable(simpletable.StyleDefault)},
	{name: "style-compact", table: goldenStyleTable(simpletable.StyleCompact)},
	{name: "style-compact-lite", table: goldenStyleTable(simpletable.StyleCompactLite)},
	{name: "style-compact-classic", table: goldenStyleTable(simpletable.StyleCompactClassic)},
	{name: "style-markdown", table: goldenStyleTable(simpletable.StyleMarkdown)},
	{name: "style-rounded", table: goldenStyleTable(simpletable.StyleRounded)},
	{name: "style-unicode", table: goldenStyleTable(simpletable.StyleUnicode)},
	{name: "multiline", table: multilineGoldenTable},
	{name: "ansi", table: ansiGoldenTable},
	{name: "cjk", table: cjkGoldenTable},
	{name: "spanned", table: spannedGoldenTable},
}

func TestGolden(t *testing.T) {
	for _, tc := range goldenCases {
		t.Run(tc.name, func(t *testing.T) {
			requireGolden(t, tc.name, tc.table().String())
		})
	}
}

// goldenStyleTable returns a builder for the benchStyleTable shape with the
// given style.
func goldenStyleTable(style *simpletable.Style) func() *simpletable.Table {
	return func() *simpletable.Table {
		return benchStyleTable(style)
	}
}

// requireGolden asserts that got equals the committed fixture for name byte
// for byte, or rewrites the fixture when -update-golden is set.
func requireGolden(t *testing.T, name, got string) {
	t.Helper()

	path := goldenPath(name)

	if *updateGolden {
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o750))
		require.NoError(t, os.WriteFile(path, []byte(got), 0o600))

		return
	}

	require.Equal(t, goldenWant(t, name), got, "rendered output differs from %s", path)
}

// goldenPath returns the fixture path for a golden case name.
func goldenPath(name string) string {
	return filepath.Join("testdata", name+".golden")
}

// goldenWant returns the committed fixture content for a golden case name.
func goldenWant(t *testing.T, name string) string {
	t.Helper()

	want, err := os.ReadFile(goldenPath(name))
	require.NoError(t, err, "golden fixture is missing; run go test -update-golden to create it")

	return string(want)
}

// multilineGoldenTable freezes newline splitting edge cases: per-line
// trimming of padding whitespace, a trailing newline, a leading newline, a
// bare newline, empty text, an interior empty line, and all three alignment
// padding branches on multiline cells.
func multilineGoldenTable() *simpletable.Table {
	t := simpletable.New()
	t.SetStyle(simpletable.StyleUnicode)

	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "MULTI\nLINE"},
			{Align: simpletable.AlignCenter, Text: "EDGE"},
			{Align: simpletable.AlignCenter, Text: "NOTE"},
		},
	}

	t.Body.Cells = [][]*simpletable.Cell{
		{
			{Text: "  padded line  \n  second   "},
			{Text: "with\ntrailing\n"},
			{Text: "row one"},
		},
		{
			{Text: "\nleading"},
			{Text: ""},
			{Text: "row two"},
		},
		{
			{Text: "\n"},
			{Text: "two\n\nnewlines"},
			{Text: "row three"},
		},
		{
			{Text: ""},
			{Text: "center\n\nblank"},
			{Text: "row four"},
		},
		{
			{Align: simpletable.AlignLeft, Text: "lft\nxx"},
			{Align: simpletable.AlignCenter, Text: "ctr\nyy"},
			{Align: simpletable.AlignRight, Text: "rgt\nzz"},
		},
	}

	t.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 2, Text: "spans\ntwo"},
			{Align: simpletable.AlignRight, Text: "5"},
		},
	}

	return t
}

// ansiGoldenTable freezes ANSI escape handling: escape sequences stay in
// the rendered output while cell widths are measured after stripping them,
// including one 8-bit CSI sequence, padding spaces inside and outside the
// escape pairs, a colored multiline cell, and a colored line whose padding
// is zero.
func ansiGoldenTable() *simpletable.Table {
	t := simpletable.New()
	t.SetStyle(simpletable.StyleUnicode)

	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "COLORED"},
			{Align: simpletable.AlignCenter, Text: "ESCAPES"},
		},
	}

	t.Body.Cells = [][]*simpletable.Cell{
		{
			{Text: "\x1b[91mred text\x1b[0m"},
			{Text: "\x1b[32mgreen\x1b[0m padded"},
		},
		{
			{Text: "\x9b94mbright blue\x9b0m"},
			{Text: "  plain  "},
		},
		{
			{Text: "\x1b[1;4mbold underline\x1b[0m"},
			{Text: "\x1b[38;5;196m256-color\x1b[0m"},
		},
		{
			{Align: simpletable.AlignRight, Text: "\x1b[31m  right-aligned  \x1b[0m"},
			{Align: simpletable.AlignCenter, Text: "  \x1b[35mcentered\x1b[0m  "},
		},
	}

	t.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "\x1b[36mfirst\x1b[0m\n\x1b[36msecond\x1b[0m"},
			{Align: simpletable.AlignRight, Text: "\x1b[33m4\x1b[0m"},
		},
	}

	return t
}

// cjkGoldenTable freezes runewidth-based measurement: CJK wide runes, an
// emoji, mixed ASCII and wide runes, and a multiline CJK cell.
func cjkGoldenTable() *simpletable.Table {
	t := simpletable.New()
	t.SetStyle(simpletable.StyleUnicode)

	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "言語"},
			{Align: simpletable.AlignCenter, Text: "NOTES"},
		},
	}

	t.Body.Cells = [][]*simpletable.Cell{
		{
			{Text: "こんにちは"},
			{Text: "five wide runes"},
		},
		{
			{Text: "中文测试"},
			{Text: "four wide runes"},
		},
		{
			{Align: simpletable.AlignCenter, Text: "混合mixed"},
			{Text: "wide plus ascii"},
		},
		{
			{Text: "絵文字🚀です"},
			{Text: "emoji is wide"},
		},
		{
			{Align: simpletable.AlignRight, Text: "日本語\nテスト"},
			{Text: "multiline cjk"},
		},
	}

	t.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 2, Text: "合計"},
		},
	}

	return t
}

// spannedGoldenTable freezes span rendering: every row fills the
// five-column layout exactly through a different Span combination, and the
// long three-column span forces the spanned columns to grow.
func spannedGoldenTable() *simpletable.Table {
	t := simpletable.New()
	t.SetStyle(simpletable.StyleUnicode)

	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ONE"},
			{Align: simpletable.AlignCenter, Text: "TWO"},
			{Align: simpletable.AlignCenter, Text: "THREE"},
			{Align: simpletable.AlignCenter, Text: "FOUR"},
			{Align: simpletable.AlignCenter, Text: "FIVE"},
		},
	}

	t.Body.Cells = [][]*simpletable.Cell{
		{
			{Span: 5, Text: "full row span"},
		},
		{
			{Span: 3, Text: "wide spanned text grows columns"},
			{Span: 2, Text: "span of two"},
		},
		{
			{Text: "one"},
			{Span: 4, Text: "span of four after a single cell"},
		},
		{
			{Span: 2, Text: "left span"},
			{Text: "mid"},
			{Span: 2, Text: "right span"},
		},
		{
			{Text: "a"},
			{Text: "b"},
			{Text: "c"},
			{Text: "d"},
			{Text: "e"},
		},
	}

	t.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 4, Text: "subtotal"},
			{Align: simpletable.AlignRight, Text: "5"},
		},
	}

	return t
}
