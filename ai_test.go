package ttt

import (
	"testing"
	"fmt"
)

func Test_score(t *testing.T) {
	game := DefaultNewGame()
	game.setUpWinningBoard("Row", "1", Player1)
	mm := Minimax{}
	mm.SetMaxMin(Player1, Player2)
	scoreWinMaxWins := mm.Score(game.Board, game.Players[0].GetPiece(),game.Players[1].GetPiece())
	if scoreWinMaxWins != 10 {
		t.Errorf("Expected 10 when winner is ownPiece,\n got: %v", scoreWinMaxWins)
	}
	mm.SetMaxMin(Player1, Player2)
	game.setUpWinningBoard("Col", "1", game.Players[0].GetPiece())
	scoreWinMinWins := mm.Score(game.Board, game.Players[1].GetPiece(),game.Players[0].GetPiece())
	if scoreWinMinWins != -10 {
		t.Errorf("Expected -10 for minimizer,\n got: %v", scoreWinMinWins)
	}

	mm.SetMaxMin(Player2, Player1)
	scoreDrawnBoard := mm.Score(DrawnBoard(), game.Players[1].GetPiece(),game.Players[0].GetPiece())
	if scoreDrawnBoard != 0 {
		t.Errorf("Expected 0 for draw,\n got: %v", scoreDrawnBoard)
	}
}

func Test_minimax(t *testing.T) {
	p1 := *HARD.create(Player1)
	p2 := *HARD.create(Player2)
	players := [2]Player{p1,p2}
	b := NewBoard(9)
	b.PlacePieces(Player1, 2,3,5)
	b.PlacePieces(Player2, 0,7,8)
	fmt.Printf("%v",b)
	mm := new (Minimax)
	choice := mm.minimax(b,players)
	if choice != 4 {
		t.Errorf("Minimax ai chose %v,\nwhen it was expected to choose %v\n", choice,4)
	}
}
