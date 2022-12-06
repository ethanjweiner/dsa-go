package structures

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (dequeued any) {
	if len(q.items) == 0 {
		return nil
	}

	dequeued = q.items[0]
	q.items = q.items[1:]
	return
}
