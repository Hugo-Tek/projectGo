package components

import (
	"reflect"
	"testing"
)

func TestCreateTruckValidInput(t *testing.T) {
	// Test with a valid input slice
	parts := []string{"truck1", "10", "20", "30", "40"}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if len(result) != len(trucks)+1 {
		t.Errorf("Expected result to have length %d, got %d", len(trucks)+1, len(result))
	}
}

func TestCreateTruckInvalidX(t *testing.T) {
	// Test with an input slice containing a non-integer string
	parts := []string{"truck1", "invalid", "20", "30", "40"}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if !reflect.DeepEqual(result, trucks) {
		t.Errorf("Expected result to be %v, got %v", trucks, result)
	}
}

func TestCreateTruckInvalidY(t *testing.T) {
	// Test with an input slice containing a non-integer string
	parts := []string{"truck1", "10", "invalid", "30", "40", "", ""}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if !reflect.DeepEqual(result, trucks) {
		t.Errorf("Expected result to be %v, got %v", trucks, result)
	}
}

func TestCreateTruckInvalidMaxLoad(t *testing.T) {
	// Test with an input slice containing a non-integer string
	parts := []string{"truck1", "10", "20", "invalid", "40"}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if !reflect.DeepEqual(result, trucks) {
		t.Errorf("Expected result to be %v, got %v", trucks, result)
	}
}

func TestCreateTruckInvalidAvailability(t *testing.T) {
	// Test with an input slice containing a non-integer string
	parts := []string{"truck1", "10", "20", "30", "invalid"}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if !reflect.DeepEqual(result, trucks) {
		t.Errorf("Expected result to be %v, got %v", trucks, result)
	}
}

func TestCreateTruckValidInputValues(t *testing.T) {
	// Test with a valid input slice and check truck fields
	parts := []string{"truck1", "10", "20", "30", "40"}
	trucks := []Truck{}
	result := CreateTruck(parts, trucks)
	if result[0].Name != parts[0] || result[0].X != 10 || result[0].Y != 20 || result[0].MaxLoad != 30 || result[0].Availability != 40 || result[0].Status != "WAITING" || result[0].CurrentLoad != 0 {
		t.Errorf("Expected truck fields to be {%s, %d, %d, %d, %d, %s}, got {%s, %d, %d, %d, %d, %s}", parts[0], 10, 20, 30, 40, "WAITING", result[0].Name, result[0].X, result[0].Y, result[0].MaxLoad, result[0].Availability, result[0].Status)
	}
}
