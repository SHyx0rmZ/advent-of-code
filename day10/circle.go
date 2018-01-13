package day10

type Circle struct {
	Marks []Mark
	Ptr   int
	Skip  int
}

func (c *Circle) SelectEnds(length int) (*Mark, *Mark) {
	return &c.Marks[c.Ptr], &c.Marks[(c.Ptr+length-1)%len(c.Marks)]
}

func (c *Circle) TwistEnds(start *Mark, end *Mark, l int) {
	//for start != end {
	//	start.Value, end.Value = end.Value, start.Value
	//
	//	start = start.Next()
	//	if start == end {
	//		break
	//	}
	//	end = end.Prev()
	//}

	it := Iter{}
	p, n := it.Prev(start), it.Next(end)
	p.Next, n.Prev = end, start
	start.Prev, end.Next = n, p
	start.Bkwd = l
	end.Fwd = l
}

func (c *Circle) Round(lengths []int) {
	for _, length := range lengths {
		//c.Marks[c.Ptr].Reverse(length)
		start, end := c.SelectEnds(length)
		c.TwistEnds(start, end, length)
		c.Ptr = (c.Ptr + length + c.Skip) % len(c.Marks)
		c.Skip++
	}
	//c.Marks[0].Value = 3
	//c.Marks[1].Value = 4
}
