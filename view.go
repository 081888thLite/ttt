package ttt

import (
	"math"
	"strconv"
	"strings"
)

const MoveError = `The move you picked: %v,\nWas out of this world! Literally!\n Try again, and GO for a number
that correlates to the open positions on the board.`

type View interface {
	ofBoard(Board) string
	ofPrompt(Player) string
	ofMove(string) string
	ofWinner(Game) string
	ofDraw() string
	ofPlayerThinking(Player) string
}

type Console struct {
	UI Sys
}

func (display *Console) ofBoard(b Board) string {
	rowSize := math.Sqrt((float64(len(b))))
	var viewableBoard strings.Builder
	for i, e := range b {
		if i == 0 || math.Remainder(float64(i), rowSize) == 0 {
			viewableBoard.WriteRune('\n')
		}
		if e == Blank {
			index := strconv.Itoa(i)
			viewableBoard.WriteString(index)
			viewableBoard.WriteString(" ")
		} else {
			viewableBoard.WriteString(string(e))
			viewableBoard.WriteString(" ")
		}
	}
	viewableBoard.WriteString("\n")
	return viewableBoard.String()
}

func (display *Console) Board(board Board) {
	printableBoard := display.ofBoard(board)
	display.Write(printableBoard)
}

func (display *Console) ofPrompt(Player) string         { return "" }
func (display *Console) ofMove(string) string           { return "" }
func (display *Console) ofWinner(Game) string           { return "" }
func (display *Console) ofDraw() string                 { return "" }
func (display *Console) ofPlayerThinking(Player) string { return "" }
func (display *Console) greeting() {
	display.Write("\nWELCOME TO TICTACTOE\nwrote in Go")
}

func (display *Console) Write(msg string) {
	display.UI.Write(msg)
}
func (display *Console) GameMenu() {
	display.Write("\n**Game Menu**\n")
}

func (display *Console) PlayerMenu(i int) (Strategy, Piece) {
	display.Write("\nPlayer Menu\n")
	display.Write("______________\n")
	display.Write("NOTE: Unless you want this menu to keep resetting,\n")
	display.Write("enter a single character and press return for all options\n\n")
	strategy := display.PickStrategy(i)
	piece := display.PickPiece(i)
	return strategy, piece
}

func (display *Console) PickPiece(order int) Piece {
	display.Write("\nWhat do you want Player ")
	display.Write(strconv.Itoa(order + 1))
	display.Write(" to mark the board with?\n")
	display.Write(":")
	display.UI.Read()
	choice := display.UI.GetLastRead()
	return Piece(choice)
}

func (display *Console) PickStrategy(order int) Strategy {
	display.Write("\nWhat kind of Mover is Player ")
	display.Write(strconv.Itoa(order + 1))
	display.Write("?\nEnter...")
	display.Write("\n1 for HUMAN\n")
	display.Write("2 for EASY Difficulty AI\n")
	display.Write("3 for MEDIUM Difficulty AI\n")
	display.Write("4 for HARD Difficulty AI\n")
	display.Write(":")
	display.UI.Read()
	choice, err := strconv.Atoi(display.UI.GetLastRead())
	if choice < 1 || choice > 4 || err != nil {
		display.Write("***YOUR ENTERED SOMETHING WRONG***\n***CHECK INPUT AND TRY AGAIN***")
		return display.PickStrategy(order)
	}
	return Strategy(choice)
}

func (display *Console) WantsSetup() bool {
	display.Write("\n**Want to configure the players?**\n")
	display.Write("DEFAULT PLAYERS: Human w/ Piece of X vs. HardComputer w/ Piece of O")
	display.Write("\nEnter y and press 'return' to setup the players")
	display.Write("\nEnter n and press 'return' to skip and use defaults\n")
	display.Write(":")
	display.UI.Read()
	choice := display.UI.GetLastRead()
	return strings.ToLower(choice) == "y"
}

func NewConsole() *Console {
	return &Console{UI: Sys{}}
}
func (display *Console) getHumanMove() int {
	display.Write("To make a move, enter a number corresponding\n")
	display.Write("to an open board position and press 'return':\n")
	display.UI.Read()
	choice, _ := strconv.Atoi(display.UI.GetLastRead())
	return choice
}
