package queue

type node struct {
	element interface{}
	link    *node
}

type queue struct {
	head *node
	tail *node
	n    int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Enqueue(data interface{}) {
	item := &node{data, nil}
	if q.tail == nil {
		q.head = item
		q.tail = item
	} else {
		q.tail.link = item
		q.tail = item
	}
	q.n++

}

func (q *queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	item := q.head
	q.head = item.link
	q.n--
	return item.element
}

func (q *queue) Size() int {
	return q.n
}
