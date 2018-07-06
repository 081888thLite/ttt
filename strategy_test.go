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

func TestMedium_GetMove(t *testing.T) {
	for i := 0; i < 7; i++ {

		game := NewGame(Configuration{Players: [2]Player{*MEDIUM.create("h"), *MEDIUM.create("d")}})
		var h = game.Players[0]
		game.Board.PlacePieces("h", WinConditions[i][0], WinConditions[i][1])
		placeOppPieces(i, game)
		move := h.GetMove(game.Board, game.Players[1])
		expectedMove := WinConditions[i][2]
		if move != expectedMove {
			t.Errorf(
				"Expected medium computer to choose a quick win like %v when available,\n got: %v",
				expectedMove, move,
			)
		} else {
			t.Log("Passed")
		}

		game.Board = NewBoard(9)
		game.Board.PlacePieces("h", WinConditions[i][2], WinConditions[i][1])
		placeOppPieces(i, game)
		move = h.GetMove(game.Board, game.Players[1])
		expectedMove = WinConditions[i][0]
		if move != expectedMove {
			t.Errorf(
				"Expected medium computer to choose a quick win like %v when available,\n got: %v",
				expectedMove, move,
			)
		} else {
			t.Log("Passed")
		}

		game.Board = NewBoard(9)
		game.Board.PlacePieces("h", WinConditions[i][0], WinConditions[i][2])
		placeOppPieces(i, game)
		move = h.GetMove(game.Board, game.Players[1])
		expectedMove = WinConditions[i][1]
		if move != expectedMove && game.Board[expectedMove] != "d" {
			t.Errorf(
				"Expected medium computer to choose a quick win like %v when available,\n got: %v and board at pos was %v",
				expectedMove, move, game.Board[expectedMove],
			)
		} else {
			t.Log("Passed")
		}
	}
}

func placeOppPieces(i int, game *Game) {
	if anyZeroOrEight(i) {
		game.Board.PlacePieces("d", 3)
	} else {
		if anyZero(i) {
			game.Board.PlacePieces("d", 8)
		} else if anyEight(i) {
			game.Board.PlacePieces("d", 0)
		}
	}
}

func anyZero(i int) bool {
	return WinConditions[i][0] == 0 || WinConditions[i][1] == 0 || WinConditions[i][2] == 0
}

func anyEight(i int) bool {
	return WinConditions[i][0] == 8 || WinConditions[i][1] == 8 || WinConditions[i][2] == 8
}

func anyZeroOrEight(i int) bool {
	return anyEight(i) || anyZero(i)
}
