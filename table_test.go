package simpletable

import "testing"

func TestNew(t *testing.T) {
	tbl := New()
	if tbl == nil {
		t.Error("Table must not be empty")
	}
}
