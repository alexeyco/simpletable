package simpletable

import "testing"

func TestNew(t *testing.T) {
	tbl := New()
	if tbl == nil {
		t.Error("Table must not be empty")
	}
}

func TestTable_Carve(t *testing.T) {
	type Test struct {
		Value  int
		Parts  int
		Result []int
	}

	s := []*Test{
		{
			Value:  5,
			Parts:  3,
			Result: []int{2, 2, 1},
		},
		{
			Value:  7,
			Parts:  2,
			Result: []int{4, 3},
		},
		{
			Value:  6,
			Parts:  3,
			Result: []int{2, 2, 2},
		},
		{
			Value:  1,
			Parts:  3,
			Result: []int{1, 0, 0},
		},
	}

	tbl := New()
	for _, x := range s {
		if !tableTestSlicesEqual(tbl.carve(x.Value, x.Parts), x.Result) {
			t.Error("Wrong column sizes calculation")
		}
	}
}

func tableTestSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
