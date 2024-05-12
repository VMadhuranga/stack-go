package main

type node struct {
	data any
	next *node
}

type stack struct {
	head *node
	size int
}

func (s *stack) push(data any) {
	n := &node{
		data: data,
		next: nil,
	}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}

	s.size++
}

func (s *stack) pop() any {
	n := &node{}

	if s.head != nil {
		n = s.head
		s.head = s.head.next
		s.size--
	}

	return n.data
}

func (s *stack) peek() any {
	return s.head.data
}

func (s *stack) getSize() int {
	return s.size
}
