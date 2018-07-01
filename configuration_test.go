package ttt

import (
	"reflect"
	"testing"
)

func TestConfiguration_configurePlayers(t *testing.T) {
	defaultHumanP1 := &Human{"X", &Sys{}}
	defaultHumanP2 := &Human{"O", &Sys{}}
	defaultCompE1 := &Easy{"X", &StubClient{}}
	defaultCompE2 := &Easy{"O", &StubClient{}}
	tests := []struct {
		name string
		mode Mode
		want [2]Player
	}{
		{
			name: "Configure Human vs Human game",
			mode: HvH,
			want: [2]Player{defaultHumanP1, defaultHumanP2},
		},
		{
			name: "Configure Human vs Computer game",
			mode: HvC,
			want: [2]Player{defaultHumanP1, defaultCompE2},
		},
		{
			name: "Configure Computer vs Computer game",
			mode: CvC,
			want: [2]Player{defaultCompE1, defaultCompE2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{
				View: NewConsole(),
			}
			c.configurePlayers(tt.mode.players())
			if got := c.Players; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n %v failed,\ngot:\n%T %T\nwanted config Like:\n %T %T",
					tt.name, got[0], got[1], tt.want[0], tt.want[1])
			}
		})
	}
}

func TestConfigure(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		{
			name: "Configure runs ok",
			want: &Configuration{View: &Console{UI: Sys{}}, Players: Mode(0).players()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Configure()
			if !reflect.DeepEqual(got.Players, tt.want.Players) || !reflect.DeepEqual(got.Mode, tt.want.Mode) {
				t.Errorf("\nConfiguration returned:\n %v %v %v %v\nwanted:\n %v %v %v %v",
					got.View, got.Mode, got.Players[0], got.Players[1],
					tt.want.View, tt.want.Mode, tt.want.Players[0], tt.want.Players[1])
			}
		})
	}
}
