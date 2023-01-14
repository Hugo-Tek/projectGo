// Package components provides main function on components files
package components

// Warehouse struct
type Warehouse struct {
	Width, Height, Turns int
	Parcels              []Parcel
	PalletTruck          []PalletTruck
	Trucks               []Truck
	Grid                 [][]int
	NoMoreParcel         bool
}

// InitWarehouse func to init the warehouse
func InitWarehouse(width, height, turns int, parcels []Parcel, palletTruck []PalletTruck, trucks []Truck, grid [][]int) *Warehouse {
	return &Warehouse{
		Width:        width,
		Height:       height,
		Turns:        turns,
		Parcels:      parcels,
		PalletTruck:  palletTruck,
		Trucks:       trucks,
		Grid:         grid,
		NoMoreParcel: false,
	}
}
