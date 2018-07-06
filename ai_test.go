package ttt

import (
	"fmt"
	"testing"
)

func Test_score(t *testing.T) {
	b := []Piece{X, Blank, Blank, O, X, Blank, O, Blank, X}
	mm := Minimax{}
	firstCallerWin := mm.Score(b, 0, X)
	if firstCallerWin != 10 {
		t.Errorf("Expected 10 when winner is in turn player and depth 0,\n got: %v", firstCallerWin)
	}
	notCallerWin := mm.Score(b, 2, O)
	if notCallerWin != 8 {
		t.Errorf("Expected -10 when depth is 2 & in turn player did not win,\n got: %v", notCallerWin)
	}
	scoreDrawnBoard := mm.Score(DrawnBoard(), 0, O)
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
	if choice != 4 {
		t.Errorf("Minimax ai chose %v,\nwhen it was expected to choose %v\n\npiece for original max was:%v", choice, 4, X)
	}
	mm = new(Minimax)
	b = NewBoard(9)
	b.PlacePieces(X, 2, 3, 5, 6)
	b.PlacePieces(O, 0, 7, 8)
	choice = mm.minimax(b, O, 0)
	if choice != 4 {
		t.Errorf("Minimax ai chose %v,\nwhen it was expected to choose %v\n\npiece for original max was:%v", choice, 4, O)
	}
}

func TestMinimax_evaluate(t *testing.T) {
	type fields struct {
		level int
		min   Piece
		max   Piece
		tree  []Node
	}
	type args struct {
		nodes []Node
		l     int
		p     Piece
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "Chooses highest score when max wins is in turn",
			fields: fields{level: 0, min: O, max: X, tree: []Node{Node{pos: 2, score: 0}, Node{pos: 3, score: 4}, Node{pos: 4, score: 10}}},
			args:   args{nodes: []Node{Node{pos: 2, score: 0}, Node{pos: 3, score: 4}, Node{pos: 4, score: 10}}, l: 1, p: X},
			want:   4,
		},
		{
			name:   "Chooses highest score when min in turn",
			fields: fields{level: 0, min: O, max: X, tree: []Node{Node{pos: 2, score: -4}, Node{pos: 3, score: 4}, Node{pos: 4, score: -25}}},
			args:   args{nodes: []Node{Node{pos: 2, score: -4}, Node{pos: 3, score: 4}, Node{pos: 4, score: -25}}, l: 2, p: O},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &Minimax{
				level: tt.fields.level,
				min:   tt.fields.min,
				max:   tt.fields.max,
				tree:  tt.fields.tree,
			}
			if got := mm.evaluate(tt.args.nodes, tt.args.l, tt.args.p); got != tt.want {
				t.Errorf("Minimax.evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
