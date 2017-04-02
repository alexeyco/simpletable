package simpletable

import (
	"testing"
)

func testCellNewCell(s string) *Cell {
	return &Cell{
		Text:    s,
		content: newContent(s),
	}
}

func TestCell_Len(t *testing.T) {
	c := testCellNewCell("12345")

	if c.len() != 5 {
		t.Error("Wrong *Cell length")
	}
}

func TestCell_IsSpanned(t *testing.T) {
	c := &Cell{Text: "12345", Span: 2}

	if !c.isSpanned() {
		t.Error("*Cell must be spanned")
	}
}

func TestCell_IsSpanned2(t *testing.T) {
	c := &Cell{Text: "12345", Span: 1}

	if c.isSpanned() {
		t.Error("*Cell must not be spanned")
	}
}

func TestCell_SetWidth(t *testing.T) {
	c := testCellNewCell("12345")

	c.setWidth(10)

	if c.width() != 10 {
		t.Error("Wrong *Cell width")
	}
}

func TestCell_SetHeight(t *testing.T) {
	c := testCellNewCell("12345")

	c.setHeight(3)

	if c.height() != 3 {
		t.Error("Wrong *Cell height")
	}
}

func TestCell_Resize(t *testing.T) {
	c := testCellNewCell("12345")

	c.setWidth(c.len())
	c.resize()

	if c.width() != 5 {
		t.Error("Wrong *Cell resizing")
	}
}

func TestCell_Lines(t *testing.T) {
	c := testCellNewCell("12345")
	c.setWidth(c.len())

	if c.lines()[0] != "12345" {
		t.Error("Wrong *Cell contents")
	}
}

func TestCell_String2(t *testing.T) {
	c := testCellNewCell("12345")
	c.setWidth(c.len() + 5)

	if c.lines()[0] != "12345     " {
		t.Error("Wrong *Cell contents (align: left)")
	}
}

func TestCell_String3(t *testing.T) {
	c := testCellNewCell("12345")
	c.Align = AlignCenter

	c.setWidth(c.len() + 4)

	if c.lines()[0] != "  12345  " {
		t.Error("Wrong *Cell contents (align: center)")
	}
}

func TestCell_String4(t *testing.T) {
	c := testCellNewCell("12345")
	c.Align = AlignRight

	c.setWidth(c.len() + 5)

	if c.lines()[0] != "     12345" {
		t.Error("Wrong *Cell contents (align: right)")
	}
}

func TestDivider_Len(t *testing.T) {
	d := &dividerCell{}
	d.setWidth(5)

	if d.len() != 1 {
		t.Error("Wrong *dividerCell length")
	}
}

func TestDivider_IsSpanned(t *testing.T) {
	d := &dividerCell{span: 1}

	if d.isSpanned() {
		t.Error("*dividerCell must not be spanned")
	}
}

func TestDivider_IsSpanned2(t *testing.T) {
	d := &dividerCell{span: 2}

	if !d.isSpanned() {
		t.Error("*dividerCell must be spanned")
	}
}

func TestDivider_SetWidth(t *testing.T) {
	d := &dividerCell{}
	d.setWidth(5)

	if d.width() != 1 {
		t.Error("Wrong *dividerCell width")
	}
}

func TestDivider_String(t *testing.T) {
	tbl := New()
	tbl.Header = &Header{
		Cells: []*Cell{
			{Text: "AAA"},
			{Text: "BBB"},
		},
	}

	tbl.Body = &Body{
		Cells: [][]*Cell{
			{
				&Cell{Span: 2, Text: "CCC"},
			},
		},
	}

	tbl.Footer = &Footer{
		Cells: []*Cell{
			{Text: "DDD"},
			{Text: "EEE"},
		},
	}

	_ = tbl.String()
	s := tbl.dividers[0].lines()[0]

	if s != "+-----+-----+" {
		t.Errorf("Wrong *dividerCell toString [%s]", s)
	}
}

func TestDivider_String2(t *testing.T) {
	tbl := New()
	tbl.Header = &Header{
		Cells: []*Cell{
			{Text: "AAA"},
			{Text: "BBB"},
		},
	}

	tbl.Body = &Body{
		Cells: [][]*Cell{
			{
				&Cell{Span: 2, Text: "CCC CCC  CCC"},
			},
		},
	}

	_ = tbl.String()
	s := tbl.dividers[0].lines()[0]

	if s != "+-------+------+" {
		t.Errorf("Wrong *dividerCell toString [%s]", s)
	}
}

func TestEmptyCell_SetHeight(t *testing.T) {
	d := &emptyCell{content: &content{c: []string{""}}}
	d.setHeight(3)

	if d.height() != 3 {
		t.Error("Wrong *emptyCell width")
	}
}

func TestEmptyCell_Lines(t *testing.T) {
	d := &emptyCell{content: &content{c: []string{""}}}
	d.setHeight(2)

	l := d.lines()
	if d.height() != 2 || (l[0] != "" && l[1] != "") {
		t.Error("Wrong *emptyCell lines")
	}
}
