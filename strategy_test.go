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
