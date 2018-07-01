package ttt

type Easy struct{
	Piece	Piece
	Client	Client
}

func (comp Easy) GetPiece() Piece {
	return comp.Piece
}

func (comp Easy) GetMove(board Board) int {
	return board.blanks()[0]
}

type Medium struct{
	Piece	Piece
	Client	Client
}

func (comp Medium) GetPiece() Piece {
	return comp.Piece
}

func (comp Medium) GetMove(board Board) int {
	blanks := board.blanks()
	for _, pos := range blanks {
		tryBoard := make(Board, len(board))
		copy(tryBoard, board)
		if tryBoard.Mark(pos, comp.Piece).wonBy() == comp.Piece {
			return pos
		}
	}
	return blanks[0]
}
