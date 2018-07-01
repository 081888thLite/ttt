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

type Console struct{
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
func (display *Console) PlayerOptions() {
	display.Write("\nSet the game mode by entering one of the following options and pressing 'return':\n")
	display.Write("\nEnter 1 For Human vs Human Mode")
	display.Write("\nEnter 2 For Human vs Computer Mode")
	display.Write("\nEnter 3 For Computer vs Computer Mode")
}
func (display *Console) GetMode() Mode {
	display.UI.Read()
	choice := display.UI.LastRead.Msg
	setting, _ := strconv.Atoi(choice)
	return Mode(setting)
}

func NewConsole() *Console {
	return &Console{UI: Sys{}}
}
