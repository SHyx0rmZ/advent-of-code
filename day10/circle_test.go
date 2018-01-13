package day10_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day10"
	"testing"
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
			c.TwistEnds(&c.Marks[test.Start], &c.Marks[test.End])
			m := &c.Marks[len(c.Marks)-1]
			for i := 0; i < len(test.Values); i++ {
				m = m.Next()
				if m.Value != test.Values[i] {
					t.Fatalf("Values[%d]: got %d, want %d", i, m.Value, test.Values[i])
				}
			}
		})
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
		c.TwistEnds(&c.Marks[0], &c.Marks[2499])
	}
}
