package ttt

import (
	"testing"
	"fmt"
)

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

func WinTest(t *testing.T, howWon string, area string, winningPlayer Player) {
	game := NewGame()
	game.setUpWinningBoard(howWon, area, winningPlayer)
	game.checkForWin()
	w := game.Winner
	if w != winningPlayer {
		t.Errorf("Expected %v %v win by %v, to return %s,\n got: %v", howWon, area, winningPlayer, winningPlayer, w)
	} else {
		fmt.Printf("Passed win detection in %v %v with %v Player\n", howWon, area, winningPlayer)
	}
}

func (game *Game) setUpWinningBoard(howWon string, area string, winningPlayer Player) {
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