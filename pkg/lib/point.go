package lib

import (
	stdsort "sort"
)

type Point struct {
	X, Y int
}

func (p Point) Less(o Point) bool {
	if p.Y == o.Y {
		return p.X < o.X
	}
	return p.Y < o.Y
}

type PointSlice []Point

func (p PointSlice) Len() int {
	return len(p)
}

func (p PointSlice) Less(i, j int) bool {
	return p[i].Less(p[j])
}

func (p PointSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortPoints(p []Point) {
	stdsort.Sort(PointSlice(p))
}
