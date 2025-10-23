package datastruct

type Queue[T any] struct {
	Data []T
	Size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(d T) {
	q.Data = append(q.Data, d)
}

func (q *Queue[T]) Dequeue() T {
	data := q.Data[0]
	q.Data = q.Data[:1]
	return data
}

func (q *Queue[T]) Len() int {
	return len(q.Data)
}

func (q *Queue[T]) IsEmpty() bool {
	if q.Len() != 0 {
		return false
	}
	return true
}

func (q *Queue[T]) Peek() T {
	return q.Data[0]
}
