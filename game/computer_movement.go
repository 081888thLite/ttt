package game

type Easy struct{}

func (comp Easy) GetMove(client Client, board Board) int {
	return comp.getEasyMove(board)
}

func (comp Easy) getEasyMove(board Board) int {
	i := 0
	for i, e := range board {
		if e == Blank {
			return i
		}
	}
	return i
}
