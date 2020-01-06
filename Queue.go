package cyDS


type Queue struct {
	contiguousList
}

func (q *Queue) Front() Elem {
	return q.data[0]
}

func (q *Queue) Push(v Elem) {
	if q.Size == 0 {
		q.data = []Elem{v}
	} else {
		q.data = append(q.data, v)
	}
	q.Size++
}

func (q *Queue) Pop() Elem {
	v := q.data[0]
	q.data = q.data[1:q.Size]
	q.Size--
	return v
}
