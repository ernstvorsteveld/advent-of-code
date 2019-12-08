package day2

import (
	"fmt"
	"testing"
)

func Test_should_execute_day2_sum(t *testing.T) {
	expected := 19690720

	i, j := findValue(expected)

	var result []int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input := readOpcodes()

			input[1] = i
			input[2] = j
			result = execute(input)

			if expected == result[0] {
				fmt.Printf("same, while i == %d and j == %d\n", i, j)
				break
			}
		}
	}
	fmt.Printf("Found at position 0 the value: %d, at postion 1: %d and postion 2: %d\n", result[0], i, j)

	fmt.Printf("The answer is: %d\n", 100*i+j)
}
