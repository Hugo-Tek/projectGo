// Package errors provides main function on errors files
package errors

import "fmt"

// Check is a function that checks if an error is nil or not
func Check(err error) int {
	if err != nil {
		fmt.Println(err)
		fmt.Println("😱")
		return 1
	}
	return 0
}
