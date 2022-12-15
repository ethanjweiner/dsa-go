package structures

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (dequeued T) {
	if len(q.items) == 0 {
		return
	}

	dequeued = q.items[0]
	q.items = q.items[1:]
	return
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.items)
}
