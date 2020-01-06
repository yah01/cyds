package cyDS

type Stack struct {
	contiguousList
}

func (q *Stack) Top() Elem {
	return q.data[q.Size-1]
}

func (q *Stack) Push(v Elem) {
	if q.Size == 0 {
		q.data = []Elem{v}
	} else {
		q.data = append(q.data, v)
	}
	q.Size++
}

func (q *Stack) Pop() Elem {
	q.Size--
	v := q.data[q.Size]
	q.data = q.data[0:q.Size]
	return v
}
