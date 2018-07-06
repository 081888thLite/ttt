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

func TestHard_ReplaceBoard(t *testing.T) {
	b := NewBoard(9)
	b.PlacePieces("F", 1, 3, 5)
	b.PlacePieces("L", 2, 6, 7)
	h := Hard{piece: "F"}
	nb := h.ReplaceBoard(b)
	if nb[1] != X || nb[3] != X || nb[5] != X {
		t.Errorf("Expected hard player's pieces to be replaced with X")
	}
	if nb[2] != O || nb[6] != O || nb[7] != O {
		t.Errorf("Expected opp of hard player's pieces to be replaced with O")
	}
	if nb[8] != NoOne || nb[4] != NoOne || nb[0] != NoOne {
		t.Errorf("Expected blank spots to remain blank but got: %v, %v, and %v", nb[8], nb[4], nb[0])
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
