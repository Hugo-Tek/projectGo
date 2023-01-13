package components

type Warehouse struct {
	Width, Height, Turns int
	Parcels              []Parcel
	Pallets              []Pallet
	Trucks               []Truck
	Grid                 [][]int
}

func InitWarehouse(width, height, turns int, parcels []Parcel, pallets []Pallet, trucks []Truck, grid [][]int) *Warehouse {
	return &Warehouse{
		Width:   width,
		Height:  height,
		Turns:   turns,
		Parcels: parcels,
		Pallets: pallets,
		Trucks:  trucks,
		Grid:    grid,
	}
}
