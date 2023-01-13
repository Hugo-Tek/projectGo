package components

import (
	"reflect"
	"testing"
)

func TestCreatePalletErrorPart1(t *testing.T) {
	// Test conversion error for parts[1]
	parts := []string{"P1", "invalid", "20"}
	pallets := []Pallet{}
	resultPallets := CreatePallet(parts, pallets)
	if !reflect.DeepEqual(resultPallets, pallets) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPallets, pallets)
	}
}

func TestCreatePalletErrorPart2(t *testing.T) {
	// Test conversion error for parts[2]
	parts := []string{"P1", "10", "invalid"}
	pallets := []Pallet{}
	resultPallets := CreatePallet(parts, pallets)
	if !reflect.DeepEqual(resultPallets, pallets) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPallets, pallets)
	}
}

func TestCreatePalletSuccess(t *testing.T) {
	parts := []string{"P1", "10", "20"}
	pallets := []Pallet{}
	expectedPallets := []Pallet{{"P1", 10, 20, "WAIT", Parcel{}, false}}
	resultPallets := CreatePallet(parts, pallets)
	if !reflect.DeepEqual(resultPallets, expectedPallets) {
		t.Errorf("CreatePallet(%v, %v) = %v, expected %v", parts, pallets, resultPallets, expectedPallets)
	}
}
