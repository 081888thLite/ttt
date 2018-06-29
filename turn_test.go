package ttt

import "testing"

func TestDisplayBoard(t *testing.T) {
	game := DefaultNewGame()
	client := &StubClient{}
	view := ConsoleView{}
	turn := Turn{}
	turn.displayBoard(client, view, game.Board)
	if turn.DisplayedBoard != true {
		t.Error("Expected DisplayedBoard set to true after the call to displayBoard,\n got:", turn.DisplayedBoard)
	}
}

func TestPromptForMove(t *testing.T) {
	game := DefaultNewGame()
	client := &StubClient{}
	view := ConsoleView{}
	turn := Turn{}
	turn.promptForMove(client, view, game.Players[0])
	if turn.PromptedForMove != true {
		t.Error("Expected PromptedForMove set to true after the call to promptForMove,\n got:", false)
	}
}

func TestReceiveMove(t *testing.T) {
	game := DefaultNewGame()
	client := &StubClient{}
	turn := Turn{}
	turn.receiveMove(client, game.Players[0], NewBoard(0))
	if turn.ReceivedMove != true {
		t.Error("Expected ReceivedMove set to true after the call to receiveMove,\n got:", false)
	}
}

func TestValidateMove(t *testing.T) {
	game := DefaultNewGame()
	game.mark(5)
	client := &StubClient{}
	client.LastRead = MsgStatus{"5", nil}
	turn := Turn{}
	turn.validateMove(client, game.Board)
	if turn.ValidMove != false {
		t.Error("Expected ValidMove to be false after the player attempts to move onto a filled space,\n got:", true)
	}
	client.LastRead = MsgStatus{"6", nil}
	turn.validateMove(client, game.Board)
	if turn.ValidMove != true {
		t.Error("Expected ValidMove to be true after the player attempts to move onto blank space,\n got:", false)
	}
}

func TestComplete(t *testing.T) {
	if CompletedTurn.Complete() != true {
		t.Error(`Expected true in response to Complete() call on turn that has completed displaying 
				\n the board, prompting for a move, receiving a move, and validating a move,
				\n got:`, false)
	}
	if NewTurn.Complete() != false {
		t.Error("Expected false in response to Complete() call on turn that has not started\n got:", true)
	}
}
