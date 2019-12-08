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
	f := openFile(name)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return f
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

func getTotalFuelForMassAndFuel(name string) int {
	var masses []string = readMasses("input.txt")
	var fuel int = 0
	for i := range masses {
		var mass, err = strconv.Atoi(masses[i])
		if err != nil {
			panic(err)
		}
		var fuelRequired = getRequiredFuel(mass)
		fuel += fuelRequired + getFuelForFuel(fuelRequired)
	}
	return fuel
}

func getFuelForFuel(fuelMass int) int {
	var fuelRequired = getRequiredFuel(fuelMass)
	if fuelRequired <= 0 {
		return 0
	} else {
		return fuelRequired + getFuelForFuel(fuelRequired)
	}
}
