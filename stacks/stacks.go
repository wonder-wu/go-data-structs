package Stacks

type Node struct {
	Element interface{}
	Link    *Node
}

type Stack struct {
	Head *Node
	N    int
}

func NewLinkedStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data interface{}) {
	item := &Node{Element: data, Link: s.Head}
	s.Head = item
	s.N++
}
func (s *Stack) Pop() interface{} {
	if s.Head == nil {
		return nil
	}
	item := s.Head
	s.Head = item.Link
	s.N--
	return item.Element
}

func (s *Stack) Size() int {
	return s.N
}

type ArrayStack struct {
	Container []interface{}
	Len       int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{nil, 0}
}

func (s *ArrayStack) Push(data interface{}) {
	s.Container = append(s.Container, data)
	s.Len++
}
func (s *ArrayStack) Pop() interface{} {
	item := s.Container[s.Len-1]
	s.Container = s.Container[:s.Len-1]
	s.Len--
	return item
}
