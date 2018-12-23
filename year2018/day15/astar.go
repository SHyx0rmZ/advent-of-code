package day15

import (
	"sort"
)

const infinity = 999999999

func aStar(start, goal point, m map[point]*tile) []point {
	visited := make(map[point]struct{})
	toVisit := make(map[point]struct{})
	toVisit[start] = struct{}{}
	for p, t := range m {
		if t.terrain == cave && t.unit == nil {
			toVisit[p] = struct{}{}
		}
	}
	distance := func(a, b point) int {
		x := b.X - a.X
		y := b.Y - a.Y
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		return x + y
	}
	heuristic := distance
	cameFrom := make(map[point]point)
	scores := make(map[point]int)
	scores[start] = 0
	estimatedScores := make(map[point]int)
	estimatedScores[start] = heuristic(start, goal)
	reconstructPath := func(current point) []point {
		var ps []point
		for {
			p, ok := cameFrom[current]
			if !ok {
				break
			}
			current = p
			ps = append(ps, current)
		}
		return ps
	}
	for len(toVisit) > 0 {
		var minimum struct {
			Score int
			Point point
		}
		minimum.Score = infinity
		var ps pointSlice
		for t := range toVisit {
			ps = append(ps, t)
		}
		sort.Sort(sort.Reverse(ps))
		for _, p := range ps {
			estimatedScore, ok := estimatedScores[p]
			if ok && estimatedScore <= minimum.Score {
				minimum.Score = estimatedScore
				minimum.Point = p
			}
			if !ok && minimum.Score == infinity {
				minimum.Point = p
			}
		}
		current := minimum.Point
		if current == goal {
			return reconstructPath(current)
		}
		delete(toVisit, current)
		visited[current] = struct{}{}
		for _, neighbor := range []point{
			{current.X, current.Y - 1},
			{current.X - 1, current.Y},
			{current.X + 1, current.Y},
			{current.X, current.Y + 1},
		} {
			if _, ok := visited[neighbor]; ok {
				continue
			}
			score := scores[current] + distance(current, neighbor)
			cameFrom[neighbor] = current
			scores[neighbor] = score
			estimatedScores[neighbor] = score + heuristic(neighbor, goal)
		}
	}
	return nil
}
