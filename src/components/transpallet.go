// Package components provides main function on components files
package components

import (
	"projectGo/src/errors"
	"strconv"
)

// PalletTruck struct
type PalletTruck struct {
	Name           string
	X, Y           int
	Status         string
	CurrentParcel  Parcel
	OnHisWayParcel Parcel
	IsCharged      bool
	IsOnHisWay     bool
}

// CreatePallet func to init/create pallet truck
func CreatePallet(parts []string, palletTruck []PalletTruck) []PalletTruck {
	ret := 0
	x, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	y, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	if ret != 0 {
		return palletTruck
	}
	p := PalletTruck{Name: parts[0], X: x, Y: y, Status: "WAIT", CurrentParcel: Parcel{}, IsCharged: false}
	palletTruck = append(palletTruck, p)
	return palletTruck
}
