package ttt

const (
	HUMAN Strategy = iota + 1
	EASY
	MEDIUM
	HARD
)

var DefaultPlayers = [2]Player{&Hard{"x", &Sys{}}, *HUMAN.create("o")}

type Configuration struct {
	View    *Console
	Players [2]Player
}
type Strategy int

func (strategy Strategy) create(piece Piece) *Player {
	players := [...]Player{
		&Human{piece, &Sys{}, Console{}},
		&Easy{piece, &StubClient{}},
		&Medium{piece, &StubClient{}},
		&Hard{piece, &StubClient{}},
	}
	player := players[strategy-1]
	return &player
}

func Configure() *Configuration {
	var setPlayers [2]Player
	c := Configuration{View: NewConsole()}
	v := c.View
	v.GameMenu()
	if v.WantsSetup() {
		for i, _ := range setPlayers {
			strategy, piece := v.PlayerMenu(i)
			if i == 1 {
				if setPlayers[0].GetPiece() == piece {
					v.CantHaveSamePiece()
					v.PlayerMenu(i)
				}
			}
			setPlayers[i] = *strategy.create(piece)
		}
	} else {
		setPlayers = DefaultPlayers
	}
	c.Players = setPlayers
	return &c
}
