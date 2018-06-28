package game

import (
	"fmt"
	"strconv"
)

const PromptForMove = "To make a move, enter a number corresponding to an open board position and press 'return':"

type Human struct{}

func (human Human) GetMove(ui Client, board Board) int {
	ui.Write(PromptForMove)
	ui.Read()
	move, _ := strconv.Atoi(ui.GetLastRead())
	if move < 0 || move > len(board) {
		fmt.Errorf(`The move you picked: %v,\n Was out of this world! Literally!\n Try again, and GO for a number
					that correlates to the open positions on the board.`, move)
		human.GetMove(ui, board)
	}
	return move
}
