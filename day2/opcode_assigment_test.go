package day2

import (
	"fmt"
	"testing"
)

func Test_should_execute_day2_assingment(t *testing.T) {
	input := readOpcodes()
	output := execute(input)
	fmt.Printf("Found at position 0 the value: %d\n", output[0])
}
