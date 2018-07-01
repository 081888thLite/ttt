package ttt

import "testing"

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
			name: "Human strategy for get move, is dependent upon the human. So we get input from their ui.",
			human: Human{Piece: "X", Client: client},
			args: args{client, NewBoard(9)},
		},
	}
		for _, tt := range tests{
		t.Run(tt.name, func (t *testing.T){
		human := tt.human
		human.GetMove(tt.args.board)
	})
	}
}