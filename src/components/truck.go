// Package components provides main function on components files
package components

import (
	"strconv"
)

// Truck struct of truck
type Truck struct {
	Name        string
	X, Y        int
	MaxLoad     int
	DayGone     int
	IsGone      int
	Status      string
	CurrentLoad int
}

// CreateTruck func to initialise trucks
func CreateTruck(parts []string, trucks []Truck) []Truck {
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return trucks
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return trucks
	}
	maxLoad, err := strconv.Atoi(parts[3])
	if err != nil {
		return trucks
	}
	DayGone, err := strconv.Atoi(parts[4])
	if err != nil {
		return trucks
	}
	Status := "WAITING"
	Load := 0
	t := Truck{Name: parts[0], X: x, Y: y, MaxLoad: maxLoad, DayGone: DayGone, IsGone: 0, Status: Status, CurrentLoad: Load}
	trucks = append(trucks, t)
	return trucks
}

// DetectTruck func to find the nearest truck
func DetectTruck(x int, y int, warehouse *Warehouse) *Truck {
	for i := range warehouse.Trucks {
		if warehouse.Trucks[i].X == x && warehouse.Trucks[i].Y == y {
			return &warehouse.Trucks[i]
		}
	}
	return nil
}
