package day10

type Circle struct {
	Marks []Mark
	Ptr   int
	Skip  int
}

func (c *Circle) SelectEnds(length int) (*Mark, *Mark) {
	return &c.Marks[c.Ptr], &c.Marks[(c.Ptr+length-1)%len(c.Marks)]
}

func (c *Circle) TwistEnds(start *Mark, end *Mark) {
	//for start != end {
	//	start.Value, end.Value = end.Value, start.Value
	//
	//	start = start.Next()
	//	if start == end {
	//		break
	//	}
	//	end = end.Prev()
	//}

	//p, n := start.Prev(), end.Next()
	//p.PtrNext, end.PtrNext, n.PtrPrev, start.PtrPrev = end.PtrNext, p.PtrNext, start.PtrPrev, n.PtrPrev
	//p.PtrNext, n.PtrPrev = n.PtrPrev, p.PtrNext
	//p.PtrNext, n.PtrPrev = end, start
	//end.DirBackward.Toggle()
	//end.DirForward.Toggle()
	//start.DirBackward.Toggle()
	//start.DirForward.Toggle()
}

func (c *Circle) Round(lengths []int) {
	for _, length := range lengths {
		//c.Marks[c.Ptr].Reverse(length)
		start, end := c.SelectEnds(length)
		c.TwistEnds(start, end)
		c.Ptr = (c.Ptr + length + c.Skip) % len(c.Marks)
		c.Skip++
	}
	//c.Marks[0].Value = 3
	//c.Marks[1].Value = 4
}
