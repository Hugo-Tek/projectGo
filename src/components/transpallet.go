package components

import (
	"projectGo/src/errors"
	"strconv"
)

type Pallet struct {
	Name      string
	X, Y      int
	Status    string
	Parcel    Parcel
	IsCharged bool
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
	p := Pallet{Name: parts[0], X: x, Y: y, Status: "WAIT", Parcel: Parcel{}, IsCharged: false}
	pallets = append(pallets, p)
	return pallets
}
