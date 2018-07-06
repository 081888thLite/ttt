package ttt

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	if len(DefaultNewGame().Board) != 9 {
		t.Errorf("Expected default board size to be: 9,\n got: %v", len(DefaultNewGame().Board))
	}
	g := NewGame(Configuration{})
	b := NewBoard(25)
	size := len(b)
	if size != 25 {
		t.Errorf("Expected default board size to be: 25,\n got: %v", size)
	}
	for _, field := range g.Board {
		if field != " " {
			t.Errorf("Expected fields with Blank (\" \") values,\n got: %v", field)
		}
	}
}

func TestGame_switchPlayers(t *testing.T) {
	game := DefaultNewGame()
	if game.CurrentPlayer != game.Players[0] {
		t.Errorf("Expected *X (X),\n got: %v", game.CurrentPlayer)
	}

	game.switchPlayers()
	if game.CurrentPlayer != game.Players[1] {
		t.Errorf("Expected *O (O),\n got: %v", game.CurrentPlayer)
	}
}

func TestGame_mark(t *testing.T) {
	game := DefaultNewGame()
	game.mark(4)
	if game.Board[4] != game.Players[0].GetPiece() {
		t.Errorf("Expected Blank value in center cell to become *X (X) value,\n got: %v", game.Board[4])
		t.Errorf("Perhaps the wrong cell was marked...\n here are the rest of the cell values: %v\n", game.Board)
	}

	game.switchPlayers()
	game.mark(5)
	if game.Board[5] != game.Players[1].GetPiece() {
		t.Errorf("Expected Blank value in cell 5 (row~2 col~3) to become *O (O) value,\n got: %v", game.Board[5])
	}

	game.switchPlayers()
	game.mark(5)
	if game.Board[5] != game.Players[1].GetPiece() {
		message := "Expected already marked *O (O) value in cell 5 (row~2 col~3) to not change,\n got: %v"
		t.Errorf(message, game.Board[5])
	}
}

func TestWinner(t *testing.T) {
	TestForDiagonalWins(t)
	TestForRowWins(t)
	TestForColumnWins(t)
}

func TestBoardFull(t *testing.T) {
	game := DefaultNewGame()
	whenEmptyResult := game.boardFull()
	if whenEmptyResult != false {
		t.Errorf("Expected empty board to return false,\n got: %v", whenEmptyResult)
	}

	game.Board = []Piece{X, O, O, O, Blank, X, X, X, O}
	whenInPlayResult := game.boardFull()
	if whenInPlayResult != false {
		t.Errorf("Expected in game board to return false,\n got: %v", whenInPlayResult)
	}

	game.Board = []Piece{X, O, O, O, X, X, X, X, O}
	whenFullResult := game.boardFull()
	if whenFullResult != true {
		t.Errorf("Expected full board to return true,\n got: %v", whenFullResult)
	}
}

func TestGame_setPlayers(t *testing.T) {
	game := DefaultNewGame()
	game = game.setPlayers(&Easy{"S", &StubClient{}}, &Easy{"P", &StubClient{}})
	player1piece := game.Players[0].GetPiece()
	player2piece := game.Players[1].GetPiece()
	if player1piece != "S" && player2piece != "P" {
		t.Errorf("Expected player pieces to be set to:\n'S' & 'P'\n got: \n%v & %v ", player1piece, player2piece)

	}
	game = DefaultNewGame()
	game.setPlayers(&Easy{"S", &StubClient{}}, &Easy{"P", &StubClient{}})
	hard1piece := game.Players[0].GetPiece()
	hard2piece := game.Players[1].GetPiece()
	if hard1piece != "S" && hard2piece != "P" {
		t.Error("Expected Hard vs. Hard to work, and GetPiece to be callable")
	}
}

func TestGame_Play(t *testing.T) {
	e1 := &Easy{}
	e2 := &Easy{}
	h1 := &Hard{}
	h2 := &Hard{}
	ePlayers := [2]Player{e1, e2}
	hPlayers := [2]Player{h1, h2}
	type fields struct {
		Board         Board
		CurrentPlayer Player
		Winner        Piece
		Players       [2]Player
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Play runs ok w/ 2 easy players",
			fields: fields{
				Board:         FullBoard(),
				CurrentPlayer: e1,
				Players:       ePlayers,
				Winner:        NoOne,
			},
		}, {
			name: "Play runs ok w/ 2 hard players",
			fields: fields{
				Board:         FullBoard(),
				CurrentPlayer: h1,
				Players:       hPlayers,
				Winner:        NoOne,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{
				Board:         tt.fields.Board,
				CurrentPlayer: tt.fields.CurrentPlayer,
				Winner:        tt.fields.Winner,
				Players:       tt.fields.Players,
				Display:       &Console{},
			}
			game.Play()
		})
	}
}

func TestGame_turn(t *testing.T) {
	type args struct {
		givenBoard    Board
		currentPlayer Player
		shouldTake    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Easy computer", args: args{NewBoard(9), &Easy{piece: X, Client: &StubClient{}}, 0}},
		{name: "Easy computer",
			args: args{Board{X, Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank}, &Easy{piece: O}, 1}},
		{name: "Easy computer",
			args: args{Board{X, O, Blank, Blank, Blank, Blank, Blank, Blank, Blank}, &Easy{piece: X}, 2}},
		{name: "Easy computer",
			args: args{Board{X, O, Blank, X, Blank, Blank, Blank, Blank, Blank}, &Easy{piece: O}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := make([]Piece, len(tt.args.givenBoard))
			copy(given, tt.args.givenBoard)
			game := &Game{Board: tt.args.givenBoard, CurrentPlayer: tt.args.currentPlayer, Display: &Console{}}
			game.turn()
			if game.Board[tt.args.shouldTake] != tt.args.currentPlayer.GetPiece() {
				e := "\nGiven board:\n %v\n%v player %v\nShould take position:\n%v\nbut board looks like:\n%v\n"
				t.Fail()
				t.Errorf(e, given, tt.name, tt.args.currentPlayer, tt.args.shouldTake, game.Board)
			}
		})
	}
}
