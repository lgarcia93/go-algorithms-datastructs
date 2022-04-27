package main

type Queue[T any] struct {
	data []T
}

// Add adds item to the queue/
func (q *Queue[T]) Add(item T) {
	q.data = append([]T{item}, q.data...)
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) LasItemIndex() int {
	return q.Size() - 1
}

func (q *Queue[T]) Deque() (T, bool) {

	if q.IsEmpty() {
		return *new(T), false
	}

	item := q.data[q.LasItemIndex()]

	q.data = q.data[:q.LasItemIndex()]

	return item, true
}
