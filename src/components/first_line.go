package components

import (
	"bufio"
	"fmt"
	"projectGo/src/errors"
	"strconv"
	"strings"
)

func ParseFirstLine(scanner *bufio.Scanner) (int, int, int, int) {
	line := scanner.Text()
	parts := strings.Split(line, " ")
	ret := 0
	if len(parts) != 3 {
		fmt.Println("Erreur de format: première ligne doit avoir 3 éléments")
		ret = 1
		return 0, 0, 0, ret
	}
	width, err := strconv.Atoi(parts[0])
	ret += errors.Check(err)
	height, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	turns, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	if ret != 0 {
		return 0, 0, 0, ret
	}
	return width, height, turns, ret
}
