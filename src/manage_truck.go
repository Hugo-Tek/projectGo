package src

import (
	"fmt"
	"projectGo/src/components"
)

func PalletIsInLoadingZone(pallet components.Pallet, truck components.Truck) bool {
	return IsAdjacent(pallet.X, pallet.Y, truck.X, truck.Y)
}

func CheckMaxLoad(truck components.Truck, weight int) bool {
	return truck.CurrentLoad+weight <= truck.MaxLoad
}

func Charging(warehouse *components.Warehouse, i int) {
	var weight int = 0

	switch warehouse.Pallets[i].Parcel.Color {
	case "YELLOW":
		weight = 100
	case "GREEN":
		weight = 200
	case "BLUE":
		weight = 500
	}

	if CheckMaxLoad(warehouse.Trucks[0], weight) {
		warehouse.Trucks[0].CurrentLoad += weight
		warehouse.Pallets[i].Status = "LEAVE"
	} else {
		warehouse.Trucks[0].Status = "GONE"
		warehouse.Trucks[0].CurrentLoad = 0
		warehouse.Pallets[i].Status = "WAIT"
	}
}

func IsLoading(warehouse *components.Warehouse) {
	for i := 0; i < len(warehouse.Pallets); i++ {
		if PalletIsInLoadingZone(warehouse.Pallets[i], warehouse.Trucks[0]) && warehouse.Pallets[i].IsCharged && warehouse.Trucks[0].Status != "GONE" {
			if warehouse.Pallets[i].Status == "GO" || warehouse.Pallets[i].Status == "WAIT" {
				fmt.Println("bite")
				Charging(warehouse, i)
			} else if warehouse.Pallets[i].Status == "LEAVE" {
				warehouse.Pallets[i].Parcel = components.Parcel{}
				warehouse.Pallets[i].IsCharged = false
				warehouse.Pallets[i].Status = "WAIT"
			}
		}
	}
}

func ManageTruck(warehouse *components.Warehouse, gone int) int {
	if warehouse.Trucks[0].Status != "GONE" {
		IsLoading(warehouse)
	} else {
		gone++
		if gone == warehouse.Trucks[0].Availability {
			warehouse.Trucks[0].Status = "WAITING"
			gone = 0
		}
	}
	return gone
}
