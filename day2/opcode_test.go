package day2

import (
	"fmt"
	"testing"
)

func Test_should_read_file(t *testing.T) {
	input := readOpcodes()
	//print(input)
	if input[1] != 12 {
		t.Errorf("The array does not contain %d at postion %d, but %d.", 0, 1, input[1])
	}
}

func Test_should_work_for_examples(t *testing.T) {
	var c = executeTest([]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	if c != 0 {
		t.Errorf("Input did not result in expected output, value %d.", c)
	}

	c = executeTest([]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	if c != 0 {
		t.Errorf("Input did not result in expected output, value %d.", c)
	}

	c = executeTest([]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	if c != 0 {
		t.Errorf("Input did not result in expected output, value %d.", c)
	}

	c = executeTest([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
	if c != 0 {
		t.Errorf("Input did not result in expected output, value %d.", c)
	}
}

func executeTest(input []int, expected []int) int {
	output := execute(input)
	c := compare(output, expected)
	if c != 0 {
		print(output)
		print(expected)
	}
	return c
}

func compare(left []int, right []int) int {
	if len(left) > len(right) {
		return -1
	}

	if len(left) < len(right) {
		return 1
	}

	for i, v := range left {
		if v != right[i] {
			return 2
		}
	}
	return 0
}

func printM(m string, opcodes []int) {
	fmt.Printf("%s", m)
	print(opcodes)
}

func print(opcodes []int) {
	for _, v := range opcodes {
		fmt.Printf("%d, ", v)
	}
	fmt.Printf("\n")
}
