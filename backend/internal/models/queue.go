package models

type Queue struct {
	Items []string
}

func NewQueue() *Queue {
	return &Queue{Items: []string{}}
}

func (q *Queue) Enqueue(item string) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.Items) == 0 {
		return "", false
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item, true
}

func (q *Queue) Len() int {
	return len(q.Items)
}
