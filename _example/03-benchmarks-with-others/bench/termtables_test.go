package bench

import "testing"

func BenchmarkTermtables(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Termtables()
	}
}
