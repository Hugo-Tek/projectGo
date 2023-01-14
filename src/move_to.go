// Package src provides main function on src files
package src

import (
	"fmt"
	"projectGo/src/components"
)

func checkPlace(a, sizeMax, element int) (res bool) {
	if a+1 < sizeMax && element == 0 {
		return true
	}
	return false
}

func checkEmptyX(warehouse *components.Warehouse, palletTruck *components.PalletTruck) (xEmpty int, yEmpty int, objectNear bool) {
	xEmpty = -1
	yEmpty = -1
	objectNear = false

	x := palletTruck.X
	y := palletTruck.Y

	if checkPlace(x, warehouse.Width, warehouse.Grid[y][x+1]) {
		xEmpty = x + 1
	} else if x+1 < warehouse.Width {
		objectNear = true
	}

	if x-1 >= 0 && warehouse.Grid[y][x-1] == 0 {
		xEmpty = x - 1
	} else if x+1 >= 0 {
		objectNear = true
	}
	return xEmpty, yEmpty, objectNear
}

func checkEmpty(warehouse *components.Warehouse, palletTruck *components.PalletTruck) (xEmpty int, yEmpty int, objectNear bool) {
	xEmpty = -1
	yEmpty = -1
	objectNear = false

	x := palletTruck.X
	y := palletTruck.Y

	xEmpty, yEmpty, objectNear = checkEmptyX(warehouse, palletTruck)
	if checkPlace(y, warehouse.Height, warehouse.Grid[y+1][x]) {
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
	FillGrid(warehouse)
}

func checkIfValueIsLower(a, aToGo, sizeMax int) (isGood bool) {
	if a < aToGo && a+1 < sizeMax {
		return true
	}
	return false
}

func checkIfValueIsUpper(a, aToGo int) (isGood bool) {
	if a > aToGo && a > 0 {
		return true
	}
	return false
}

// MoveTo func to move
func MoveTo(warehouse *components.Warehouse, palletTruck *components.PalletTruck, xToGo, yToGo int) {
	palletTruck.Status = "GO"
	x := palletTruck.X
	y := palletTruck.Y

	switch {
	case checkIfValueIsLower(x, xToGo, warehouse.Width) && warehouse.Grid[y][x+1] == 0:
		palletTruck.X++
	case checkIfValueIsUpper(x, xToGo) && warehouse.Grid[y][x-1] == 0:
		palletTruck.X--
	case checkIfValueIsLower(y, yToGo, warehouse.Height) && warehouse.Grid[y+1][x] == 0:
		palletTruck.Y++
	case checkIfValueIsUpper(y, yToGo) && warehouse.Grid[y-1][x] == 0:
		palletTruck.Y--
	default:
		MoveIfNotEmpty(warehouse, palletTruck)
		return
	}
	FillGrid(warehouse)
	fmt.Printf("%s GO [%d,%d]\n", palletTruck.Name, palletTruck.X, palletTruck.Y)
}
