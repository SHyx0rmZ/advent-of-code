package day11

import "math"

type position struct {
	X, Y, Z float64
}

var directions = map[string]position{
	"se": {X: +1, Y: -1, Z: +0},
	"s":  {X: +1, Y: +0, Z: -1},
	"sw": {X: +0, Y: +1, Z: -1},
	"nw": {X: -1, Y: +1, Z: +0},
	"n":  {X: -1, Y: +0, Z: +1},
	"ne": {X: +0, Y: -1, Z: +1},
}

func (p position) Add(o position) position {
	return position{
		X: p.X + o.X,
		Y: p.Y + o.Y,
		Z: p.Z + o.Z,
	}
}

func (p position) Distance(o position) float64 {
	return (math.Abs(p.X-o.X) + math.Abs(p.Y-o.Y) + math.Abs(p.Z-o.Z)) / 2
}
