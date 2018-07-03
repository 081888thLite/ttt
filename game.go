package ttt

const (
	Blank = " "
	NoOne = Blank
)

var WinConditions = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {6, 4, 2}}

//Types related to the Board
type BoardSize int
type Board []Piece

type Client interface {
	Write(string)
	Read()
	GetLastSent() string
	GetLastRead() string
}

type Game struct {
	Board         Board
	CurrentPlayer Player
	Waiting       Player
	Winner        Piece
	Players       [2]Player
	Console       *Console
	mv int
}

func NewGame(c Configuration) *Game {
	b := NewBoard(9)
	p := c.Players
	return &Game{
		Board:         b,
		Players:       p,
		CurrentPlayer: p[0],
		Winner:        NoOne,
		Console:       NewConsole(),
	}
}

func NewBoard(boardSize BoardSize) Board {
	var b Board
	for i := 0; i < int(boardSize); i++ {
		b = append(b, Blank)
	}
	return b
}

func (g *Game) setPlayers(player1 Player, player2 Player) *Game {
	g.Players = [2]Player{player1, player2}
	return g
}

func (g *Game) switchPlayers() *Game {
	if g.CurrentPlayer == g.Players[0] {
		g.CurrentPlayer = g.Players[1]
		g.Waiting = g.Players[0]
	} else {
		g.CurrentPlayer = g.Players[0]
		g.Waiting = g.Players[1]
	}
	return g
}

func (board Board) Mark(position int, piece Piece) Board {
	if board[position] == Blank {
		board[position] = piece
	}
	return board
}

func (g *Game) mark(i int) *Game {
	g.Board.Mark(i, g.CurrentPlayer.GetPiece())
	return g
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

func (g *Game) CheckForWin() {
	g.Winner = g.Board.wonBy()
}

func (g *Game) boardFull() bool {
	for _, a := range g.Board {
		if a == Blank {
			return false
		}
	}
	return true
}
func (g *Game) Play() {
	g.Console.greeting()
	for !g.over() {
		g.turn()
	}
	g.end()
}

func (g *Game) end() {
	switch {
	case g.Winner != NoOne:
		g.Console.Write("Game Won By:\n")
		g.Console.Write(string(g.Winner))
		g.Console.Write("\n")
		g.Console.Board(g.Board)
		break
	case g.boardFull():
		g.Console.Write("Game Ends in Draw!\n")
		g.Console.Write("\n")
		break
	}
}

func (g *Game) turn() *Game {
	g.Console.Board(g.Board)
	g.move()
	g.mark(g.mv)
	g.CheckForWin()
	g.switchPlayers()
	return g
}

func (g *Game) move() {
	g.mv = g.CurrentPlayer.GetMove(g.Board, g.Waiting)
}

func (g *Game) over() bool {
	return g.Winner != NoOne || g.boardFull()
}
