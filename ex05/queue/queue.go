package queue

type Queue []interface{}

func (q *Queue) Push(param interface{}) {
	*q = append(*q, param)
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
