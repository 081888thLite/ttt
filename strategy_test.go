package ttt

import (
	"testing"
)

func TestEasy_GetMove(t *testing.T) {
	game := NewGame(Configuration{Players: [2]Player{*EASY.create("e"), *EASY.create("z")}})
	easyComputer := game.Players[0]
	move := easyComputer.GetMove(game.Board, game.Players[1])
	if move != 0 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 0, move)
	}
	game.mark(move)
	secondEasyMove := easyComputer.GetMove(game.Board, game.Players[1])
	if secondEasyMove != 1 {
		t.Errorf("Expected easy computer to choose first available position: %v,\n got: %v", 1, secondEasyMove)
	}
}

//func TestMedium_GetMove(t *testing.T) {
//	game := NewGame(Configuration{Players: [2]Player{MEDIUM.create("m"), MEDIUM.create("d")}})
//	m := MEDIUM.create("m")
//	for i := 0; i < 7 ; i++ {
//		t.Logf(
//			"Testing Medium takes checkmate move of %v on the condition that spots %v %v are filled.",
//			WinConditions[i][2], WinConditions[i][0], WinConditions[i][1],
//		)
//		PlacePieces(game.Board, "m", WinConditions[i][0], WinConditions[i][1])
//		expectedMove := WinConditions[i][2]
//		if got := m.GetMove(game.Board); got != expectedMove {
//			blnk := game.Board.blanks()
//			t.Logf(":::::::::\n %v", blnk)
//			t.Errorf(
//				"Expected Medium computer to choose a quick win like %v when available,\n got: %v",
//				expectedMove, got)
//		}
//	}
//}

func TestHard_GetMove(t *testing.T) {
	for i := 0; i < 7; i++ {
		LogIncrementingCheckmateTest(t, i)
		game := NewGame(Configuration{Players: [2]Player{&Hard{piece: X}, &Hard{piece: O}}})
		var x = game.Players[0]
		var o = game.Players[1]
		game.Board.PlacePieces(X, WinConditions[i][0], WinConditions[i][1])
		move := x.GetMove(game.Board, o)
		expectedMove := WinConditions[i][2]
		if move != expectedMove {
			t.Errorf(
				"Expected hard computer to choose a quick win like %v when available,\n got: %v",
				expectedMove, move,
			)
		}
		LogDecrementingCheckmateTest(t, i)
		ClearBoard(game)
		game.Board.PlacePieces(X, WinConditions[i][2], WinConditions[i][1])
		game.Board.PlacePieces(O, 5, 7)
		move = x.GetMove(game.Board, o)
		expectedMove = WinConditions[i][0]
		if move != expectedMove {
			t.Errorf(
				"Expected hard computer to choose a quick win like %v when available,\n got: %v",
				expectedMove, move,
			)
		}
		LogCenterCellCheckmateTest(t, i)
		ClearBoard(game)
		game.Board.PlacePieces(X, WinConditions[i][0], WinConditions[i][2])
		move = x.GetMove(game.Board, o)
		expectedMove = WinConditions[i][1]
		if move != expectedMove {
			t.Errorf(
				"Expected hard computer to choose a quick win like %v when available,\n got: %v",
				expectedMove, move,
			)
		}
	}
}

func ClearBoard(game *Game) {
	game.Board = NewBoard(9)
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
		opp   Hard
		args  args
	}{
		{
			name:  "Human strategy for get move, is dependent upon the human. So we get input from their ui.",
			human: Human{piece: X, Client: client},
			opp:   Hard{piece: O},
			args:  args{client, NewBoard(9)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.human.GetMove(tt.args.board, &tt.opp)
		})
	}
}

func Test_PlacePieces(t *testing.T) {
	actualBoard := NewBoard(9)
	actualBoard.PlacePieces("p", 1, 2, 3)
	expectedBoard := Board{Blank, "p", "p", "p", Blank, Blank, Blank, Blank, Blank}
	for i := 0; i < 9; i++ {
		if actualBoard[i] != expectedBoard[i] {
			t.Errorf("Expected:\n%v,\nGot:\n%v", expectedBoard, actualBoard)
		}
		t.Log("Passed Place Pieces Test")
	}
}
