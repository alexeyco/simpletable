package bench_test

import (
	"testing"

	"github.com/alexeyco/simpletable/_example/03-benchmarks-with-others/bench"
)

func BenchmarkUITable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bench.UITable()
	}
}
