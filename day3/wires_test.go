package day3

import (
	"testing"
)

func Test_should_execute_day2_assingment(t *testing.T) {
	var dim = 10
	var bp *SwitchBoard = createSwitchBoard(dim)

	if bp.dim != dim {
		t.Errorf("Dimension of the array in the crossing is not %d, but %d.", dim, bp.dim)
	}

	for i := range bp.crossings {
		var crossing *Crossings = new(Crossings)
		crossing.x = i
		crossing.y = i
		bp.crossings[i][i] = crossing
	}
}
