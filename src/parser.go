// Package src provides main function on src files
package src

import (
	"bufio"
	"fmt"
	"os"
	"projectGo/src/components"
	"strings"
)

// Parser func to parse the parameter file
func Parser(file *os.File, scanner *bufio.Scanner) (w *components.Warehouse, nb int) {
	// Parse the first row (number of squares in width, number of squares in length, number of turns)
	width, height, turns, ret := components.ParseFirstLine(scanner)
	nb = 0

	if ret == 1 {
		fmt.Println("😱")
		return nil, ret
	}

	// Initializes tabs to store parcels, pallet trucks and trucks
	var parcels []components.Parcel
	var PalletTruck []components.PalletTruck
	var trucks []components.Truck
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Parse the following lines
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch len(parts) {
		case 4:
			parcels = components.CreatePackage(parts, parcels)
		case 3:
			PalletTruck = components.CreatePallet(parts, PalletTruck)
		case 5:
			trucks = components.CreateTruck(parts, trucks)
		default:
			fmt.Println("Erreur de format: ligne non reconnue")
			fmt.Println("😱")
			return nil, 1
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			fmt.Println("😱")
			return nil, 1
		}
	}
	return components.InitWarehouse(width, height, turns, parcels, PalletTruck, trucks, grid), nb
}
