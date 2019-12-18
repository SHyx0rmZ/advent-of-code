package day08

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	l := make([]byte, 25*6)
	var m [3]int
	m[0] = 25 * 6
	for {
		n, err := r.Read(l)
		if n > 0 {
			l = l[:n]
			if len(l) == 1 && l[0] == 10 {
				continue
			}
			var c [3]int
			for _, b := range l {
				switch b {
				case '0':
					c[0]++
				case '1':
					c[1]++
				case '2':
					c[2]++
				}
			}
			if c[0] < m[0] {
				m = c
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}

	return fmt.Sprintf("%d", m[1]*m[2]), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	l := make([]byte, 25*6)
	pa := color.Palette{
		color.Gray{Y: 0},
		color.Gray{Y: 255},
		color.Alpha{A: 0},
	}
	i := &gif.GIF{
		Image:     nil,
		Delay:     nil,
		LoopCount: -1,
		Disposal:  nil,
		Config: image.Config{
			ColorModel: pa,
			Width:      25,
			Height:     6,
		},
		BackgroundIndex: 2,
	}
	re := image.Rect(0, 0, 25, 6)
	for {
		n, err := r.Read(l)
		if n > 0 {
			l = l[:n]
			if len(l) == 1 && l[0] == 10 {
				continue
			}
			il := image.NewPaletted(re, pa)
			for a, b := range l {
				il.SetColorIndex(a%25, a/25, b-'0')
			}
			i.Image = append(i.Image, il)
			i.Delay = append(i.Delay, 1)
			i.Disposal = append(i.Disposal, gif.DisposalNone)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	var is []*image.Paletted
	for ii := range i.Image {
		is = append(is, i.Image[len(i.Image)-ii-1])
	}
	i.Image = is
	f, err := os.Create("day08.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = gif.EncodeAll(f, i)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%d", 0), nil
}
