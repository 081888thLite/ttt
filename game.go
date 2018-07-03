package ttt

const (
	Blank = " "
	NoOne = Blank
)

var WinConditions = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {6, 4, 2}}

//Types related to the Board
type BoardSize int
type Board []Piece

type Player interface {
	GetMove(b Board, opp Player) int
	GetPiece() Piece
}

type Game struct {
	Board         Board
	CurrentPlayer Player
	Waiting       Player
	Winner        Piece
	Players       [2]Player
	Display       *Console
}

func NewGame(c Configuration) *Game {
	b := NewBoard(9)
	p := c.Players
	return &Game{
		Board:         b,
		Players:       p,
		CurrentPlayer: p[0],
		Winner:        NoOne,
		Display:       NewConsole(),
	}
}

func NewBoard(boardSize BoardSize) Board {
	var b Board
	for i := 0; i < int(boardSize); i++ {
		b = append(b, Blank)
	}
	return b
}

func (game *Game) setPlayers(player1 Player, player2 Player) *Game {
	game.Players = [2]Player{player1, player2}
	return game
}

func (game *Game) switchPlayers() *Game {
	if game.CurrentPlayer == game.Players[0] {
		game.CurrentPlayer = game.Players[1]
		game.Waiting = game.Players[0]
	} else {
		game.CurrentPlayer = game.Players[0]
		game.Waiting = game.Players[1]
	}
	return game
}

func (board Board) Mark(position int, piece Piece) Board {
	if board[position] == Blank {
		board[position] = piece
	}
	return board
}

func (game *Game) mark(position int) *Game {
	game.Board.Mark(position, game.CurrentPlayer.GetPiece())
	return game
}

func (b Board) allCellsMatch(cells ...int) bool {
	for _, cell := range cells {
		if b[cell] != b[cells[0]] {
			return false
		}
	}
	return true
}

func (b Board) blanks() []int {
	blanks := []int{}
	for i, e := range b {
		if e == Blank {
			blanks = append(blanks, i)
		}
	}
	return blanks
}

func (b Board) wonBy() Piece {
	for _, cells := range WinConditions {
		possibleWinner := b[cells[0]]
		if possibleWinner != NoOne && b.allCellsMatch(cells...) {
			return possibleWinner
		}
	}
	return NoOne
}

func (game *Game) CheckForWin() {
	game.Winner = game.Board.wonBy()
}

func (game *Game) boardFull() bool {
	for _, a := range game.Board {
		if a == Blank {
			return false
		}
	}
	return true
}
func (game *Game) Play() {
	game.Display.greeting()
	for !over(game) {
		game.turn()
	}
	switch {
	case game.Winner != NoOne:
		game.Display.Write("Game Won By:\n")
		game.Display.Write(string(game.Winner))
		game.Display.Write("\n")
		game.Display.Board(game.Board)
		break
	case game.boardFull():
		game.Display.Write("Game Ends in Draw!\n")
		game.Display.Write("\n")
		break
	}
}

func (game *Game) turn() *Game {
	game.Display.Board(game.Board)
	game.mark(game.CurrentPlayer.GetMove(game.Board, game.Waiting))
	game.CheckForWin()
	game.switchPlayers()
	return game
}

func over(game *Game) bool {
	return game.Winner != NoOne || game.boardFull()
}
