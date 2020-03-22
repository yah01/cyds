package cyds

type Stack struct {
	contiguousList
}

func (q *Stack) Top() Elem {
	return q.data[len(q.data)-1]
}

func (q *Stack) Pop() Elem {
	length := len(q.data)
	v := q.data[length-1]
	q.data = q.data[0 : length-1]
	return v
}
