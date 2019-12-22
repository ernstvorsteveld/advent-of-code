package day3

type Crossings struct {
	x, y  int
	wires []int
}

type SwitchBoard struct {
	crossings [][]*Crossings
	dim       int
}

func extend(board SwitchBoard) SwitchBoard {
	board.crossings = addCrossing(board)
	board.dim += 1
	return board
}

func addCrossing(board SwitchBoard) [][]*Crossings {
	newCrossings := make([][]*Crossings, board.dim+1)
	newDim := board.dim + 1

	for i := range board.crossings {
		newCrossings[i] = make([]*Crossings, newDim)
		for j := range board.crossings[i] {
			newCrossings[i][j] = board.crossings[i][j]
		}
	}

	return newCrossings
}

func createSwitchBoard(dim int) *SwitchBoard {
	var bp *SwitchBoard
	bp = new(SwitchBoard)
	bp.dim = dim
	bp.crossings = make([][]*Crossings, dim)
	return bp
}
