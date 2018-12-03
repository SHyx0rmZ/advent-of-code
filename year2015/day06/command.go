package day06

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s := NewParser(r)
	i := image.NewGray(image.Rect(0, 0, 1000, 1000))
	go func() {
		for ops := range s.Operations() {
			switch op := ops.(type) {
			case Toggle:
				for x := op.From.X; x <= op.To.X; x++ {
					for y := op.From.Y; y <= op.To.Y; y++ {
						c := i.GrayAt(x, y)
						switch c {
						case color.Gray{255}:
							i.Set(x, y, color.Black)
						case color.Gray{0}:
							i.Set(x, y, color.White)
						}
					}
				}
			case TurnOff:
				draw.Draw(i, image.Rect(op.From.X, op.From.Y, op.To.X+1, op.To.Y+1), image.Black, image.ZP, draw.Src)
			case TurnOn:
				draw.Draw(i, image.Rect(op.From.X, op.From.Y, op.To.X+1, op.To.Y+1), image.White, image.ZP, draw.Src)
			}
		}
	}()
	err := s.Parse()
	var c int
	for x := 0; x <= 999; x++ {
		for y := 0; y <= 999; y++ {
			if i.GrayAt(x, y) == color.GrayModel.Convert(color.White) {
				c++
			}
		}
	}
	f, e := os.Create(filepath.Join(os.Getenv("HOME"), "Desktop", "aoc-2015-06a.png"))
	if e != nil {
		panic(err)
	}
	defer f.Close()
	e = png.Encode(f, i)
	if e != nil {
		panic(e)
	}
	return strconv.Itoa(c), err
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s := NewParser(r)
	i := image.NewGray(image.Rect(0, 0, 1000, 1000))
	go func() {
		for ops := range s.Operations() {
			switch op := ops.(type) {
			case Toggle:
				for x := op.From.X; x <= op.To.X; x++ {
					for y := op.From.Y; y <= op.To.Y; y++ {
						c := i.GrayAt(x, y)
						c.Y += 2
						i.SetGray(x, y, c)
					}
				}
			case TurnOff:
				for x := op.From.X; x <= op.To.X; x++ {
					for y := op.From.Y; y <= op.To.Y; y++ {
						c := i.GrayAt(x, y)
						if c.Y > 0 {
							c.Y -= 1
						}
						i.SetGray(x, y, c)
					}
				}
			case TurnOn:
				for x := op.From.X; x <= op.To.X; x++ {
					for y := op.From.Y; y <= op.To.Y; y++ {
						c := i.GrayAt(x, y)
						c.Y += 1
						i.SetGray(x, y, c)
					}
				}
			}
		}
	}()
	err := s.Parse()
	var c int
	for x := 0; x <= 999; x++ {
		for y := 0; y <= 999; y++ {
			c += int(i.GrayAt(x, y).Y)
		}
	}
	f, e := os.Create(filepath.Join(os.Getenv("HOME"), "Desktop", "aoc-2015-06b.png"))
	if e != nil {
		panic(err)
	}
	defer f.Close()
	e = png.Encode(f, i)
	if e != nil {
		panic(e)
	}
	return strconv.Itoa(c), err
}
