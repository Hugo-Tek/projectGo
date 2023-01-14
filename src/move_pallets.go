// Package src provides main function on src files
package src

import (
	"fmt"
	"projectGo/src/components"
)

func notChargedPallet(warehouse *components.Warehouse, palletTruck *components.PalletTruck) {
	// chercher la palette la plus proche
	x, y, dist := components.FindClosestOne(warehouse.Grid, palletTruck.X, palletTruck.Y, 1)
	if dist >= 2147483647 {
		// wait ou move to empty si quelque chose a coté
		palletTruck.Status = "WAIT"
		MoveIfNotEmpty(warehouse, palletTruck)
		return
	}
	targetedParcel := *components.DetectParcel(x, y, warehouse)

	// SI SANS PALETTE A ALLER CHERCHER
	if dist == 1 {
		// TAKE
		palletTruck.Status = "TAKE"
		palletTruck.CurrentParcel = targetedParcel
		palletTruck.IsCharged = true

		i := components.IndexToRemove(warehouse.Parcels, targetedParcel)
		warehouse.Parcels = append(warehouse.Parcels[:i], warehouse.Parcels[i+1:]...)

		palletTruck.OnHisWayParcel = components.Parcel{}
		palletTruck.IsOnHisWay = false
		warehouse.Grid[y][x] = 0
		fmt.Println(palletTruck.Name, "TAKE", targetedParcel.Name, targetedParcel.Color)
	} else {
		if !palletTruck.IsOnHisWay {
			// ajouter que je suis en route pour une palette et donc non disponible (onHisWay)
			palletTruck.IsOnHisWay = true
			palletTruck.OnHisWayParcel = targetedParcel
			// transformer la palette la plus proche en un autre numero
			warehouse.Grid[y][x] = -1
		}
		// MOVE
		MoveTo(warehouse, palletTruck, x, y)
	}
}

func movePalletTruck(warehouse *components.Warehouse, palletTruck *components.PalletTruck) {
	// ICI ON RECOIT LE COLIS LE PLUS PROCHE ET LA DISTANCE

	if !palletTruck.IsCharged {
		notChargedPallet(warehouse, palletTruck)
	} else if palletTruck.IsCharged {
		x, y, dist := components.FindClosestOne(warehouse.Grid, palletTruck.X, palletTruck.Y, 3)
		truckToGo := components.DetectTruck(x, y, warehouse)
		// SI IL EST A COTÉ DU CAMION
		if dist == 1 {
			Charging(warehouse, palletTruck, truckToGo)
		} else {
			// se deplacer vers le camion le plus proche.
			MoveTo(warehouse, palletTruck, x, y)
		}
	} else {
		palletTruck.Status = "WAIT"
	}
}

// MovePalletTrucks func to manage the pallet truck
func MovePalletTrucks(warehouse *components.Warehouse) {
	for i := 0; i < len(warehouse.PalletTruck); i++ {
		movePalletTruck(warehouse, &warehouse.PalletTruck[i])
	}
}
