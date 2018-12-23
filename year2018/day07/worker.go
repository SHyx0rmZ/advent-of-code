package day07

type Payload func() Payload

type Queue struct {
	Size int

	payloads []Payload
}

func (q *Queue) Add(p Payload) {
	if p == nil {
		return
	}
	q.payloads = append(q.payloads, p)
}

func (q *Queue) remove(i int) {
	q.payloads = append(q.payloads[:i], q.payloads[i+1:]...)
}

func (q *Queue) Process() bool {
	if len(q.payloads) == 0 {
		return false
	}
	for i := range q.payloads {
		q.payloads[i] = q.payloads[i]()
		if q.payloads[i] == nil {
			q.remove(i)
		}
	}
	return true
}
