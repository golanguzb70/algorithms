package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v, true
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())

	s.Push(10)

	fmt.Println(s.items)
}
