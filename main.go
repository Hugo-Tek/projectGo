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

	_, _, _, _ = src.Parser(file, scanner)
}
