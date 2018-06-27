package game

import (
	"strings"
	"math"
	"strconv"
)

type View interface {
	ofBoard(Board) string
	ofPrompt(string) string
	ofMove(string) string
	ofWinner(Game) string
	ofDraw(Game) string
	ofPlayerThinking(Player) string
}

type ConsoleView struct {}

func (view ConsoleView) ofBoard(b Board) string {
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
	return viewableBoard.String()
}

func (view ConsoleView) ofPrompt(string) string { return "" }
func (view ConsoleView) ofMove(string) string { return "" }
func (view ConsoleView) ofWinner(Game) string { return "" }
func (view ConsoleView) ofDraw(Game) string { return "" }
func (view ConsoleView) ofPlayerThinking(Player) string { return "" }