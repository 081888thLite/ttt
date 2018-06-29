package ttt

import "testing"

func TestConsoleView_OfBoard(t *testing.T) {
	game := DefaultNewGame()
	view := ConsoleView{}
	t.Log("ConsoleView.ofBoard prints index in place of Blank Cells")
	result := view.ofBoard(game.Board)
	expected := "\n0 1 2 \n3 4 5 \n6 7 8 \n"
	if result != expected {
		t.Errorf("Expected empty board to return: %v ,\n got: %v", expected, result)
	}

	t.Log("ConsoleView.ofBoard prints Player piece for marked Cells")
	game.mark(1)
	game.switchPlayers()
	game.mark(5)
	result = view.ofBoard(game.Board)
	expected = "\n0 X 2 \n3 4 O \n6 7 8 \n"
	if result != expected {
		t.Errorf("Expected marked board to return: %v ,\n got: %v", expected, result)
	}
}
