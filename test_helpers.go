package ttt

import (
	"testing"
)

const X, O = "X", "O"

func DefaultNewGame() *Game {
	players := [2]Player{&Human{Piece: "X", Client: &Sys{}}, &Human{Piece: "O", Client: &Sys{}}}
	c := Configuration{Players: players}
	return NewGame(c)
}

func FullBoard() Board {
	b := NewBoard(9)
	for i, _ := range b {
		if i%2 == 0 {
			b[i] = "X"
		} else {
			b[i] = "O"
		}
	}
	return b
}

func TestForDiagonalWins(t *testing.T) {
	WinTest(t, "Diagonal", "LtR", X)
	WinTest(t, "Diagonal", "LtR", O)
	WinTest(t, "Diagonal", "RtL", X)
	WinTest(t, "Diagonal", "RtL", O)
}

func TestForRowWins(t *testing.T) {
	WinTest(t, "Row", "1", X)
	WinTest(t, "Row", "1", O)
	WinTest(t, "Row", "2", X)
	WinTest(t, "Row", "2", O)
	WinTest(t, "Row", "3", X)
	WinTest(t, "Row", "3", O)
}

func TestForColumnWins(t *testing.T) {
	WinTest(t, "Column", "1", X)
	WinTest(t, "Column", "1", O)
	WinTest(t, "Column", "2", X)
	WinTest(t, "Column", "2", O)
	WinTest(t, "Column", "3", X)
	WinTest(t, "Column", "3", O)
}

func WinTest(t *testing.T, howWon string, area string, winningPlayer Piece) {
	game := DefaultNewGame()
	game.setUpWinningBoard(howWon, area, winningPlayer)
	game.CheckForWin()
	w := game.Winner
	if w != winningPlayer {
		t.Errorf("Expected %v %v win by %v, to return %s,\n got: %v", howWon, area, winningPlayer, winningPlayer, w)
	} else {
		t.Logf("Passed win detection in %v %v with %v Piece\n", howWon, area, winningPlayer)
	}
}

func (game *Game) setUpWinningBoard(howWon string, area string, winningPlayer Piece) *Game {
	switch howWon {
	case "Diagonal":
		switch area {
		case "LtR":
			game.Board =
				[]Piece{winningPlayer, Blank, Blank, Blank, winningPlayer, Blank, Blank, Blank, winningPlayer}
		case "RtL":
			game.Board =
				[]Piece{Blank, Blank, winningPlayer, Blank, winningPlayer, Blank, winningPlayer, Blank, Blank}
		}
	case "Row":
		switch area {
		case "1":
			game.Board =
				[]Piece{winningPlayer, winningPlayer, winningPlayer, Blank, Blank, Blank, Blank, Blank, Blank}
		case "2":
			game.Board =
				[]Piece{Blank, Blank, Blank, winningPlayer, winningPlayer, winningPlayer, Blank, Blank, Blank}
		case "3":
			game.Board =
				[]Piece{Blank, Blank, Blank, Blank, Blank, Blank, winningPlayer, winningPlayer, winningPlayer}
		}
	case "Column":
		switch area {
		case "1":
			game.Board =
				[]Piece{winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank}
		case "2":
			game.Board =
				[]Piece{Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank}
		case "3":
			game.Board =
				[]Piece{Blank, Blank, winningPlayer, Blank, Blank, winningPlayer, Blank, Blank, winningPlayer}
		}
	}
	return game
}
func LogIncrementingCheckmateTest(t *testing.T, i int) {
	t.Log(
		"Testing Hard takes checkmate move for incrementing win condition:",
		WinConditions[i][0], WinConditions[i][1], WinConditions[i][2],
	)
}
func LogCenterCellCheckmateTest(t *testing.T, i int) {
	t.Logf(
		"Testing Hard takes checkmate move for center cell: %v, when %v %v filled.",
		WinConditions[i][1], WinConditions[i][2], WinConditions[i][0],
	)
}
func LogDecrementingCheckmateTest(t *testing.T, i int) {
	t.Log(
		"Testing Hard takes checkmate move for decrementing win conditions:",
		WinConditions[i][2], WinConditions[i][1], WinConditions[i][0],
	)
}
func (b *Board) PlacePieces(piece Piece, moves ...int) {
	for i := 0; i < len(*b); i++ {
		for _, m := range moves {
			if i == m {
				b.Mark(m, piece)
			}
		}
	}
}

func DrawnBoard() Board {
	return Board{
		X, X, O,
		O, O, X,
		X, X, O,
	}
}
