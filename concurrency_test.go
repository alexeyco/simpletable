package simpletable_test

import (
	"sync"
	"testing"
)

// TestReRenderStability renders each golden fixture's table twice in a row
// and asserts both outputs match the fixture. The second render recycles
// the content objects released by the first one, including lines appended
// to them by setHeight during the first render, so it proves pooled
// content is fully reset between String() calls.
func TestReRenderStability(t *testing.T) {
	for _, tc := range goldenCases {
		t.Run(tc.name, func(t *testing.T) {
			tbl := tc.table()

			requireGolden(t, tc.name, tbl.String())
			requireGolden(t, tc.name, tbl.String())
		})
	}
}

// TestConcurrentDistinctTables renders distinct tables from separate
// goroutines, each repeatedly, and compares every render against its
// golden fixture. The goroutines share only the content pool, so this is
// the race-detector proof for recycling content objects across concurrent
// renders of different tables. Concurrent String() calls on one shared
// table are not a supported guarantee and are intentionally not tested.
func TestConcurrentDistinctTables(t *testing.T) {
	const goroutines = 8
	const renders = 25

	var wg sync.WaitGroup

	for _, tc := range goldenCases {
		want := goldenWant(t, tc.name)

		for g := 0; g < goroutines; g++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				tbl := tc.table()

				for i := 0; i < renders; i++ {
					if got := tbl.String(); got != want {
						t.Errorf("%s: render %d differs from golden fixture", tc.name, i)

						return
					}
				}
			}()
		}
	}

	wg.Wait()
}
