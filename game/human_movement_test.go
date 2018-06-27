package game

import (
	"testing"
)

func TestHumanGetMove(t *testing.T) {
	client := &TestIO{"5", PromptForMove, ""}
	game := NewGame(9, Player{"X", Human{}}, Player{"O", Human{}})
	human := game.Players[0].Strategy
	entered := human.getMove(client, game.Board)
	if entered != 5 {
		t.Errorf("Expected %v", 5)
	}
}
