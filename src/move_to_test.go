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

func TestMoveTo(t *testing.T) {
	// Test case 1
	w := &components.Warehouse{
		Width:  5,
		Height: 5,
		Grid: [][]int{
			{-1, -1, -1, -1, -1},
			{-1, 1, 0, 0, 0},
			{-1, 0, 0, 1, 0},
			{-1, 0, 0, 3, 0},
			{-1, 0, 0, 0, 2},
		},
	}
	pt := &components.PalletTruck{
		X:      2,
		Y:      2,
		Name:   "pt1",
		Status: "WAIT",
	}
	MoveTo(w, pt, 4, 4)
	expectedX, expectedY, expectedStatus := 3, 3, "GO"
	if pt.X != expectedX || pt.Y != expectedY || pt.Status != expectedStatus {
		t.Errorf("MoveTo() = (%d, %d, %s), expected (%d, %d, %s)", pt.X, pt.Y, pt.Status, expectedX, expectedY, expectedStatus)
	}
}
