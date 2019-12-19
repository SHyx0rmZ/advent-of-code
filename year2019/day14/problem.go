package day14

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
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

type tile int

const (
	empty tile = iota
	wall
	block
	paddle
	ball
)

func (t tile) String() string {
	s, ok := map[tile]string{
		empty:  " ",
		wall:   "+",
		block:  "#",
		paddle: "^",
		ball:   "o",
	}[t]
	if !ok {
		return "?"
	}
	return s
}

type board map[point]tile

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
	bounds := b.Bounds()
	i := image.NewPaletted(bounds, g.Config.ColorModel.(color.Palette))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			i.SetColorIndex(x, y, uint8(b[point{x, y}]))
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
	output := make(chan int, 1)
	go prg.Run(nil, output)
	b := make(board)
	for {
		x, ok := <-output
		if !ok {
			break
		}
		y := <-output
		t := tile(<-output)
		b[point{x, y}] = t
	}
	var c int
	for _, t := range b {
		if t != block {
			continue
		}
		c++
	}
	return strconv.Itoa(c), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	prg[0] = 2
	input := make(chan int, 1)
	output := make(chan int, 1)
	go prg.Run(input, output)
	var ba, pa point
	b := make(board)
	for {
		x := <-output
		if x == -1 {
			break
		}
		y := <-output
		t := tile(<-output)
		b[point{x, y}] = t
		switch t {
		case ball:
			ba = point{x, y}
		case paddle:
			pa = point{x, y}
		}
	}
	g := &gif.GIF{
		LoopCount: -1,
		Config: image.Config{
			ColorModel: color.Palette{
				color.Black,
				color.Gray{Y: 128},
				color.Gray{Y: 64},
				color.Gray{Y: 255},
				color.Gray{Y: 255},
			},
			Width:  b.Bounds().Dx(),
			Height: b.Bounds().Dy(),
		},
	}
	f, err := os.Create("day14.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer gif.EncodeAll(f, g)
	b.Print(g)
	input <- 0
	var score int
	for {
		<-output
		score = <-output
		for {
			x, ok := <-output
			if !ok {
				b.Print(g)
				return strconv.Itoa(score), nil
			}
			if x == -1 {
				break
			}
			y := <-output
			t := tile(<-output)
			b[point{x, y}] = t

			switch t {
			case ball:
				ba = point{x, y}
			case paddle:
				pa = point{x, y}
			}

			if t != ball {
				continue
			}

			var cmd int
			switch {
			case ba.X < pa.X:
				cmd = -1
			case ba.X > pa.X:
				cmd = 1
			}

			select {
			case input <- cmd:
				b.Print(g)
			default:
			}
		}
	}
}
