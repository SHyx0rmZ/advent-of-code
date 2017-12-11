package day11

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	var p pos
	for _, step := range bytes.Split(data, []byte(",")) {
		p = p.Add(directions[string(bytes.TrimSpace(step))])
	}
	return fmt.Sprintf("%+v", p.Distance(pos{})), nil
}

func (problem) PartTwo(data []byte) (string, error) {
	return "", errors.New("not implemented yet")
}

type pos struct {
	X, Y, Z int
}

var directions = map[string]pos{
	"se": {X: +1, Y: -1, Z: +0},
	"s":  {X: +1, Y: +0, Z: -1},
	"sw": {X: +0, Y: +1, Z: -1},
	"nw": {X: -1, Y: +1, Z: +0},
	"n":  {X: -1, Y: +0, Z: +1},
	"ne": {X: +0, Y: -1, Z: +1},
}

func (p pos) Add(o pos) pos {
	return pos{
		X: p.X + o.X,
		Y: p.Y + o.Y,
		Z: p.Z + o.Z,
	}
}

func (p pos) Distance(o pos) float64 {
	return (math.Abs(float64(p.X-o.X)) + math.Abs(float64(p.Y-o.Y)) + math.Abs(float64(p.Z-o.Z))) / 2
}
