package ttt

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if len(game.Board) != 9 {
		t.Errorf("Expected board size 9,\n got: %v", len(game.Board))
		t.Errorf("Check the constants \n BoardSize: %v or CellsPerSide: %v\n", BoardSize, CellsPerSide)
	}

	for _, field := range game.Board {
		if field != " " {
			t.Errorf("Expected fields with Blank (\" \") values,\n got: %v", field)
		}
	}
}

func TestSwitchPlayers(t *testing.T) {
	game := NewGame()
	if game.CurrentPlayer != Player1 {
		t.Errorf("Expected Player1 (X),\n got: %v", game.CurrentPlayer)
	}

	game.switchPlayers()
	if game.CurrentPlayer != Player2 {
		t.Errorf("Expected Player2 (O),\n got: %v", game.CurrentPlayer)
	}
}

func TestMark(t *testing.T) {
	game := NewGame()
	game.mark(4)
	if game.Board[4] != Player1 {
		t.Errorf("Expected Blank value in center cell to become Player1 (X) value,\n got: %v", game.Board[4])
		t.Errorf("Perhaps the wrong cell was marked...\n here are the rest of the cell values: %v\n", game.Board)
	}

	game.switchPlayers()
	game.mark(5)
	if game.Board[5] != Player2 {
		t.Errorf("Expected Blank value in center cell to become Player2 (O) value,\n got: %v", game.Board[5])
	}
}

func TestWinner(t *testing.T) {
	TestForDiagonalWins(t)
	TestForRowWins(t)
	TestForColumnWins(t)
}

func TestBoardFull(t *testing.T) {
	game := NewGame()
	whenEmptyResult := game.boardFull()
	if  whenEmptyResult != false { t.Errorf("Expected empty board to return false,\n got: %v", whenEmptyResult) }

	game.Board = [BoardSize]Player { Player1, Player2, Player2, Player2, Blank, Player1, Player1, Player1, Player2 }
	whenInPlayResult := game.boardFull()
	if  whenInPlayResult != false { t.Errorf("Expected in play board to return false,\n got: %v", whenInPlayResult) }

	game.Board = [BoardSize]Player { Player1, Player2, Player2, Player2, Player1, Player1, Player1, Player1, Player2 }
	whenFullResult := game.boardFull()
	if  whenFullResult != true { t.Errorf("Expected full board to return true,\n got: %v", whenFullResult) }
}
