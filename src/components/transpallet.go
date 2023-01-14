// Package components provides main function on components files
package components

import (
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
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return palletTruck
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return palletTruck
	}
	p := PalletTruck{Name: parts[0], X: x, Y: y, Status: "WAIT", CurrentParcel: Parcel{}, IsCharged: false}
	palletTruck = append(palletTruck, p)
	return palletTruck
}
