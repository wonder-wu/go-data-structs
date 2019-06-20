package queue

type Node struct {
	Element interface{}
	Link    *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	N    int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {
	item := &Node{data, nil}
	if q.Tail == nil {
		q.Head = item
		q.Tail = item
	} else {
		q.Tail.Link = item
		q.Tail = item
	}
	q.N++

}

func (q *Queue) Dequeue() interface{} {
	if q.Head == nil {
		return nil
	}
	item := q.Head
	q.Head = item.Link
	q.N--
	return item.Element
}

func (q *Queue) Size() int {
	return q.N
}
