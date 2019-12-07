package day1

import (
	"bufio"
	"os"
	"strconv"
)

func getRequiredFuel(mass string) int {
	i, err := strconv.Atoi(mass)
	if err != nil {
		panic(err)
	}
	return (i / 3) - 2
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
	var totalMass int = 0
	for i := range masses {
		var mass = masses[i]
		var fuelRequired = getRequiredFuel(mass)
		totalMass += fuelRequired
	}
	return totalMass
}
