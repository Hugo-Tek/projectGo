package components

import (
	"fmt"
	"projectGo/src/errors"
	"strconv"
	"strings"
)

type Parcel struct {
	Name  string
	X, Y  int
	Color string
}

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
