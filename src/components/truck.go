package components

import (
	"projectGo/src/errors"
	"strconv"
)

type Truck struct {
	Name         string
	X, Y         int
	MaxLoad      int
	Availability int
	Status       string
}

func CreateTruck(parts []string, trucks []Truck) []Truck {
	ret := 0
	x, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	y, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	maxLoad, err := strconv.Atoi(parts[3])
	ret += errors.Check(err)
	availability, err := strconv.Atoi(parts[4])
	ret += errors.Check(err)
	Status := "WAITING"
	if ret != 0 {
		return trucks
	}
	t := Truck{Name: parts[0], X: x, Y: y, MaxLoad: maxLoad, Availability: availability, Status: Status}
	trucks = append(trucks, t)
	return trucks
}
