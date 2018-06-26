package ttt

import (
	"testing"
	"fmt"
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

func (game *Game)setUpWinningBoard(howWon string, area string, winningPlayer Player) {
	switch howWon {
	case "Diagonal":
		switch area {
		case "LtR": game.Board = 
			[BoardSize]Player {winningPlayer, Blank, Blank, Blank, winningPlayer, Blank, Blank, Blank, winningPlayer}
		case "RtL": game.Board = 
			[BoardSize]Player {Blank, Blank, winningPlayer, Blank, winningPlayer, Blank, winningPlayer, Blank, Blank}
		}
	case "Row":
		switch area {
		case "1": game.Board = 
			[BoardSize]Player {winningPlayer, winningPlayer, winningPlayer, Blank, Blank, Blank, Blank, Blank, Blank}
		case "2": game.Board = 
			[BoardSize]Player {Blank, Blank, Blank, winningPlayer, winningPlayer, winningPlayer, Blank, Blank, Blank}
		case "3": game.Board = 
			[BoardSize]Player {Blank, Blank, Blank, Blank, Blank, Blank, winningPlayer, winningPlayer, winningPlayer}
		}
	case "Column":
		switch area {
		case "1": game.Board = 
			[BoardSize]Player{winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank}
		case "2": game.Board = 
			[BoardSize]Player{Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank}
		case "3": game.Board = 
			[BoardSize]Player{Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer}
		}
	}
}

func WinTest(t *testing.T, howWon string, area string, winningPlayer Player) {
	game := NewGame()
	game.setUpWinningBoard(howWon, area, winningPlayer)
	actualWinner := game.winner()
	if actualWinner != winningPlayer {
		t.Errorf("Expected %v %v win by %v, to return %s,\n got: %v",
				 howWon, area, winningPlayer, winningPlayer, actualWinner)
	} else {
		fmt.Printf("Passed win detection in %v %v with %v Player\n", howWon, area, winningPlayer)
	}
}

func TestForDiagonalWins(t *testing.T) {
	WinTest(t, "Diagonal", "LtR", Player1)
	WinTest(t, "Diagonal", "LtR", Player2)
	WinTest(t, "Diagonal", "RtL", Player1)
	WinTest(t, "Diagonal", "RtL", Player2)
}

func TestForRowWins(t *testing.T) {
	WinTest(t, "Row", "1", Player1)
	WinTest(t, "Row", "1", Player2)
	WinTest(t, "Row", "2", Player1)
	WinTest(t, "Row", "2", Player2)
	WinTest(t, "Row", "3", Player1)
	WinTest(t, "Row", "3", Player2)
}

func TestForColumnWins(t *testing.T) {
	WinTest(t, "Column", "1", Player1)
	WinTest(t, "Column", "1", Player2)
	WinTest(t, "Column", "2", Player1)
	WinTest(t, "Column", "2", Player2)
	WinTest(t, "Column", "3", Player1)
	WinTest(t, "Column", "3", Player2)
}

func TestWinner(t *testing.T) {
	TestForDiagonalWins(t)
	TestForRowWins(t)
	TestForColumnWins(t)
}