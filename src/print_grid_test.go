package src

import (
	"bytes"
	"fmt"
	"projectGo/src/components"
	"testing"
)

func TestPrintGrid(t *testing.T) {
	// Test case 1
	w := &components.Warehouse{
		Width:  5,
		Height: 5,
		Grid: [][]int{
			{0, 0, 0, -1, -1},
			{-1, 1, 0, 0, 0},
			{-1, 0, 0, 1, 0},
			{-1, 0, 0, 3, 0},
			{-1, 0, 0, 0, 2},
		},
	}

	var buf bytes.Buffer
	fmt.Println("original Output :")
	PrintGrid(w)
	fmt.Println("Test Output :")
	fmt.Print(&buf)
	PrintGrid(w)
}
