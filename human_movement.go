package ttt

const PromptForMove = "To make a move, enter a number corresponding to an open board position and press 'return':"

type Human struct{}

func (human Human) GetMove(ui Client, board Board) {
	ui.Write(PromptForMove)
	ui.Read()
}
