package day18

import (
	"image/color"
	"math"
)

func hsl(hue, saturation, lightness float64) color.RGBA {
	c := (1 - math.Abs(2.0*lightness-1)) * saturation
	h := hue / 60.0
	x := c * (1 - math.Abs(math.Mod(h, 2)-1))
	var r, g, b float64
	switch {
	case 0 <= h && h <= 1:
		r, g = c, x
	case 1 <= h && h <= 2:
		r, g = x, c
	case 2 <= h && h <= 3:
		g, b = c, x
	case 3 <= h && h <= 4:
		g, b = x, c
	case 4 <= h && h <= 5:
		r, b = x, c
	case 5 <= h && h <= 6:
		r, b = c, x
	}
	m := lightness - c/2
	return color.RGBA{
		R: uint8((r + m) * 255.0),
		G: uint8((g + m) * 255.0),
		B: uint8((b + m) * 255.0),
		A: 255,
	}
}
