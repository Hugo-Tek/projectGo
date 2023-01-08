package errors

import (
	"bufio"
	"fmt"
)

func ScannerError(scanner *bufio.Scanner) int {
	// Vérifie qu'il n'y a pas eu d'erreur pendant la lecture du fichier
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
