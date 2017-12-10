package day09_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day09"
	"testing"
)

func TestProblem_PartOne1(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{}"))
	if r != "1" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "1", nil)
	}
}

func TestProblem_PartOne2(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{{}}}"))
	if r != "6" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "6", nil)
	}
}

func TestProblem_PartOne3(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{},{}}"))
	if r != "5" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "5", nil)
	}
}

func TestProblem_PartOne4(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{{},{},{{}}}}"))
	if r != "16" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "16", nil)
	}
}

func TestProblem_PartOne5(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{<a>,<a>,<a>,<a>}"))
	if r != "1" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "1", nil)
	}
}

func TestProblem_PartOne6(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{<a>},{<a>},{<a>},{<a>}}"))
	if r != "9" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "9", nil)
	}
}

func TestProblem_PartOne7(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}"))
	if r != "9" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "9", nil)
	}
}

func TestProblem_PartOne8(t *testing.T) {
	r, err := day09.Problem().PartOne([]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}"))
	if r != "3" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "3", nil)
	}
}

func TestProblem_PartTwo1(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<>"))
	if r != "0" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "0", nil)
	}
}

func TestProblem_PartTwo2(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<random characters>"))
	if r != "17" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "17", nil)
	}
}

func TestProblem_PartTwo3(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<<<<>"))
	if r != "3" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "3", nil)
	}
}

func TestProblem_PartTwo4(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<{!>}>"))
	if r != "2" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "2	", nil)
	}
}

func TestProblem_PartTwo5(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<!!>"))
	if r != "0" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "0", nil)
	}
}

func TestProblem_PartTwo6(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte("<!!!>>"))
	if r != "0" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "0", nil)
	}
}

func TestProblem_PartTwo7(t *testing.T) {
	r, err := day09.Problem().PartTwo([]byte(`<{o"i!a,<{i<a>`))
	if r != "10" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "10", nil)
	}
}
