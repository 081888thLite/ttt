package ttt

import (
	"reflect"
	"testing"
)

func TestConfiguration_configurePlayers(t *testing.T) {
	type args struct {
		mode Mode
	}
	tests := []struct {
		name string
		args args
		want [2]Player
	}{
		{
			name: "Configure Human vs Human game",
			args: args{mode: HvH},
			want: [2]Player{&Human{}, &Human{}},
		},
		{
			name: "Configure Human vs Computer game",
			args: args{mode: HvC},
			want: [2]Player{&Human{}, &Easy{}},
		},
		{
			name: "Configure Computer vs Computer game",
			args: args{mode: CvC},
			want: [2]Player{&Easy{}, &Easy{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{
				View: NewConsole(),
			}
			c.configurePlayers(tt.args.mode)
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
			want: &Configuration{View: &Console{UI:Sys{}}, Players: HvH.Set()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Configure();
			if !reflect.DeepEqual(got.Players, tt.want.Players) || !reflect.DeepEqual(got.Mode, tt.want.Mode) {
				t.Errorf("\nConfiguration returned:\n %v %v %v %v\nwanted:\n %v %v %v %v",
						 got.View, got.Mode, got.Players[0], got.Players[1],
						 tt.want.View, tt.want.Mode, tt.want.Players[0], tt.want.Players[1])
			}
		})
	}
}
