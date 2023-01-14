// Package components provides main function on components files
package components

import (
	"testing"
)

func TestFindClosestOne(t *testing.T) {
	// Test case 1
	slice := [][]int{
		{0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 3, 0},
		{1, 0, 0, 0, 0},
	}
	x3, y3 := 3, 3
	wantedNumber := 1
	expectedX, expectedY, expectedDist := 3, 0, 3
	x, y, dist := FindClosestOne(slice, x3, y3, wantedNumber)
	if x != expectedX || y != expectedY || dist != expectedDist {
		t.Errorf("FindClosestOne(%v, %d, %d, %d) = (%d, %d, %d), expected (%d, %d, %d)", slice, x3, y3, wantedNumber, x, y, dist, expectedX, expectedY, expectedDist)
	}

	// Test case 2
	slice = [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0},
	}
	x3, y3 = 3, 3
	wantedNumber = 1
	expectedX, expectedY, expectedDist = 2, 1, 3
	x, y, dist = FindClosestOne(slice, x3, y3, wantedNumber)
	if x != expectedX || y != expectedY || dist != expectedDist {
		t.Errorf("FindClosestOne(%v, %d, %d, %d) = (%d, %d, %d), expected (%d, %d, %d)", slice, x3, y3, wantedNumber, x, y, dist, expectedX, expectedY, expectedDist)
	}
}
