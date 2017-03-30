package simpletable

import "testing"

func TestRow_String(t *testing.T) {
	c := &TextCell{Content: "12345"}
	c.SetWidth(10)

	r := &Row{
		Cells: []Cell{c},
		Table: New(),
	}

	s := r.String()
	if s != " 12345      " {
		t.Error("Wrong row formatting")
	}
}

func TestRow_IsDivider(t *testing.T) {
	d := &Row{
		Cells: []Cell{&Divider{}},
		Table: New(),
	}

	if !d.IsDivider() {
		t.Error("Must be divider row")
	}
}

func TestRow_IsDivider2(t *testing.T) {
	n := &Row{
		Cells: []Cell{&TextCell{Content: ""}},
		Table: New(),
	}

	if n.IsDivider() {
		t.Error("Must not be divider row")
	}
}
