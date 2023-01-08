package components

import (
	"projectGo/src/errors"
	"strconv"
)

type Pallet struct {
	Name string
	X, Y int
}

func CreatePallet(parts []string, pallets []Pallet) []Pallet {
	ret := 0
	x, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	y, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	if ret != 0 {
		return pallets
	}
	p := Pallet{Name: parts[0], X: x, Y: y}
	pallets = append(pallets, p)
	return pallets
}
