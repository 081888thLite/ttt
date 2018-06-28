package game

import "testing"

func TestTurn(t *testing.T) {
	t.Log("Game.turn represents a single move cycle.")
	game := DefaultNewGame()
	firstPlayer := game.Players[0]
	secondPlayer := game.Players[1]
	game.turn()
	if game.CurrentPlayer != secondPlayer {
		t.Errorf("Expected current player of %v after the first turn,\n got: %v", secondPlayer, game.CurrentPlayer)
	}
	game.turn()
	if game.CurrentPlayer != firstPlayer {
		t.Errorf("Expected current player of %v after the second turn,\n got: %v", firstPlayer, game.CurrentPlayer)
	}
}
