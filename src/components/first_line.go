// Package components provides main function on components files
package components

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// ParseFirstLine func to parse first line
func ParseFirstLine(scanner *bufio.Scanner) (w int, h int, t int, r int) {
	line := scanner.Text()
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		fmt.Println("Erreur de format: première ligne doit avoir 3 éléments")
		return 0, 0, 0, 1
	}
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("Erreur de format: première ligne doit avoir 3 entiers")
		return 0, 0, 0, 1
	}
	height, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Erreur de format: première ligne doit avoir 3 entiers")
		return 0, 0, 0, 1
	}
	turns, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("Erreur de format: première ligne doit avoir 3 entiers")
		return 0, 0, 0, 1
	}
	return width, height, turns, 0
}
