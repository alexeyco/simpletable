package simpletable

import (
	"testing"
)

func TestTextCell_Len(t *testing.T) {
	c := &TextCell{Content: "12345"}

	if c.Len() != 5 {
		t.Error("Wrong *TextCell length")
	}
}

func TestTextCell_IsSpanned(t *testing.T) {
	c := &TextCell{Content: "12345", Span: 2}

	if !c.IsSpanned() {
		t.Error("*TextCell must be spanned")
	}
}

func TestTextCell_IsSpanned2(t *testing.T) {
	c := &TextCell{Content: "12345", Span: 1}

	if c.IsSpanned() {
		t.Error("*TextCell must not be spanned")
	}
}

func TestTextCell_SetWidth(t *testing.T) {
	c := &TextCell{Content: "12345"}
	c.SetWidth(10)

	if c.width != 10 {
		t.Error("Wrong *TextCell width")
	}
}

func TestTextCell_Resize(t *testing.T) {
	c := &TextCell{Content: "12345"}
	c.SetWidth(c.Len())
	c.Resize()

	if c.width != 5 {
		t.Error("Wrong *TextCell resizing")
	}
}

func TestTextCell_Resize2(t *testing.T) {

}

func TestTextCell_Resize3(t *testing.T) {

}

func TestTextCell_String(t *testing.T) {
	c := &TextCell{Content: "12345"}
	c.SetWidth(c.Len())

	if c.String() != "12345" {
		t.Error("Wrong *TextCell contents")
	}
}

func TestTextCell_String2(t *testing.T) {
	c := &TextCell{Content: "12345"}
	c.SetWidth(c.Len() + 5)

	if c.String() != "12345     " {
		t.Error("Wrong *TextCell contents (align: left)")
	}
}

func TestTextCell_String3(t *testing.T) {
	c := &TextCell{Content: "12345", Align: AlignCenter}
	c.SetWidth(c.Len() + 4)

	if c.String() != "  12345  " {
		t.Error("Wrong *TextCell contents (align: center)")
	}
}

func TestTextCell_String4(t *testing.T) {
	c := &TextCell{Content: "12345", Align: AlignRight}
	c.SetWidth(c.Len() + 5)

	if c.String() != "     12345" {
		t.Error("Wrong *TextCell contents (align: right)")
	}
}

func TestDivider_Len(t *testing.T) {
	d := &Divider{}
	d.SetWidth(5)

	if d.Len() != 5 {
		t.Error("Wrong *Divider length")
	}
}

func TestDivider_IsSpanned(t *testing.T) {
	d := &Divider{Span: 1}

	if d.IsSpanned() {
		t.Error("*Divider must not be spanned")
	}
}

func TestDivider_IsSpanned2(t *testing.T) {
	d := &Divider{Span: 2}

	if !d.IsSpanned() {
		t.Error("*Divider must be spanned")
	}
}

func TestDivider_SetWidth(t *testing.T) {
	d := &Divider{}
	d.SetWidth(5)

	if d.width != 5 {
		t.Error("Wrong *Divider width")
	}
}

func TestDivider_String(t *testing.T) {
	tbl := New()
	tbl.Header = &Header{
		Cells: []Cell{
			&TextCell{Content: "AAA"},
			&TextCell{Content: "BBB"},
		},
	}

	tbl.Body = &Body{
		Cells: [][]Cell{
			{
				&TextCell{Span: 2, Content: "CCC"},
			},
		},
	}

	tbl.Footer = &Footer{
		Cells: []Cell{
			&TextCell{Content: "DDD"},
			&TextCell{Content: "EEE"},
		},
	}

	tbl.String()
	s := tbl.dividers[0].String()

	if s != "+-----+-----+" {
		t.Errorf("Wrong *Divider string [%s]", s)
	}
}
