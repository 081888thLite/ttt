package ttt

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame()
	if len(game.board) != 9 {
		t.Errorf("Expected board size 9, got: %v", len(game.board))
	}

	for _, field := range game.board {
		if field != Blank {
			t.Errorf("Expected field with value %v, got: %v", Blank, field)
		}
	}
}
