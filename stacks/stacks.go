package stacks

type node struct {
	element interface{}
	link    *node
}

type stack struct {
	head *node
	n    int
}

func NewLinkedStack() *stack {
	return &stack{}
}

func (s *stack) Push(data interface{}) {
	item := &node{element: data, link: s.head}
	s.head = item
	s.n++
}
func (s *stack) Pop() interface{} {
	if s.head == nil {
		return nil
	}
	item := s.head
	s.head = item.link
	s.n--
	return item.element
}

func (s *stack) Size() int {
	return s.n
}

type arrayStack struct {
	container []interface{}
	len       int
}

func NewArrayStack() *arrayStack {
	return &arrayStack{nil, 0}
}

func (s *arrayStack) Push(data interface{}) {
	s.container = append(s.container, data)
	s.len++
}
func (s *arrayStack) Pop() interface{} {
	item := s.container[s.len-1]
	s.container = s.container[:s.len-1]
	s.len--
	return item
}
