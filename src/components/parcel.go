// Package components provides main function on components files
package components

import (
	"fmt"
	"projectGo/src/errors"
	"strconv"
	"strings"
)

// Parcel struct
type Parcel struct {
	Name  string
	X, Y  int
	Color string
}

// CreatePackage func to initialize/create Parcels
func CreatePackage(parts []string, parcels []Parcel) []Parcel {
	ret := 0
	x, err := strconv.Atoi(parts[1])
	ret += errors.Check(err)
	y, err := strconv.Atoi(parts[2])
	ret += errors.Check(err)
	if ret != 0 {
		return parcels
	}
	color := strings.ToUpper(parts[3])
	if color != "GREEN" && color != "YELLOW" && color != "BLUE" {
		fmt.Println("Erreur de format: couleur de colis non valide")
	} else {
		p := Parcel{Name: parts[0], X: x, Y: y, Color: color}
		parcels = append(parcels, p)
	}
	return parcels
}

// DetectParcel func to find the nearest parcel
func DetectParcel(x int, y int, warehouse *Warehouse) *Parcel {
	for i := range warehouse.Parcels {
		if warehouse.Parcels[i].X == x && warehouse.Parcels[i].Y == y {
			return &warehouse.Parcels[i]
		}
	}
	return nil
}

// CountingParcels func to count the number of parcel left
func CountingParcels(warehouse *Warehouse) int {
	count := 0

	for i := 0; i < len(warehouse.PalletTruck); i++ {
		if warehouse.PalletTruck[i].IsCharged {
			count++
		}
	}
	return count + len(warehouse.Parcels)
}

// IndexToRemove func to return the index of the element to remove
func IndexToRemove(parcels []Parcel, target Parcel) int {
	for i, parcel := range parcels {
		if parcel == target {
			return i
		}
	}
	return -1
}
