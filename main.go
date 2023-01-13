package main

import (
	"bufio"
	"os"
	"projectGo/src"
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
	src.PrintGrid(warehouse)
}
