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

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	f.Close()
	return fileTextLines
}

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return f
}

func getTotalFuelRequired(name string) int {
	var masses []string = readMasses("input.txt")
	s1 := masses.([]string)
	var totalMass int = 0
	for mass := range masses {
		s := mass.(string)
		totalMass += getRequiredFuel(mass)
	}
	return totalMass
}
