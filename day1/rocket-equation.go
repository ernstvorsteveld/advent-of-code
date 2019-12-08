package day1

import (
	"bufio"
	"os"
	"strconv"
)

func getRequiredFuel(mass int) int {
	return (mass / 3) - 2
}

func readMasses(name string) []string {
	f, _ := os.Open(name)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getTotalFuelRequired(name string) int {
	var masses []string = readMasses("input.txt")
	var fuel int = 0
	for i := range masses {
		var mass, _ = strconv.Atoi(masses[i])
		var fuelRequired = getRequiredFuel(mass)
		fuel += fuelRequired
	}
	return fuel
}

func getFuelForFuel(fuelMass int) int {
	var fuelRequired = getRequiredFuel(fuelMass)
	if fuelRequired <= 0 {
		return 0
	}
	return fuelRequired + getFuelForFuel(fuelRequired)
}

func getTotalFuelForMassAndFuel(name string) int {
	var masses []string = readMasses("input.txt")
	var fuel int = 0
	for _, mass := range masses {
		var mass, _ = strconv.Atoi(mass)
		var fuelRequired = getRequiredFuel(mass)
		fuel += fuelRequired + getFuelForFuel(fuelRequired)
	}
	return fuel
}

func getTotalFuelForMass(mass int, massFunc func(mass int) int, fuelFunc func(fuel int) int) (int, int) {
	var fuelForMass = massFunc(mass)
	var fuelForFuel = fuelFunc(fuelForMass)
	return fuelForMass, fuelForFuel
}
