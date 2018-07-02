package ttt

import (
	"testing"
)

func TestEasy_GetMove(t *testing.T) {
	game := NewGame(Configuration{Players: [2]Player{EASY.getPlayer("e"), EASY.getPlayer("z")}})
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

func TestHard_GetMove(t *testing.T) {
	for i := 0; i < 7 ; i++ {
		t.Logf(
				"Testing Hard takes Winning move of %v on the condition that spots %v %v are filled.",
			 	WinConditions[i][2], WinConditions[i][0], WinConditions[i][1],
			 	)
		game := NewGame(Configuration{Players: [2]Player{HARD.getPlayer("h"), HARD.getPlayer("d")}})
		h := game.Players[0]
		PlacePieces(game.Board, "h", WinConditions[i][0], WinConditions[i][1])
		move := h.GetMove(game.Board)
		expectedMove := WinConditions[i][2]
		if move != 0 {
			t.Errorf(
					"Expected hard computer to choose a quick win like %v when available,\n got: %v",
				 	expectedMove, move,
				 	)
		}
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

func PlacePieces(board Board, piece Piece, i ...int) Board {
	for index := range i {
		board[index+1] = piece
	}
	return board
}

func Test_PlacePieces(t *testing.T) {
	startBoard := NewBoard(9)
	actualBoard := PlacePieces( startBoard[:], "p", 1, 2, 3 )
	expectedBoard := Board{Blank, "p", "p", "p", Blank, Blank, Blank, Blank, Blank }
	for i := 0; i < 9; i++ {
		if actualBoard[i] != expectedBoard[i] {
			t.Errorf("Expected:\n%v,\nGot:\n%v", expectedBoard, actualBoard)
		}
		t.Log("Passed Place Pieces Test")
	}
}