package day19

import (
	"fmt"
	aoc "github.com/SHyx0rmZ/advent-of-code"
	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
	"strconv"
)

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
	b := make(board)
	g := gif.GIF{
		Config: image.Config{
			ColorModel: color.Palette{
				color.Black,
				color.White,
			},
			Width:  50,
			Height: 50,
		},
	}
	for y := 0; y < g.Config.Height; y++ {
		for x := 0; x < g.Config.Width; x++ {
			output := make(chan int, 1)
			go prg.Run(input, output)
			input <- x
			input <- y
			d := <-output
			if d == 0 {
				continue
			}
			b[point{x, y}] = d
		}
	}
	b.Print(&g)
	f, err := os.Create("day19.gif")
	if err != nil {
		return "", err
	}
	defer f.Close()
	defer gif.EncodeAll(f, &g)
	return strconv.Itoa(len(b)), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	prg, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	beam := func(x, y int) bool {
		output := make(chan int)
		go prg.Run(input, output)
		input <- x
		input <- y
		return <-output == 1
	}
	var y, minX, maxX int
	for {
		line := make([]bool, 0, maxX-minX+1)
		var present bool
		for x := minX; x <= maxX; x++ {
			pulled := beam(x, y)
			line = append(line, pulled)
			if pulled {
				if x == maxX {
					maxX++
				}
				present = true
			}
		}
		fmt.Println(y, minX, len(line), line[0], line[1], line[len(line)-2], line[len(line)-1])
		if !present {
			minX++
			maxX++
		} else {
			for _, x := range line {
				if x {
					break
				}
				minX++
			}
		}
		y++
		if y < 1000 {
			continue
		}
		return "", nil
	}
}
