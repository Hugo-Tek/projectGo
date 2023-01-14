// Package src provides main function on src files
package src

import (
	"fmt"
	"projectGo/src/components"
)

func checkEmpty(warehouse *components.Warehouse, palletTruck *components.PalletTruck) (int, int, bool) {
	xEmpty := -1
	yEmpty := -1
	objectNear := false

	x := palletTruck.X
	y := palletTruck.Y

	if x+1 < warehouse.Width && warehouse.Grid[y][x+1] == 0 {
		xEmpty = x + 1
	} else if x+1 < warehouse.Width {
		objectNear = true
	}

	if x-1 >= 0 && warehouse.Grid[y][x-1] == 0 {
		xEmpty = x - 1
	} else if x+1 >= 0 {
		objectNear = true
	}

	if y+1 < warehouse.Height && warehouse.Grid[y+1][x] == 0 {
		yEmpty = y + 1
	} else if y+1 < warehouse.Height {
		objectNear = true
	}

	if y-1 >= 0 && warehouse.Grid[y-1][x] == 0 {
		yEmpty = y - 1
	} else if y-1 >= 0 {
		objectNear = true
	}
	return xEmpty, yEmpty, objectNear
}

// MoveIfNotEmpty func to move if there is some things around it
func MoveIfNotEmpty(warehouse *components.Warehouse, palletTruck *components.PalletTruck) {
	xEmpty, yEmpty, objectNear := checkEmpty(warehouse, palletTruck)

	if objectNear {
		MoveTo(warehouse, palletTruck, xEmpty, yEmpty)
	} else {
		palletTruck.Status = "WAIT"
		fmt.Println(palletTruck.Name, "WAIT")
	}
}

// MoveTo func to move
func MoveTo(warehouse *components.Warehouse, palletTruck *components.PalletTruck, xToGo, yToGo int) {
	palletTruck.Status = "GO"
	x := palletTruck.X
	y := palletTruck.Y

	if x < xToGo && x+1 < warehouse.Width && warehouse.Grid[y][x+1] == 0 {
		palletTruck.X++
	} else if x > xToGo && x-1 >= 0 && warehouse.Grid[y][x-1] == 0 {
		palletTruck.X--
	} else if y < yToGo && y+1 < warehouse.Height && warehouse.Grid[y+1][x] == 0 {
		palletTruck.Y++
	} else if y > yToGo && y-1 >= 0 && warehouse.Grid[y-1][x] == 0 {
		palletTruck.Y--
	} else {
		palletTruck.Status = "WAIT"
		fmt.Println(palletTruck.Name, "WAIT")
	}
	FillGrid(warehouse)
	fmt.Printf("%s GO [%d,%d]\n", palletTruck.Name, palletTruck.X, palletTruck.Y)
}
