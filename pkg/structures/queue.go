package structures

type QueueStr struct {
	items []string
}

func (q *QueueStr) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Dequeue removes the first item from the queue and returns it
// Boolean value indicates whether the queue was empty
func (q *QueueStr) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *QueueStr) IsEmpty() bool {
	return len(q.items) == 0
}
