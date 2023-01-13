package src

import (
	"projectGo/src/components"
)

func MovePallets(warehouse *components.Warehouse) {
	for i := 0; i < len(warehouse.Pallets); i++ {
		for j := 0; j < len(warehouse.Parcels); j++ {
			if ParcelNearPallet(warehouse.Pallets[i], warehouse.Parcels[j]) && !warehouse.Pallets[i].IsCharged {
				warehouse.Pallets[i].IsCharged = true
				warehouse.Pallets[i].Parcel = warehouse.Parcels[j]
				warehouse.Pallets[i].Status = "TAKE"
				return
			} else if warehouse.Pallets[i].IsCharged && warehouse.Pallets[i].Status == "WAIT" {
				warehouse.Pallets[i].Status = "WAIT"
			} else if warehouse.Pallets[i].IsCharged && warehouse.Pallets[i].Status != "LEAVE" && warehouse.Pallets[i].Status != "WAIT" {
				warehouse.Pallets[i].Status = "GO"
				MoveToTruck(warehouse, i)
				return
			}
		}
	}
}

func ParcelNearPallet(pallet components.Pallet, parcel components.Parcel) bool {
	return IsAdjacent(pallet.X, pallet.Y, parcel.X, parcel.Y)
}
