package day1

import (
	"testing"
)

func Test_should_read_file(t *testing.T) {
	masses := readMasses("input.txt")
	if len(masses) != 100 {
		t.Errorf("Did not find 100 lines, but %d.", len(masses))
	}
}

func Test_should_get_required_fuel(t *testing.T) {
	var expected int = 654
	var fuelRequired int = getRequiredFuel(1969)
	if fuelRequired != expected {
		t.Errorf("Fuel calculation failed, got %d, expected %d.", fuelRequired, expected)
	}

	expected = 19890
	fuelRequired = getRequiredFuel(59677)
	if fuelRequired != expected {
		t.Errorf("Fuel calculation failed, got %d, expected %d.", fuelRequired, expected)
	}
}

func Test_should_get_total_required_fuel(t *testing.T) {
	required := getTotalFuelRequired("input.txt")
	if required != 3375962 {
		t.Errorf("Fuel calculation failed, got %d, expected %d.", required, 0)
	}

	var forFuel = getFuelForFuel(36)
	if forFuel != 11 {
		t.Errorf("Fuel for fuel calculation failed, got %d, expected %d.", forFuel, 11)
	}

	forFuel = getFuelForFuel(3375962)
	if forFuel != 1687937 {
		t.Errorf("Fuel for fuel calculation failed, got %d, expected %d.", forFuel, 1687937)
	}
}
