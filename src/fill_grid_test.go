package src

import (
	"projectGo/src/components"
	"testing"
)

func TestFillGrid(t *testing.T) {
	// Test case 1
	w := &components.Warehouse{
		Width:  5,
		Height: 5,
		Grid: [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		Parcels: []components.Parcel{
			{X: 1, Y: 1},
			{X: 2, Y: 3},
		},
		PalletTruck: []components.PalletTruck{
			{X: 4, Y: 4},
		},
		Trucks: []components.Truck{
			{X: 3, Y: 3},
		},
	}
	expectedGrid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 3, 0},
		{0, 0, 0, 0, 2},
	}
	FillGrid(w)
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Grid[i][j] != expectedGrid[i][j] {
				t.Errorf("FillGrid() = %v, expected %v", w.Grid, expectedGrid)
				return
			}
		}
	}
}
