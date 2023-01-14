// Package src provides main function on src files
package src

import (
	"fmt"
	"projectGo/src/components"
)

// PrintGrid func to print the grid
func PrintGrid(w *components.Warehouse) {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			switch w.Grid[i][j] {
			case -1:
				fmt.Print("X")
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("P")
			case 2:
				fmt.Print("T")
			case 3:
				fmt.Print("C")
			}
		}
		fmt.Println()
	}
}
