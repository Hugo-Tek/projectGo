package src

import (
	"projectGo/src/components"
)

func IndexToRemove(parcels []components.Parcel, target components.Parcel) int {
	for i, parcel := range parcels {
		if parcel == target {
			return i
		}
	}
	return -1
}
