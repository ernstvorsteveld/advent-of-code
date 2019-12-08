package day2

import (
	"fmt"
	"testing"
)

func Test_should_read_file(t *testing.T) {
	input := readAndPrint()
	fmt.Printf("Read: %+q", input)

	if input[1] != 0 {
		t.Errorf("The array does not contain %d at postion %d, but %d.", 0, 1, input[1])
	}
}
