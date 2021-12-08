package game

import (
	"reflect"
	"testing"
)

func TestNewGeneration(t *testing.T) {
	type tcase struct {
		liveCells []Coordinate
		want      Generation
	}

	tests := map[string]tcase{
		"No liveCells": {
			liveCells: nil,
			want:      Generation{},
		},

		"With liveCells": {
			liveCells: []Coordinate{{X: 0, Y: 0}},
			want:      Generation{[25]byte{1}},
		},

		// TODO: Add more cases here.
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := NewGeneration(tc.liveCells...); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NewGeneration() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGeneration_isCellAlive(t *testing.T) {
	type tcase struct {
		g    Generation
		x, y int
		want byte
	}
	tests := map[string]tcase{
		"Alive": {
			g: Generation{[25]byte{1}},
			x: 0, y: 0, want: 1,
		},

		"Dead": {
			g: Generation{[25]byte{0}},
			x: 0, y: 0, want: 0,
		},

		// TODO: Add more cases here.
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tc.g.isCellAlive(tc.x, tc.y); got != tc.want {
				t.Errorf("isCellAlive() = %v, want %v", got, tc.want)
			}
		})
	}
}
