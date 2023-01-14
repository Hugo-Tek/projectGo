// Package components provides main function on components files
package components

import (
	"projectGo/src/errors"
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
	ret := 0
	x, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	y, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	maxLoad, err := strconv.Atoi(parts[3])
	ret += errors.Check(err)
	DayGone, err := strconv.Atoi(parts[4])
	ret += errors.Check(err)
	Status := "WAITING"
	Load := 0
	if ret != 0 {
		return trucks
	}
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
