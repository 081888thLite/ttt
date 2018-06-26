package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	defaultNewGame := DefaultNewGame()
	if len(defaultNewGame.Board) != 9 {
		t.Errorf("Expected default board size to be: 9,\n got: %v", len(defaultNewGame.Board))
	}
	game := NewGame(25, Player{"1", &Easy{}}, Player{"2", &Easy{}})
	if len(game.Board) != 25 {
		t.Errorf("Expected default board size to be: 9,\n got: %v", len(game.Board))
	}
	for _, field := range game.Board {
		if field != " " {
			t.Errorf("Expected fields with Blank (\" \") values,\n got: %v", field)
		}
	}
}

func TestSwitchPlayers(t *testing.T) {
	game := DefaultNewGame()
	if game.CurrentPlayer != *Player1 { t.Errorf("Expected *Player1 (X),\n got: %v", game.CurrentPlayer) }

	game.switchPlayers()
	if game.CurrentPlayer != *Player2 { t.Errorf("Expected *Player2 (O),\n got: %v", game.CurrentPlayer) }
}

func TestMark(t *testing.T) {
	game := DefaultNewGame()
	game.mark(4)
	if game.Board[4] != *Player1 {
		t.Errorf("Expected Blank value in center cell to become *Player1 (X) value,\n got: %v", game.Board[4])
		t.Errorf("Perhaps the wrong cell was marked...\n here are the rest of the cell values: %v\n", game.Board)
	}

	game.switchPlayers()
	game.mark(5)
	if game.Board[5] != *Player2 {
		t.Errorf("Expected Blank value in cell 5 (row~2 col~3) to become *Player2 (O) value,\n got: %v", game.Board[5])
	}

	game.switchPlayers()
	game.mark(5)
	if game.Board[5] != *Player2 {
		message := "Expected already marked *Player2 (O) value in cell 5 (row~2 col~3) to not change,\n got: %v"
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
	if  whenEmptyResult != false { t.Errorf("Expected empty board to return false,\n got: %v", whenEmptyResult) }

	game.Board = []Piece{ *Player1, *Player2, *Player2, *Player2, Blank, *Player1, *Player1, *Player1, *Player2 }
	whenInPlayResult := game.boardFull()
	if  whenInPlayResult != false { t.Errorf("Expected in game board to return false,\n got: %v", whenInPlayResult) }

	game.Board = []Piece{ *Player1, *Player2, *Player2, *Player2, *Player1, *Player1, *Player1, *Player1, *Player2 }
	whenFullResult := game.boardFull()
	if  whenFullResult != true { t.Errorf("Expected full board to return true,\n got: %v", whenFullResult) }
}
