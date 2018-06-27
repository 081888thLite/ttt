package game

import (
	"strconv"
)

const PromptForMove = "To make a move, enter a number corresponding to an open board position and press 'return':"

//TODO: var Output = bufio.NewWriter(default to os.Stdout, but make this configurable for testing to capture output)

type Human struct{}

func (human Human) getMove(ui Client, board Board) int {
	ui.Send(PromptForMove)
	entered := ui.Receive()
	move, _ := strconv.Atoi(entered)
	//validate move entered is g2g
	return move
}
