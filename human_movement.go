package ttt

import "strconv"

const PromptForMove = "To make a move, enter a number corresponding\n to an open board position and press 'return':\n"

type Human struct{
	Piece    Piece
	Client	Client
}

func (human Human) GetPiece() Piece {
	return human.Piece
}

func (human *Human) GetMove(board Board) int {
	human.Client.Write(PromptForMove)
	human.Client.Read()
	m := human.Client.GetLastRead()
	choice, _ := strconv.Atoi(m)
	return choice
}
