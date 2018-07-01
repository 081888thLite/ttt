package ttt

import (
	"testing"
)

func TestEasyComputerGetMove(t *testing.T) {
	game := NewGame(9, Easy{}, Easy{})
	easyComputer := game.Players[0]
	move := easyComputer.GetMove(game.Board)
	if move != 0 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 0, move)
	}
	game.mark(move)
	secondEasyMove := easyComputer.GetMove(game.Board)
	if secondEasyMove != 1 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 1, secondEasyMove)
	}
}
