package simpletable

import "testing"

func TestColumn_Resize(t *testing.T) {
	if testColumn().getWidth() != 5 {
		t.Error("Wrong *tblColumn width calculation")
	}
}

func TestColumn_IncrementWidth(t *testing.T) {
	c := testColumn()
	c.incrementWidth(5)

	if c.getWidth() != 10 {
		t.Error("Wrong *tblColumn width calculation")
	}
}

func testColumn() *tblColumn {
	c := &tblColumn{
		Cells: []cellInterface{
			&Cell{Content: "123"},
			&Cell{Content: "12345"},
		},
		Table: New(),
	}

	c.resize()
	return c
}
