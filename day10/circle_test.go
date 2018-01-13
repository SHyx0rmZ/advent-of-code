package day10_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day10"
)

func TestCircle_SelectEnds(t *testing.T) {
	tests := []struct {
		Name   string
		Marks  int
		Ptr    int
		Length int
		Start  int
		End    int
	}{
		{
			Name:   "4: 0 -> 3",
			Marks:  4,
			Ptr:    0,
			Length: 3,
			Start:  0,
			End:    2,
		},
		{
			Name:   "4: 2 -> 3",
			Marks:  4,
			Ptr:    2,
			Length: 3,
			Start:  2,
			End:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := day10.Circle{
				Marks: day10.NewMarks(test.Marks),
				Ptr:   test.Ptr,
				Skip:  0,
			}
			start, end := c.SelectEnds(test.Length)
			if start != &c.Marks[test.Start] || end != &c.Marks[test.End] {
				t.Errorf("got (%p, %p), want (%p, %p)", start, end, &c.Marks[test.Start], &c.Marks[test.End])
			}
		})
	}
}

func BenchmarkCircle_SelectEnds(b *testing.B) {
	c := day10.Circle{
		Marks: day10.NewMarks(5000),
		Ptr:   0,
		Skip:  0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.SelectEnds(2500)
	}
}

func TestCircle_TwistEnds(t *testing.T) {
	tests := []struct {
		Name   string
		Start  int
		End    int
		Values []int
	}{
		{
			Name:   "3: 0 -> 3",
			Start:  0,
			End:    2,
			Values: []int{2, 1, 0},
		},
		{
			Name:   "5: 0 -> 3",
			Start:  0,
			End:    2,
			Values: []int{2, 1, 0, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			c := day10.Circle{
				Marks: day10.NewMarks(len(test.Values)),
			}
			c.TwistEnds(&c.Marks[test.Start], &c.Marks[test.End], test.End-test.Start+1)
			m := &c.Marks[2]
			it := day10.Iter{
				Direction: day10.Forward,
			}
			//ctr := day10.Counter(3)
			//go ctr.Run()
			//it.Set.Add(ctr)
			for i := 0; i < len(test.Values); i++ {
				if m.Value != test.Values[i] {
					t.Fatalf("Values[%d]: got %d, want %d", i, m.Value, test.Values[i])
				}
				m = it.Next(m)
			}
		})
	}
}

func TestCircle_TwistEnds2(t *testing.T) {
	c := day10.Circle{
		Marks: day10.NewMarks(7),
	}
	c.TwistEnds(&c.Marks[1], &c.Marks[4], 4)
	vars := []struct {
		Actual, Expected *day10.Mark
	}{
		{c.Marks[0].Next, &c.Marks[4]}, // 0
		{c.Marks[1].Next, &c.Marks[2]},
		{c.Marks[2].Next, &c.Marks[3]},
		{c.Marks[3].Next, &c.Marks[4]},
		{c.Marks[4].Next, &c.Marks[0]}, // 4
		{c.Marks[5].Next, &c.Marks[6]},
		{c.Marks[6].Next, &c.Marks[0]},
		{c.Marks[0].Prev, &c.Marks[6]},
		{c.Marks[1].Prev, &c.Marks[5]}, // 8
		{c.Marks[2].Prev, &c.Marks[1]},
		{c.Marks[3].Prev, &c.Marks[2]},
		{c.Marks[4].Prev, &c.Marks[3]},
		{c.Marks[5].Prev, &c.Marks[1]}, // 12
		{c.Marks[6].Prev, &c.Marks[5]},
	}
	for i, v := range vars {
		if v.Actual != v.Expected {
			t.Errorf("%d: got %p, want %p", i, v.Actual, v.Expected)
		}
	}
	vals := []struct {
		F func(*day10.Iter, *day10.Mark) *day10.Mark
		V []int
	}{
		{(*day10.Iter).Next, []int{0, 4, 3, 2, 1, 5, 6, 0}},
		{(*day10.Iter).Prev, []int{0, 6, 5, 1, 2, 3, 4, 0}},
	}
	for i, v := range vals {
		m := &c.Marks[0]
		it := &day10.Iter{
			Direction: day10.Forward,
		}
		for j := 0; j < len(v.V); j++ {
			if m.Value != v.V[j] {
				t.Errorf("%d: Values[%d]: got %d, want %d", i, j, m.Value, v.V[j])
			}
			m = v.F(it, m)
		}
	}
}

func BenchmarkCircle_TwistEnds(b *testing.B) {
	c := day10.Circle{
		Marks: day10.NewMarks(5000),
		Ptr:   0,
		Skip:  0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.TwistEnds(&c.Marks[0], &c.Marks[2499], 2500)
	}
}
