package bench

import "testing"

func BenchmarkSimpletable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Simpletable()
	}
}
