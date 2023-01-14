// Package src provides main function on src files
package src

import (
	"projectGo/src/components"
)

// FillGrid func to fill the grid
func FillGrid(w *components.Warehouse) *components.Warehouse {
	// Initialize all squares in the grid as empty
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Grid[i][j] != -1 {
				w.Grid[i][j] = 0
			}
		}
	}

	// Fill the grid with parcels
	for _, parcel := range w.Parcels {
		if w.Grid[parcel.Y][parcel.X] != -1 {
			w.Grid[parcel.Y][parcel.X] = 1
		}
	}

	// Fill the grid with PalletTruck
	for i := range w.PalletTruck {
		pallet := w.PalletTruck[i]
		w.Grid[pallet.Y][pallet.X] = 2
	}

	// Fill the grid with trucks
	for i := range w.Trucks {
		truck := w.Trucks[i]
		w.Grid[truck.Y][truck.X] = 3
	}
	return w
}
