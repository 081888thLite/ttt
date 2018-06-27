package game

import (
	"strconv"
)

const PromptForMove = "To make a move, enter a number corresponding to an open board position and press 'return':"

type Human struct{}

func (human Human) getMove(ui Client, board Board) int {
	ui.Send(PromptForMove)
	entered := ui.Receive()
	move, _ := strconv.Atoi(entered)
	//validate move entered is g2g
	return move
}
