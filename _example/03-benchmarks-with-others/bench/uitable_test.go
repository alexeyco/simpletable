package bench

import "testing"

func BenchmarkUITable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UITable()
	}
}
