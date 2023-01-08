package components

import (
	"testing"
)

func TestCreatePackageSuccess(t *testing.T) {
	// Test valid input
	parts := []string{"p1", "1", "2", "green"}
	parcels := []Parcel{}
	parcels = CreatePackage(parts, parcels)
	if len(parcels) != 1 {
		t.Errorf("Expected length of parcels slice to be 1, got %d", len(parcels))
	}
	if parcels[0].Name != "p1" {
		t.Errorf("Expected Name to be p1, got %s", parcels[0].Name)
	}
	if parcels[0].X != 1 {
		t.Errorf("Expected X to be 1, got %d", parcels[0].X)
	}
	if parcels[0].Y != 2 {
		t.Errorf("Expected Y to be 2, got %d", parcels[0].Y)
	}
	if parcels[0].Color != "GREEN" {
		t.Errorf("Expected Color to be GREEN, got %s", parcels[0].Color)
	}
}

func TestCreatePackageInvalidX(t *testing.T) {
	// Test invalid input where X is not an integer
	parts := []string{"p1", "a", "2", "green"}
	parcels := []Parcel{}
	parcels = CreatePackage(parts, parcels)
	if len(parcels) != 0 {
		t.Errorf("Expected length of parcels slice to be 0, got %d", len(parcels))
	}
}

func TestCreatePackageInvalidY(t *testing.T) {
	// Test invalid input where Y is not an integer
	parts := []string{"p1", "1", "b", "green"}
	parcels := []Parcel{}
	parcels = CreatePackage(parts, parcels)
	if len(parcels) != 0 {
		t.Errorf("Expected length of parcels slice to be 0, got %d", len(parcels))
	}
}

func TestCreatePackageInvalidColor(t *testing.T) {
	// Test invalid input where color is not a valid color
	parts := []string{"p1", "1", "2", "invalid"}
	parcels := []Parcel{}
	parcels = CreatePackage(parts, parcels)
	if len(parcels) != 0 {
		t.Errorf("Expected length of parcels slice to be 0, got %d", len(parcels))
	}
}
