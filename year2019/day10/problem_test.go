package day10_test

import (
"strings"
"testing"

"github.com/SHyx0rmZ/advent-of-code/year2019/day10"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day10.Problem().PartOneWithReader(strings.NewReader(`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
`))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day10.Problem().PartTwoWithReader(strings.NewReader(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}


