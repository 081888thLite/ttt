package ttt

import (
	"testing"
)

func TestEasy_GetMove(t *testing.T) {
	game := NewGame(Configuration{Players: [2]Player{*EASY.create("e"), *EASY.create("z")}})
	easyComputer := game.Players[0]
	move := easyComputer.GetMove(game.Board, game.Players[1])
	if move != 0 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 0, move)
	}
	game.mark(move)
	secondEasyMove := easyComputer.GetMove(game.Board, game.Players[1])
	if secondEasyMove != 1 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 1, secondEasyMove)
	}
}
