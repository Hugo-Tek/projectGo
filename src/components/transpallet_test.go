// Package components provides main function on components files
package components

import (
	"testing"
)

func TestCreatePalletErrorPart1(t *testing.T) {
	// Test conversion error for parts[1]
	parts := []string{"P1", "invalid", "20"}
	pallets := []PalletTruck{}
	resultPalletTruck := CreatePallet(parts, pallets)
	if len(resultPalletTruck) != len(pallets) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPalletTruck, pallets)
	}
}

func TestCreatePalletErrorPart2(t *testing.T) {
	// Test conversion error for parts[2]
	parts := []string{"P1", "10", "invalid"}
	pallets := []PalletTruck{}
	resultPalletTruck := CreatePallet(parts, pallets)
	if len(resultPalletTruck) != len(pallets) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPalletTruck, pallets)
	}
}

func arePalletsEqual(p1, p2 PalletTruck) bool {
	return p1.Name == p2.Name && p1.X == p2.X && p1.Y == p2.Y && p1.Status == p2.Status && p1.CurrentParcel == p2.CurrentParcel && p1.OnHisWayParcel == p2.OnHisWayParcel && p1.IsCharged == p2.IsCharged && p1.IsOnHisWay == p2.IsOnHisWay
}

func TestCreatePalletSuccess(t *testing.T) {
	parts := []string{"P1", "10", "20"}
	pallets := []PalletTruck{}
	expectedPalletTruck := []PalletTruck{{"P1", 10, 20, "WAIT", Parcel{}, Parcel{}, false, false}}
	resultPalletTruck := CreatePallet(parts, pallets)

	if len(resultPalletTruck) != len(expectedPalletTruck) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPalletTruck, expectedPalletTruck)
	}
	for i := range resultPalletTruck {
		if !arePalletsEqual(resultPalletTruck[i], expectedPalletTruck[i]) {
			t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPalletTruck, expectedPalletTruck)
		}
	}
}
