package game

import "testing"

func TestTurn(t *testing.T) {
	t.Log("Game.turn represents a single move cycle.")
	game := DefaultNewGame()
	secondPlayer := game.Players[1]
	game.turn()
	if game.CurrentPlayer != secondPlayer {
		t.Errorf("Expected current player of %v after the first turn,\n got: %v", secondPlayer, game.CurrentPlayer)
	}
}
