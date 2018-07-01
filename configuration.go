package ttt

type Mode int

const (
	HvH Mode = iota
	HvC
	CvC
)

func (mode Mode) Set() [2]Player {
	players := [...][2]Player {
		{&Human{}, &Human{}},
		{&Human{}, &Easy{}},
		{&Easy{}, &Easy{}},
	}
	return players[mode]
}

type Configuration struct {
	View 	*Console
	Mode 	Mode
	Players [2]Player
}

func (c *Configuration) configurePlayers(mode Mode) {
	c.Players = mode.Set()
}

func Configure() *Configuration {
	c := Configuration{View: NewConsole()}
	c.View.PlayerOptions()
	mode := c.View.GetMode()
	c.configurePlayers(mode)
	return &c
}


