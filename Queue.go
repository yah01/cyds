package cyds

type Queue struct {
	contiguousList
}

func (q *Queue) Front() Elem {
	return q.data[0]
}

func (q *Queue) Pop() Elem {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
