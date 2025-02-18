package main

/*Problem:
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after
the ith day to get a warmer temperature. If there is no future day for which this is possible,
keep answer[i] == 0 instead.
*/

/*Solution:

Monothonic stack:
 - Saving elements indexes that are only increasing/Decreasing order.

Algorithm:
1. Loop through the array.
2. Chech if any smaller values exist in stack
 > If exists, put the difference to new array pop from the stack
 > If not exists push the index of the element to the stack
3. Return the response


Time complexity: O(n)
Space complexity: O(n)
*/

type MonStack struct {
	items []int
}

func (s *MonStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v, true
}

func (s *MonStack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *MonStack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func DailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))
	stack := MonStack{}

	for i, e := range temperatures {
		if v, ok := stack.Peek(); !ok || temperatures[v] >= e {
			stack.Push(i)
			continue
		}

		for {
			v, ok := stack.Peek()
			if !ok || temperatures[v] >= e {
				stack.Push(i)
				break
			}

			response[v] = i - v

			stack.Pop()
		}
	}

	return response
}
