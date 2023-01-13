package src

import (
	"fmt"
	"projectGo/src/components"
)

func PrintGrid(w *components.Warehouse) {
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			switch w.Grid[i][j] {
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
