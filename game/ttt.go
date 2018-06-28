package game

const (
	Blank = " "
	NoOne = Blank
)

//Types related to the Board
type BoardSize int
type Board []Piece

//Types related to Players
type Piece string
type Strategy interface {
	GetMove(ui Client, board Board) int
}

//Core Types
type Player struct {
	Piece    Piece
	Strategy Strategy
}

type Game struct {
	Board         Board
	CurrentPlayer Player
	Winner        Piece
	Players       [2]Player
	UI            Client
}

func NewGame(boardSize BoardSize, player1 Player, player2 Player) *Game {
	var b Board
	for i := 0; i < int(boardSize); i++ {
		b = append(b, Blank)
	}
	return &Game{
		Board:         b,
		Players:       [2]Player{player1, player2},
		CurrentPlayer: player1,
		Winner:        NoOne,
	}
}

func (game *Game) setPlayers(player1 Player, player2 Player) *Game {
	game.Players = [2]Player{player1, player2}
	return game
}

func (game *Game) switchPlayers() *Game {
	if game.CurrentPlayer == game.Players[0] {
		game.CurrentPlayer = game.Players[1]
	} else {
		game.CurrentPlayer = game.Players[0]
	}
	return game
}

func (game *Game) mark(position int) *Game {
	if game.Board[position] == Blank {
		game.Board[position] = game.CurrentPlayer.Piece
	}
	return game
}

func (game *Game) wonBy(cell1, cell2, cell3 int) bool {
	if b := game.Board; b[cell1] != NoOne {
		return b[cell1] == b[cell2] && b[cell2] == b[cell3]
	} else {
		return false
	}
}

func (game *Game) checkDiagonalWin() {
	leftToRightWin, rightToLeftWin := game.wonBy(0, 4, 8), game.wonBy(6, 4, 2)
	if leftToRightWin {
		game.Winner = game.Board[0]
	} else if rightToLeftWin {
		game.Winner = game.Board[6]
	}
}

func (game *Game) checkRowWin() {
	row1Win, row2Win, row3Win := game.wonBy(0, 1, 2), game.wonBy(3, 4, 5), game.wonBy(6, 7, 8)
	if row1Win {
		game.Winner = game.Board[0]
	} else if row2Win {
		game.Winner = game.Board[3]
	} else if row3Win {
		game.Winner = game.Board[6]
	}
}

func (game *Game) checkColumnWin() {
	col1Win, col2Win, col3Win := game.wonBy(0, 3, 6), game.wonBy(1, 4, 7), game.wonBy(2, 5, 8)
	if col1Win {
		game.Winner = game.Board[0]
	} else if col2Win {
		game.Winner = game.Board[1]
	} else if col3Win {
		game.Winner = game.Board[2]
	}
}

func (game *Game) CheckForWin() {
	game.checkDiagonalWin()
	game.checkRowWin()
	game.checkColumnWin()
}

func (game *Game) boardFull() bool {
	for _, a := range game.Board {
		if a == Blank {
			return false
		}
	}
	return true
}
