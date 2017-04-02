package simpletable

import "testing"

func TestRow_String(t *testing.T) {
	c := testCellNewCell("12345")
	c.setWidth(10)

	r := &tblRow{
		Cells: []cellInterface{c},
		Table: New(),
	}

	s := r.toStringSlice()[0]
	if s != " 12345      " {
		t.Error("Wrong row formatting")
	}
}

func TestRow_IsDivider(t *testing.T) {
	d := &tblRow{
		Cells: []cellInterface{&dividerCell{}},
		Table: New(),
	}

	if !d.isDivider() {
		t.Error("Must be divider row")
	}
}

func TestRow_IsDivider2(t *testing.T) {
	n := &tblRow{
		Cells: []cellInterface{&Cell{Text: ""}},
		Table: New(),
	}

	if n.isDivider() {
		t.Error("Must not be divider row")
	}
}
