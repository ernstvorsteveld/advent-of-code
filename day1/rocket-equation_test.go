package day1

import (
	"testing"
)

func Test_should_get_required_fuel(t *testing.T) {
	var expected int = 654
	var fuelRequired int = getRequiredFuel("1969")
	if fuelRequired != expected {
		t.Errorf("Fuel calculation failed, got %d, expected %d.", fuelRequired, expected)
	}

	expected = 33583
	fuelRequired = getRequiredFuel("100756")
	if fuelRequired != expected {
		t.Errorf("Fuel calculation failed, got %d, expected %d.", fuelRequired, expected)
	}
}
