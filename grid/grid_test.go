package grid_test

import (
	"testing"

	"github.com/alexeyco/simpletable/grid"
	"github.com/stretchr/testify/assert"
)

func TestGrid_Cells(t *testing.T) {
	t.Parallel()

	type handler func(*grid.Grid)

	testData := [...]struct {
		name    string
		grid    *grid.Grid
		prepare handler
		cells   []*grid.Cell
	}{
		{
			name: "ResizeColumns",
			grid: grid.New(2, 2),
			prepare: func(g *grid.Grid) {
				g.SetWidth(0, 2)
				g.SetWidth(1, 3)
			},
			cells: []*grid.Cell{
				{Row: 0, Col: 0, HSpan: 1, VSpan: 1, Width: 2, Height: 0},
				{Row: 0, Col: 1, HSpan: 1, VSpan: 1, Width: 3, Height: 0},
				{Row: 1, Col: 0, HSpan: 1, VSpan: 1, Width: 2, Height: 0},
				{Row: 1, Col: 1, HSpan: 1, VSpan: 1, Width: 3, Height: 0},
			},
		},
		{
			name: "ResizeRows",
			grid: grid.New(2, 2),
			prepare: func(g *grid.Grid) {
				g.SetHeight(0, 2)
				g.SetHeight(1, 3)
			},
			cells: []*grid.Cell{
				{Row: 0, Col: 0, HSpan: 1, VSpan: 1, Width: 0, Height: 2},
				{Row: 0, Col: 1, HSpan: 1, VSpan: 1, Width: 0, Height: 2},
				{Row: 1, Col: 0, HSpan: 1, VSpan: 1, Width: 0, Height: 3},
				{Row: 1, Col: 1, HSpan: 1, VSpan: 1, Width: 0, Height: 3},
			},
		},
		{
			name: "ExpandColumnHorizontally",
			grid: grid.New(2, 2),
			prepare: func(g *grid.Grid) {
				g.SetWidth(0, 2)
				g.SetWidth(1, 3)

				g.Expand(0, 0, 2)
			},
			cells: []*grid.Cell{
				{Row: 0, Col: 0, HSpan: 2, VSpan: 1, Width: 5, Height: 0},
				{Row: 1, Col: 0, HSpan: 1, VSpan: 1, Width: 2, Height: 0},
				{Row: 1, Col: 1, HSpan: 1, VSpan: 1, Width: 3, Height: 0},
			},
		},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			if testDatum.prepare != nil {
				testDatum.prepare(testDatum.grid)
			}

			assert.Equal(t, testDatum.cells, testDatum.grid.Cells())
		})
	}
}
