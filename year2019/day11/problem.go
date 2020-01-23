package day11

import (
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
	"strconv"
)

type point struct {
	X, Y int
}

func (p point) Less(o point) bool {
	if p.Y == o.Y {
		return p.X < o.X
	}
	return p.Y < o.Y
}

func (p point) Move(d direction) point {
	switch d {
	case north:
		return point{p.X, p.Y - 1}
	case south:
		return point{p.X, p.Y + 1}
	case west:
		return point{p.X - 1, p.Y}
	case east:
		return point{p.X + 1, p.Y}
	default:
		panic("invalid direction")
	}
}

type direction int

const (
	north direction = iota
	south
	west
	east
)

func (d direction) Left() direction {
	switch d {
	case north:
		return west
	case south:
		return east
	case west:
		return south
	case east:
		return north
	default:
		panic("invalid direction")
	}
}

func (d direction) Right() direction {
	switch d {
	case north:
		return east
	case south:
		return west
	case west:
		return north
	case east:
		return south
	default:
		panic("invalid direction")
	}
}

type board map[point]int

func (b board) Bounds() image.Rectangle {
	if len(b) == 0 {
		return image.Rectangle{}
	}
	ps := make([]point, 0, len(b))
	for p := range b {
		ps = append(ps, p)
	}
	min := ps[0]
	max := ps[0]
	for _, p := range ps[1:] {
		if p.Less(min) {
			min = p
		}
		if max.Less(p) {
			max = p
		}
	}
	return image.Rect(min.X, min.Y, max.X+1, max.Y+1)
}

func (b board) Print(g *gif.GIF) {
	bounds := b.Bounds().Sub(b.Bounds().Min)
	i := image.NewPaletted(bounds, g.Config.ColorModel.(color.Palette))
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			t := b[point{x + bounds.Min.X, y + bounds.Min.Y}]
			i.SetColorIndex(x, y, uint8(t))
		}
	}
	g.Image = append(g.Image, i)
	g.Delay = append(g.Delay, 1)
	g.Disposal = append(g.Disposal, gif.DisposalBackground)
}

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	go prg.Run(input, output)
	panels := make(board)
	var robot struct {
		point
		direction
	}
	for {
		input <- panels[robot.point]
		color, ok := <-output
		if !ok {
			break
		}
		panels[robot.point] = color
		switch <-output {
		case 0:
			robot.direction = robot.Left()
		case 1:
			robot.direction = robot.Right()
		}
		robot.point = robot.Move(robot.direction)
	}
	f, err := os.Create("day11a.gif")
	if err != nil {
		return "", err
	}
	defer f.Close()
	g := &gif.GIF{
		LoopCount: -1,
		Config: image.Config{
			ColorModel: color.Palette{
				color.Black,
				color.White,
			},
			Width:  panels.Bounds().Dx(),
			Height: panels.Bounds().Dy(),
		},
	}
	defer gif.EncodeAll(f, g)
	panels.Print(g)
	return strconv.Itoa(len(panels)), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	output := make(chan int, 1)
	go prg.Run(input, output)
	panels := make(board)
	var robot struct {
		point
		direction
	}
	panels[robot.point] = 1
	for {
		input <- panels[robot.point]
		color, ok := <-output
		if !ok {
			break
		}
		panels[robot.point] = color
		switch <-output {
		case 0:
			robot.direction = robot.Left()
		case 1:
			robot.direction = robot.Right()
		}
		robot.point = robot.Move(robot.direction)
	}
	f, err := os.Create("day11b.gif")
	if err != nil {
		return "", err
	}
	defer f.Close()
	g := &gif.GIF{
		LoopCount: -1,
		Config: image.Config{
			ColorModel: color.Palette{
				color.Black,
				color.White,
			},
			Width:  panels.Bounds().Dx(),
			Height: panels.Bounds().Dy(),
		},
	}
	defer gif.EncodeAll(f, g)
	panels.Print(g)
	return strconv.Itoa(len(panels)), nil
}
