// Package errors provides main function on errors files
package errors

import (
	"bufio"
	"fmt"
)

// ScannerError is a function that checks if an error is nil or not during the scan of the file
func ScannerError(scanner *bufio.Scanner) int {
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
