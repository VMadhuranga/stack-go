package main

import "fmt"

// Doubly linked list implementation

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

type stack struct {
	doublyLinkedList
}

func (s *stack) push(data int) {
	node := &node{
		data: data,
		prev: s.tail,
		next: nil,
	}

	if s.head == nil {
		s.head = node
	} else {
		s.tail.next = node
	}

	s.tail = node
	s.size++
}

func (s *stack) pop() {
	if s.head == nil {
		fmt.Println("Empty stack")
	} else if s.size == 1 {
		s.head = nil
		s.tail = nil
		s.size--
	} else {
		s.tail = s.tail.prev
		s.tail.next = nil
		s.size--
	}
}

func (s stack) printTop() {
	if s.head == nil {
		fmt.Println("Empty stack")
	} else {
		fmt.Println(s.tail.data)
	}
}

func (s stack) printSize() {
	fmt.Println(s.size)
}

func main() {
	s := stack{}        // nil
	s.printTop()        // Empty stack
	s.printSize()       // 0
	s.pop()             // Empty stack
	fmt.Println(s.head) // nil
	fmt.Println(s.tail) // nil

	s.push(1)           // 1 -> nil
	s.printTop()        // 1
	s.printSize()       // 1
	fmt.Println(s.head) // 1
	fmt.Println(s.tail) // 1

	s.pop()             // nil
	s.printTop()        // Empty stack
	s.printSize()       // 0
	fmt.Println(s.head) // nil
	fmt.Println(s.tail) // nil

	s.push(3)           // 3 -> nil
	s.push(5)           // 3 -> 5 -> nil
	s.push(7)           // 3 -> 5 -> 7 -> nil
	s.printTop()        // 7
	s.printSize()       // 3
	fmt.Println(s.head) // 3
	fmt.Println(s.tail) // 7

	s.pop()             // 3 -> 5 -> nil
	s.printTop()        // 5
	s.printSize()       // 2
	fmt.Println(s.head) // 3
	fmt.Println(s.tail) // 5
}

// Slices implementation

// type stack struct {
// 	data []int
// }

// func (s *stack) push(data int) {
// 	s.data = append(s.data, data)
// }

// func (s *stack) pop() {
// 	if len(s.data) <= 0 {
// 		fmt.Println("Empty stack")
// 	} else {
// 		s.data = s.data[:len(s.data)-1]
// 	}
// }

// func (s stack) printTop() {
// 	if len(s.data) <= 0 {
// 		fmt.Println("Empty stack")
// 	} else {
// 		fmt.Println(s.data[len(s.data)-1])
// 	}
// }

// func (s *stack) printSize() {
// 	fmt.Println(len(s.data))
// }

// func main() {
// 	stack := stack{}  // []
// 	stack.printSize() // 0
// 	stack.printTop()  // Empty stack
// 	stack.pop()       // Empty stack

// 	stack.push(1)     // [1]
// 	stack.printSize() // 1
// 	stack.printTop()  // 1

// 	stack.pop()       // []
// 	stack.printSize() // 0
// 	stack.printTop()  // Empty stack

// 	stack.push(3)     // [3]
// 	stack.push(5)     // [3, 5]
// 	stack.push(7)     // [3, 5, 7]
// 	stack.printSize() // 3
// 	stack.printTop()  // 7

// 	stack.pop()       // [3, 5]
// 	stack.printSize() // 2
// 	stack.printTop()  // 5
// }
