package ttt

import (
	"testing"
)

func TestEasyComputerGetMove(t *testing.T) {
	game := NewGame(Configuration{Players: [2]Player{EASY.setPlayer("e"), EASY.setPlayer("z")}})
	easyComputer := game.Players[0]
	move := easyComputer.GetMove(game.Board)
	if move != 0 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 0, move)
	}
	game.mark(move)
	secondEasyMove := easyComputer.GetMove(game.Board)
	if secondEasyMove != 1 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 1, secondEasyMove)
	}
}

func TestHuman_GetMove(t *testing.T) {
	client := &StubClient{LastRead: MsgStatus{"5", nil}}
	type args struct {
		ui    Client
		board Board
	}
	tests := []struct {
		name  string
		human Human
		args  args
	}{
		{
			name:  "Human strategy for get move, is dependent upon the human. So we get input from their ui.",
			human: Human{Piece: "X", Client: client},
			args:  args{client, NewBoard(9)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			human := tt.human
			human.GetMove(tt.args.board)
		})
	}
}

func TestEasy_GetMove(t *testing.T) {
	type fields struct {
		Piece  Piece
		Client Client
	}
	type args struct {
		board Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := Easy{
				Piece:  tt.fields.Piece,
				Client: tt.fields.Client,
			}
			if got := comp.GetMove(tt.args.board); got != tt.want {
				t.Errorf("Easy.GetMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedium_GetMove(t *testing.T) {
	type fields struct {
		Piece  Piece
		Client Client
	}
	type args struct {
		board Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := Medium{
				Piece:  tt.fields.Piece,
				Client: tt.fields.Client,
			}
			if got := comp.GetMove(tt.args.board); got != tt.want {
				t.Errorf("Medium.GetMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHard_GetMove(t *testing.T) {
	tests := []struct {
		name   string
		want   int
	}{
		{
			name: "hard computer always picks obvious wins",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := Hard{
				Piece:  Player1,
				Client: &Sys{},
			}
			board := PlacePieces( NewBoard(9), "X", 1, 4 )
			PlacePieces( board, "O", 0, 2, 3 )
			if got := comp.GetMove(board); got != tt.want {
				t.Errorf("Hard.GetMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
func PlacePieces(board Board, piece Piece, i ...int) Board {
	for i, _ := range i {
		board[i] = piece
	}
	return board
}
