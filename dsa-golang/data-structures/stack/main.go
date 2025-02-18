package main

import "fmt"

/*
	Stack is data structure that is LIFO:
	Push -> Adding value. -> O(n) (Worst case when resizing is needed)
	Pop -> Reading and removing value from end of the stack -> O(1)
	Peek -> Reading from the end of the stack. ->  O(1)
*/

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
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(DailyTemperatures(input))
}
