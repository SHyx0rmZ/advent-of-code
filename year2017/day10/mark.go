package day10

type Mark struct {
	// Value stores an integer Value.
	Value int

	// Next points to the next Mark in the list.
	Next *Mark

	// Prev points to the previous Mark in the list.
	Prev *Mark

	//
	Fwd  int
	Bkwd int
}

func NewMarks(n int) []Mark {
	var marks []Mark
	for i := 0; i < n; i++ {
		marks = append(marks, Mark{
			Value: i,
		})
	}
	for i := 0; i < n-1; i++ {
		marks[i].Next = &marks[i+1]
		marks[i+1].Prev = &marks[i]
	}
	if marks != nil {
		marks[n-1].Next = &marks[0]
		marks[0].Prev = &marks[n-1]
	}
	return marks
}
