package ttt

import (
	"fmt"
	"testing"
)

func Test_score(t *testing.T) {
	b := Board{X, Blank, Blank, O, X, Blank, O, Blank, X}
	mm := Minimax{depth: 0}
	xwin := mm.Score(b.wonBy(), X, 0)
	if xwin != 10 {
		t.Errorf("Expected 10 when X wins and depth is 0,\n got: %v", xwin)
	}
	b = []Piece{O, Blank, Blank, X, O, Blank, X, Blank, O}
	owin := mm.Score(b.wonBy(), X, 0)
	if owin != -10 {
		t.Errorf("Expected -10 when O wins and depth is 0,\n got: %v", owin)
	}
	scoreDrawnBoard := mm.Score(NoOne, X, 0)
	if scoreDrawnBoard != 0 {
		t.Errorf("Expected 0 for draw,\n got: %v", scoreDrawnBoard)
	}
}

func Test_minimax(t *testing.T) {
	b := NewBoard(9)
	b.PlacePieces(X, 2, 3, 5)
	b.PlacePieces(O, 0, 7, 8)
	fmt.Printf("%v", b)
	mm := new(Minimax)
	choice := mm.minimax(b, X, 0)
	if choice != 1 {
		t.Errorf("Minimax ai chose %v,\nwhen it was expected to choose %v\n\npiece for original max was:%v", choice, 1, X)
	}
}
