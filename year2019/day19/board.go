package day19

import (
	"image"
	"image/color"
	"image/gif"
)

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
