// Package src provides main function on src files
package src

import (
	"fmt"
	"projectGo/src/components"
)

func checkMaxLoad(truck *components.Truck, weight int) bool {
	return truck.CurrentLoad+weight <= truck.MaxLoad
}

// Charging func to charge the parcel from the pallettruck to the truck
func Charging(warehouse *components.Warehouse, palletTruck *components.PalletTruck, truck *components.Truck) {
	var weight int

	switch palletTruck.CurrentParcel.Color {
	case "YELLOW":
		weight = 100
	case "GREEN":
		weight = 200
	case "BLUE":
		weight = 500
	}

	if checkMaxLoad(truck, weight) {
		fmt.Println(palletTruck.Name, "LEAVE", palletTruck.CurrentParcel.Name, palletTruck.CurrentParcel.Color)

		truck.CurrentLoad += weight
		palletTruck.IsCharged = false
		palletTruck.Status = "LEAVE"
		palletTruck.CurrentParcel = components.Parcel{}
	} else {
		truck.Status = "GONE"
		palletTruck.Status = "WAIT"
		fmt.Println(palletTruck.Name, "WAIT")
	}
}

func manageTruck(warehouse *components.Warehouse, truck *components.Truck) {
	if truck.Status != "GONE" && truck.CurrentLoad != truck.MaxLoad && components.CountingParcels(warehouse) == 0 {
		fmt.Printf("%s WAITING %d/%d\n", truck.Name, truck.CurrentLoad, truck.MaxLoad)
	} else {
		fmt.Printf("%s GONE %d/%d\n", truck.Name, truck.CurrentLoad, truck.MaxLoad)
		truck.Status = "GONE"
		truck.IsGone++
		if truck.IsGone == truck.DayGone {
			truck.Status = "WAITING"
			truck.CurrentLoad = 0
			truck.IsGone = 0
		}
	}
}

// ManageTrucks func to manage the truck if GONE or WAITING
func ManageTrucks(warehouse *components.Warehouse) {
	for i := 0; i < len(warehouse.Trucks); i++ {
		manageTruck(warehouse, &warehouse.Trucks[i])
	}
}
