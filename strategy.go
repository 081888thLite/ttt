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
	return b.blanks()[3]
}

type Hard struct {
	piece  Piece
	Client Client
}

func (h *Hard) GetPiece() Piece {
	return h.piece
}

func (h *Hard) GetMove(b Board, opp Player) int {
	mm := new(Minimax)
	mm.SetCaller(h)
	if len(b.blanks()) == len(b) {
		opp = &Easy{piece: O}
	}
	pSet := &[2]Player{h, opp}
	mv := mm.minimax(b, pSet)
	return mv
}

type Human struct {
	piece  Piece
	Client Client
}

func (hu *Human) GetPiece() Piece {
	return hu.piece
}

func (hu *Human) GetMove(b Board, opp Player) int {
	c := &Console{}
	mv, err := c.getHumanMove()
	open := b.blanks()
	var valid bool
	for _, e := range open {
		if mv == e {
			valid = true
		}
	}
	if mv > len(b) - 1 || mv < 0 || err != nil || !valid {
		c.Write(MoveError)
		mv = hu.GetMove(b, opp)
	}
	return mv
}
