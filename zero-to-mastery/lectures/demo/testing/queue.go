package queue

type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

func (q *Queue) Append(items ...int) bool {
	if len(items)+len(q.items) > q.capacity {
		return false
	}

	q.items = append(q.items, items...)
	return true
}

func (q *Queue) Next() (int, bool) {
	if len(q.items) <= 0 {
		return 0, false
	}

	next := q.items[0]
	q.items = q.items[1:]
	return next, true
}
