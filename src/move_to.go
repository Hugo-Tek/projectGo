package src

import (
	"projectGo/src/components"
)

func MoveToTruck(warehouse *components.Warehouse, i int) {

	if warehouse.Pallets[i].X < warehouse.Trucks[0].X {
		warehouse.Pallets[i].X++
	} else if warehouse.Pallets[i].X > warehouse.Trucks[0].X {
		warehouse.Pallets[i].X--
	} else if warehouse.Pallets[i].Y < warehouse.Trucks[0].Y {
		warehouse.Pallets[i].Y++
	} else if warehouse.Pallets[i].Y > warehouse.Trucks[0].Y {
		warehouse.Pallets[i].Y--
	}
}

func MoveToParcel(warehouse *components.Warehouse) {
	if warehouse.Pallets[0].X < warehouse.Parcels[0].X {
		warehouse.Pallets[0].X++
	} else if warehouse.Pallets[0].X > warehouse.Parcels[0].X {
		warehouse.Pallets[0].X--
	} else if warehouse.Pallets[0].Y < warehouse.Parcels[0].Y {
		warehouse.Pallets[0].Y++
	} else if warehouse.Pallets[0].Y > warehouse.Parcels[0].Y {
		warehouse.Pallets[0].Y--
	}
}
