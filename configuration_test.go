package ttt

import (
	"reflect"
	"testing"
)

func TestConfigure(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		{
			name: "configure runs ok w/ Defaults",
			want: &Configuration{View: &Console{UI: Sys{}}, Players: DefaultPlayers},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Configure()
			if !reflect.DeepEqual(got.Players, tt.want.Players) || !reflect.DeepEqual(tt.want.Players, got.Players) {
				t.Errorf("\n%v %T %T\nwanted:\n%T %T",
					tt.name,
					tt.want.Players[0], tt.want.Players[1],
					got.Players[0], got.Players[1])
			}
		})
	}
}

func TestStrategy_setPlayer(t *testing.T) {
	type args struct {
		piece Piece
	}
	tests := []struct {
		name     string
		strategy Strategy
		args     args
		want     Player
	}{
		{
			name:     "returns a Human Player",
			strategy: HUMAN,
			args:     args{piece: "x"},
			want:     &Human{Piece: "X"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.strategy.getPlayer(tt.args.piece); got == tt.want {
				t.Errorf("Strategy.getPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
