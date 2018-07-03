package ttt

type Piece string

type Player interface {
	GetMove(board Board, opp Player) int
	GetPiece() Piece
}

type Easy struct {
	piece  Piece
	Client Client
}

func (comp *Easy) GetPiece() Piece {
	return comp.piece
}

func (comp *Easy) GetMove(board Board, opp Player) int {
	return board.blanks()[0]
}

type Medium struct {
	piece  Piece
	Client Client
}

func (comp *Medium) GetPiece() Piece {
	return comp.piece
}

func (comp *Medium) GetMove(board Board, opp Player) int {
	return board.blanks()[3]
}

type Hard struct {
	piece  Piece
	Client Client
}

func (comp *Hard) GetPiece() Piece {
	return comp.piece
}

func (comp *Hard) GetMove(board Board, opp Player) int {
	mm := new(Minimax)
	mm.SetCaller(*comp)
	choice := mm.minimax(board, [2]Player{comp, opp})
	return choice
}

type Human struct {
	piece  Piece
	Client Client
}

func (human *Human) GetPiece() Piece {
	return human.piece
}

func (human *Human) GetMove(board Board, opp Player) int {
	ui := &Console{}
	return ui.getHumanMove()
}
