package ttt

const (
	CellsPerSide = 3
	BoardSize = CellsPerSide*CellsPerSide
	Blank      = " "
	Player1    = "X"
	Player2    = "O"
	NoOne = Blank
)

type Player string

type Game struct {
	Board         [BoardSize]Player
	CurrentPlayer Player
	Winner Player
}

func NewGame() *Game {
	return &Game{
		Board: [BoardSize]Player{Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank},
		CurrentPlayer: Player1,
		Winner: NoOne,
	}
}

func (game *Game) switchPlayers() *Game {
	if game.CurrentPlayer == Player1 {
		game.CurrentPlayer = Player2
	} else {
		game.CurrentPlayer = Player1
	}
	return &Game{}
}

func (game *Game) mark(position int) *Game {
	game.Board[position] = game.CurrentPlayer
	return &Game{}
}

func (game *Game) isWonByCells(cell1, cell2, cell3 int) bool {
	var isWon bool
	if b := game.Board; b[cell1] != NoOne {
		isWon = b[cell1] == b[cell2] && b[cell2] == b[cell3]
	}
	return isWon
}

func (game *Game) checkDiagonalWin() {
	leftToRightWin := game.isWonByCells(0,4,8)
	rightToLeftWin := game.isWonByCells(6,4,2)
	if leftToRightWin { game.Winner = game.Board[0]
	} else if rightToLeftWin { game.Winner = game.Board[6] }
}

func (game *Game) checkRowWin() {
	row1Win := game.isWonByCells(0,1,2)
	row2Win := game.isWonByCells(3,4,5)
	row3Win := game.isWonByCells(6,7,8)
	if row1Win { game.Winner = game.Board[0]
	} else if row2Win { game.Winner = game.Board[3]
	} else if row3Win { game.Winner = game.Board[6] }
}

func (game *Game) checkColumnWin() {
	col1Win := game.isWonByCells(0,3,6)
	col2Win := game.isWonByCells(1,4,7)
	col3Win := game.isWonByCells(2,5,8)
	if col1Win { game.Winner = game.Board[0]
	} else if col2Win { game.Winner = game.Board[1]
	} else if col3Win { game.Winner = game.Board[2] }
}

func (game *Game) winner() Player {
	game.checkDiagonalWin()
	game.checkRowWin()
	game.checkColumnWin()
	return game.Winner;
}