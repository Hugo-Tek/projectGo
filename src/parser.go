package src

import (
	"bufio"
	"fmt"
	"os"
	"projectGo/src/components"
	"projectGo/src/errors"
	"strings"
)

func Parser(file *os.File, scanner *bufio.Scanner) (int, int, int, int) {
	// Parse the first row (number of squares in width, number of squares in length, number of turns)
	width, height, turns, ret := components.ParseFirstLine(scanner)

	if ret == 1 {
		return width, height, turns, ret
	}

	// Initializes tabs to store parcels, pallet trucks and trucks
	var parcels []components.Parcel
	var pallets []components.Pallet
	var trucks []components.Truck

	// Parse the following lines
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		fmt.Printf("Longueur: %d\n", len(parts))
		if len(parts) == 4 {
			parcels = components.CreatePackage(parts, parcels)
		} else if len(parts) == 3 {
			pallets = components.CreatePallet(parts, pallets)
		} else if len(parts) == 5 {
			trucks = components.CreateTruck(parts, trucks)
		} else {
			fmt.Println("Erreur de format: ligne non reconnue")
			return width, height, turns, 1
		}
		if errors.ScannerError(scanner) == 1 {
			return width, height, turns, 1
		}
	}
	// Displays the contents of the file
	// fmt.Printf("Largeur de l'entrepôt: %d\n", width)
	// fmt.Printf("Longueur de l'entrepôt: %d\n", height)
	// fmt.Printf("Nombre de tours: %d\n", turns)
	// fmt.Println("Parcels:")
	// for _, p := range parcels {
	// 	fmt.Printf("- %s (%d,%d) %s\n", p.Name, p.X, p.Y, p.Color)
	// }
	// fmt.Println("Pallets:")
	// for _, palette := range pallets {
	// 	fmt.Printf("- %s (%d,%d)\n", palette.Name, palette.X, palette.Y)
	// }
	// fmt.Println("Trucks:")
	// for _, truck := range trucks {
	// 	fmt.Printf("- %s (%d,%d) Charge max: %d Tours avant disponibilité: %d\n", truck.Name, truck.X, truck.Y, truck.MaxLoad, truck.Availability)
	// }
	return width, height, turns, 0
}
