package ttt

import (
	"testing"
)

func TestEasyComputerGetMove(t *testing.T) {
	client := &StubClient{}
	game := DefaultNewGame()
	easyComputer := game.Players[0].Strategy
	move := easyComputer.GetMove(client, game.Board)
	if move != 0 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 0, move)
	}
	game.mark(move)
	secondEasyMove := easyComputer.GetMove(client, game.Board)
	if secondEasyMove != 1 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 1, secondEasyMove)
	}
}
