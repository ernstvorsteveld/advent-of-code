package day2

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func readOpcodes() []int {
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)

	var result []string
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, record...)
	}

	var ints []int

	for _, r := range result {
		var i, _ = strconv.Atoi(r)
		ints = append(ints, i)
	}
	return ints
}

func execute(opcodes []int) []int {
	var p = 0
	var operation int = 99

	for {
		operation = opcodes[p]
		switch operation {
		case 99:
			return opcodes
		case 1:
			opcodes, p = add(opcodes, p)
		case 2:
			opcodes, p = mult(opcodes, p)
		}
	}
	return opcodes
}

func add(opcodes []int, p int) ([]int, int) {
	opcodes[opcodes[p+3]] = opcodes[opcodes[p+1]] + opcodes[opcodes[p+2]]
	return opcodes, p + 4
}

func mult(opcodes []int, p int) ([]int, int) {
	opcodes[opcodes[p+3]] = opcodes[opcodes[p+1]] * opcodes[opcodes[p+2]]
	return opcodes, p + 4
}

func findValue(expected int) (int, int) {
	var result []int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input := readOpcodes()

			input[1] = i
			input[2] = j
			result = execute(input)

			if expected == result[0] {
				return i, j
			}
		}
	}
	return 99, 99
}
