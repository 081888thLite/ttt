package game

import (
	"testing"
)

var GameOfHumanPlayers = NewGame(9, Player{"X", Human{}}, Player{"O", Human{}})

func TestHumanGetMove(t *testing.T) {
	client := &StubIO{"5", PromptForMove, ""}
	game := GameOfHumanPlayers
	human := game.Players[0].Strategy
	choice := human.getMove(client, game.Board)
	if choice != 5 {
		t.Errorf("Expected int %v,\n got: %v", 5, choice)
	}
}
