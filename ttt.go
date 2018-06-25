package ttt

const (
	GridSquare = 3
	Blank    = GridSquare + 1
	Player1   = 0
	Player2   = 1
)

type Player int

type Game struct {
	board         []Player
	currentPlayer Player
}

func NewGame() *Game {
	b := make([]Player, GridSquare*GridSquare)
	for i := range b {
		b[i] = Blank
	}
	return &Game{board: b}
}
