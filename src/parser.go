package src

import (
	"bufio"
	"fmt"
	"os"
	"projectGo/src/components"
	"projectGo/src/errors"
	"strings"
)

func Parser(file *os.File, scanner *bufio.Scanner) (*components.Warehouse, int) {
	// Parse the first row (number of squares in width, number of squares in length, number of turns)
	width, height, turns, ret := components.ParseFirstLine(scanner)

	if ret == 1 {
		return nil, ret
	}

	// Initializes tabs to store parcels, pallet trucks and trucks
	var parcels []components.Parcel
	var pallets []components.Pallet
	var trucks []components.Truck
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Parse the following lines
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 4 {
			parcels = components.CreatePackage(parts, parcels)
		} else if len(parts) == 3 {
			pallets = components.CreatePallet(parts, pallets)
		} else if len(parts) == 5 {
			trucks = components.CreateTruck(parts, trucks)
		} else {
			fmt.Println("Erreur de format: ligne non reconnue")
			return nil, 1
		}
		if errors.ScannerError(scanner) == 1 {
			return nil, 1
		}
	}
	return components.InitWarehouse(width, height, turns, parcels, pallets, trucks, grid), 0
}
