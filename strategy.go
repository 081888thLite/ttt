package ttt

type Piece string

type Easy struct {
	piece  Piece
	Client Client
}

func (e *Easy) GetPiece() Piece {
	return e.piece
}

func (e *Easy) GetMove(b Board, opp Player) int {
	return b.blanks()[0]
}

type Medium struct {
	piece  Piece
	Client Client
}

func (m *Medium) GetPiece() Piece {
	return m.piece
}

func (m *Medium) GetMove(b Board, opp Player) int {
	bw := BlockOrWin(b)
	if bw == -1 {
		return b.blanks()[0]
	} else {
		return bw
	}
}

type Hard struct {
	piece  Piece
	Client Client
}

func (h *Hard) GetPiece() Piece {
	return h.piece
}

func (h *Hard) GetMove(b Board, opp Player) int {
	if bw := BlockOrWin(b); bw != -1 {
		return bw
	}
	mm := new(Minimax)
	nb := h.ReplaceBoard(b)[:]
	mv := mm.minimax(nb, X, 0)
	return mv
}

func (h *Hard) ReplaceBoard(b Board) Board {
	nb := NewBoard(9)
	for i, e := range b {
		if e == h.piece {
			nb[i] = X
		} else if e == Blank {
			nb[i] = NoOne
		} else {
			nb[i] = O
		}
	}
	return nb
}

type Human struct {
	piece   Piece
	Client  Client
	console Console
}

func (hu *Human) GetPiece() Piece {
	return hu.piece
}

func (hu *Human) GetMove(b Board, opp Player) int {
	mv, err := hu.console.getHumanMove()
	open := b.blanks()
	var valid bool
	for _, e := range open {
		if mv == e {
			valid = true
		}
	}
	if mv > len(b)-1 || mv < 0 || err != nil || !valid {
		hu.console.Write(MoveError)
		mv = hu.GetMove(b, opp)
	}
	return mv
}

func BlockOrWin(b Board) int {
	for i := 0; i < 8; i++ {
		if b[WinConditions[i][0]] == b[WinConditions[i][1]] &&
			b[WinConditions[i][0]] != NoOne &&
			b[WinConditions[i][2]] == Blank {
			return WinConditions[i][2]
		} else if b[WinConditions[i][1]] == b[WinConditions[i][2]] &&
			b[WinConditions[i][1]] != NoOne &&
			b[WinConditions[i][0]] == Blank {
			return WinConditions[i][0]
		} else if b[WinConditions[i][0]] == b[WinConditions[i][2]] &&
			b[WinConditions[i][0]] != NoOne &&
			b[WinConditions[i][1]] == Blank {
			return WinConditions[i][1]
		}
	}
	return -1
}
