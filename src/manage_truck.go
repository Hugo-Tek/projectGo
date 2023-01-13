package src

import (
	"projectGo/src/components"
)

func PalletIsInLoadingZone(pallet components.Pallet, truck components.Truck) bool {
	return IsAdjacent(pallet.X, pallet.Y, truck.X, truck.Y)
}

func CheckMaxLoad(truck components.Truck, weight int) bool {
	return truck.CurrentLoad+weight < truck.MaxLoad
}

func Charging(parcel components.Parcel, truck components.Truck) {
	var weight int = 0

	switch parcel.Color {
	case "YELLOW":
		weight = 100
	case "GREEN":
		weight = 200
	case "BLUE":
		weight = 500
	}

	if CheckMaxLoad(truck, weight) {
		truck.CurrentLoad += weight
	} else {
		truck.Status = "GONE"
	}
}

func IsLoading(pallet components.Pallet, truck components.Truck) {
	if PalletIsInLoadingZone(pallet, truck) && pallet.IsCharged {
		if pallet.Status == "WAIT" {
			pallet.Status = "LEAVE"
			Charging(pallet.Parcel, truck)
		} else if pallet.Status == "LEAVE" {
			pallet.Parcel = components.Parcel{}
			pallet.IsCharged = false
			pallet.Status = "WAIT"
		}
	}
}

func ManageTruck(pallet components.Pallet, truck components.Truck, gone int) {
	if truck.Status != "GONE" {
		IsLoading(pallet, truck)
	} else {
		gone++
		if gone == truck.Availability {
			truck.Status = "WAITING"
			gone = 0
		}
	}
}
