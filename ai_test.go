package ttt

import (
	"fmt"
	"testing"
)

func Test_score(t *testing.T) {
	firstCaller := Hard{piece: X}
	b := []Piece{X, Blank, Blank, O, X, Blank, O, Blank, X}
	mm := Minimax{}
	mm.SetCaller(firstCaller)
	firstCallerWin := mm.Score(b)
	if firstCallerWin != 10 {
		t.Errorf("Expected 10 when winner is caller,\n got: %v", firstCallerWin)
	}
	secondCaller := Hard{piece: O}
	mm.SetCaller(secondCaller)
	notCallerWin := mm.Score(b)
	if notCallerWin != -10 {
		t.Errorf("Expected -10 when not the caller,\n got: %v", notCallerWin)
	}
	scoreDrawnBoard := mm.Score(DrawnBoard())
	if scoreDrawnBoard != 0 {
		t.Errorf("Expected 0 for draw,\n got: %v", scoreDrawnBoard)
	}
}

func Test_minimax(t *testing.T) {
	p1 := Hard{piece: X}
	p2 := Hard{piece: O}
	players := [2]Player{&p1, &p2}
	b := NewBoard(9)
	b.PlacePieces(X, 2, 3, 5)
	b.PlacePieces(O, 0, 7, 8)
	fmt.Printf("%v", b)
	mm := new(Minimax)
	mm.SetCaller(p1)
	choice := mm.minimax(b, players)
	if choice != 4 {
		t.Errorf("Minimax ai chose %v,\nwhen it was expected to choose %v\n\npiece for original max was:%v", choice, 4, players[0].GetPiece())
	}
}
