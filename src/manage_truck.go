package src

import (
	"projectGo/src/components"
)

func PalletIsInLoadingZone(pallet components.Pallet, truck components.Truck) bool {
	return IsAdjacent(pallet.X, pallet.Y, truck.X, truck.Y)
}

func IsLoading(pallet components.Pallet, truck components.Truck) string {
	if PalletIsInLoadingZone(pallet, truck) {
		return "chargé"
	}
	return "non chargé"
}
