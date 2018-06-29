package ttt

import (
	"strconv"
)

var CompletedTurn = Turn{true, true, true, true}
var NewTurn = Turn{false, false, false, false}

type Turn struct {
	DisplayedBoard bool
	PromptedForMove bool
	ReceivedMove bool
	ValidMove bool
}

func (turn *Turn) displayBoard(client Client, view View, board Board) {
	client.Write(view.ofBoard(board))
	turn.DisplayedBoard = true
}
func (turn *Turn) promptForMove(client Client, view View, player Player) {
	client.Write(view.ofPrompt(player))
	turn.PromptedForMove = true
}

func (turn *Turn) receiveMove(client Client, player Player, board Board) {
	client.Read()
	player.Strategy.GetMove(client, board)
	turn.ReceivedMove = true
}

func (turn *Turn) validateMove(client Client, board Board) {
	move, _ := strconv.Atoi(client.GetLastRead())
	if move < 0 || move > len(board) || board[move] != Blank {
		client.Write(MoveError)
	} else {
		turn.ValidMove = true
	}
}

func (turn Turn) Complete() bool {
	return turn == CompletedTurn
}

