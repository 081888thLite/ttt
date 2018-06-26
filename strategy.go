package ttt

type Piece string

type Player interface {
	GetMove(board Board, opp Player) int
	GetPiece() Piece
}

type Easy struct {
	Piece    Piece
	Client   Client
}

func (comp *Easy) GetPiece() Piece {
	return comp.Piece
}

func (comp *Easy) GetMove(board Board, opp Player) int {
	return board.blanks()[0]
}

type Medium struct {
	Piece  Piece
	Client Client
}

func (comp *Medium) GetPiece() Piece {
	return comp.Piece
}

func (comp *Medium) GetMove(board Board, opp Player) int {
	return board.blanks()[3]
}

type Hard struct {
	Piece    Piece
	Client   Client
}

func (comp *Hard) GetPiece() Piece {
	return comp.Piece
}

func (comp *Hard) GetMove(board Board, opp Player) int {
	mm := new ( Minimax )
	choice := mm.minimax(board[:], [2]Player{comp,opp})
	return choice
}

type Human struct {
	Piece  Piece
	Client Client
}

func (human *Human) GetPiece() Piece {
	return human.Piece
}

func (human *Human) GetMove(board Board, opp Player) int {
	ui := &Console{}
	return ui.getHumanMove()
}
