package day10

type Mark struct {
	// Value stores an integer Value.
	Value int

	// PtrNext points to the next Mark in the list.
	PtrNext *Mark

	// PtrPrev points to the previous Mark in the list.
	PtrPrev *Mark

	//
	DirForward  Direction
	DirBackward Direction
}

func NewMarks(n int) []Mark {
	var marks []Mark
	for i := 0; i < n; i++ {
		marks = append(marks, Mark{
			Value: i,
		})
	}
	for i := 0; i < n-1; i++ {
		marks[i].PtrNext = &marks[i+1]
		marks[i+1].PtrPrev = &marks[i]
	}
	if marks != nil {
		marks[n-1].PtrNext = &marks[0]
		marks[0].PtrPrev = &marks[n-1]
	}
	return marks
}

func (m *Mark) Next() *Mark {
	return m.DirBackward.Next(m)
}

func (m *Mark) Prev() *Mark {
	return m.DirBackward.Prev(m)
}
