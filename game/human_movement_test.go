package game

import (
	"testing"
)

func TestHumanGetMove(t *testing.T) {
	game := NewGame(9, Player{"X", Human{}}, Player{"O", Human{}} )
	human := game.Players[0].Strategy
	//os.Stdin <- "5"
	entered := human.getMove(game.Board)
	if entered != 5 {
		t.Errorf("Expected %v", 5)
	}
}
