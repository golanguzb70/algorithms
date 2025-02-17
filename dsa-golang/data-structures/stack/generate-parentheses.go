package main

import "strings"

/*Problem:
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

/*Solution:
1. Combinations should calculated
2. Each every combination check if value is going right or not.
3. Append to response if valid combination.

Algorithm:
1. Write extra helper recursive function(stack, n)
2. Stop point: n==0, add to response if combination is valid
3. if not valid paratheses return


Time complexity: (2^n)
Space complexity: (2^n)
*/

type StrStack struct {
	items  []string
	items2 []string
}

func (s *StrStack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v, true
}

func (s *StrStack) Push(v string) {
	s.items = append(s.items, v)
	s.items2 = append(s.items2, v)
}

func (s *StrStack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}

	return s.items[len(s.items)-1], true
}

func (s *StrStack) ToString() string {
	return strings.Join(s.items2, "")
}

func generateParenthesisRecursive(n int, currStack *StrStack, response *StrStack) {
	if n == 0 {
		if len(currStack.items) == 0 {
			response.Push(currStack.ToString())
		}
		return
	}

	newStack := &StrStack{
		items:  currStack.items,
		items2: currStack.items2,
	}

	newStack2 := &StrStack{
		items:  currStack.items,
		items2: currStack.items2,
	}

	newStack.Push("(")
	generateParenthesisRecursive(n-1, newStack, response)

	_, ok := newStack2.Pop()
	if !ok {
		return
	}
	newStack2.Push(")")
	newStack2.Pop()
	generateParenthesisRecursive(n-1, newStack2, response)
}

func GenerateParenthesis(n int) []string {
	response := &StrStack{}
	generateParenthesisRecursive(n*2, &StrStack{}, response)

	return response.items
}
