// Package main provides the main function with a loop
package main

import (
	"bufio"
	"fmt"
	"os"
	"projectGo/src"
	"projectGo/src/components"
	"projectGo/src/errors"
)

func main() {
	// Open the file to read
	file, err := os.Open("file.txt")
	ret := errors.Check(err)
	if ret == 1 {
		return
	}
	defer file.Close()

	// Create a scanner to scan the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	warehouse, ret := src.Parser(file, scanner)
	if ret == 1 {
		return
	}
	warehouse = src.FillGrid(warehouse)
	for turn := 1; turn <= warehouse.Turns; turn++ {
		fmt.Printf("\nTurn %d\n", turn)
		src.PrintGrid(warehouse)

		src.MovePalletTrucks(warehouse)
		src.ManageTrucks(warehouse)
		if components.CountingParcels(warehouse) == 0 {
			fmt.Println("😎")
			return
		}
	}
	fmt.Println("🙂")
}
