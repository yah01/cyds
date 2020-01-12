package cyDS

type Stack struct {
	contiguousList
}

func (q *Stack) Top() Elem {
	return q.data[q.size-1]
}

func (q *Stack) Pop() Elem {
	q.size--
	v := q.data[q.size]
	q.data = q.data[0:q.size]
	return v
}
