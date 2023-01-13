package src

import (
	"projectGo/src/components"
)

func FillGrid(w *components.Warehouse) *components.Warehouse {
	// Initialize all squares in the grid as empty
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			w.Grid[i][j] = 0
		}
	}

	// Fill the grid with parcels
	for _, parcel := range w.Parcels {
		w.Grid[parcel.Y][parcel.X] = 1
	}

	// Fill the grid with pallets
	for _, pallet := range w.Pallets {
		w.Grid[pallet.Y][pallet.X] = 2
	}

	// Fill the grid with trucks
	for _, truck := range w.Trucks {
		w.Grid[truck.Y][truck.X] = 3
	}
	return w
}
