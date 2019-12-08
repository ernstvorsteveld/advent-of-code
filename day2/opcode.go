package day2

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func readAndPrint() []int {
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
