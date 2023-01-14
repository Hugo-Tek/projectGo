package src

import (
	"projectGo/src/components"
	"testing"
)

func TestCheckEmpty(t *testing.T) {
	// Test case 1
	warehouse := &components.Warehouse{
		Width:  5,
		Height: 5,
		Grid: [][]int{
			{1, 1, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 2, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 1, 1, 1, 1},
		},
	}
	palletTruck := &components.PalletTruck{X: 2, Y: 2}
	expectedX, expectedY, expectedObjectNear := 1, 1, false
	x, y, objectNear := checkEmpty(warehouse, palletTruck)
	if x != expectedX || y != expectedY || objectNear != expectedObjectNear {
		t.Errorf("checkEmpty(%v, %v) = (%d, %d, %v), expected (%d, %d, %v)", warehouse, palletTruck, x, y, objectNear, expectedX, expectedY, expectedObjectNear)
	}
}
