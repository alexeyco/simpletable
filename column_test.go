package simpletable

import "testing"

func TestColumn_Resize(t *testing.T) {
	if testColumn().Width() != 5 {
		t.Error("Wrong *Column width calculation")
	}
}

func TestColumn_IncrementWidth(t *testing.T) {
	c := testColumn()
	c.IncrementWidth(5)

	if c.Width() != 10 {
		t.Error("Wrong *Column width calculation")
	}
}

func testColumn() *Column {
	c := &Column{
		Cells: []Cell{
			&TextCell{Content: "123"},
			&TextCell{Content: "12345"},
		},
		Table: New(),
	}

	c.Resize()
	return c
}
