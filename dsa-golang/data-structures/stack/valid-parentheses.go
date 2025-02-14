package main

/*Problem:
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

/*Solution:

1. Map pairs of all brackets.
2. Loop through the input
3. If element is open bracket, push to stack
4. If element is close bracket, pop stack and return false if it is not pair of the bracket.
5. At last return  if stack size is zero


Time complexity: O(n)
Space complexity: O(n)
*/

type BracketsStack struct {
	items []byte
}

func (s *BracketsStack) Pop() (byte, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v, true
}

func (s *BracketsStack) Push(v byte) {
	s.items = append(s.items, v)
}

func IsValid(s string) bool {
	mp := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stackDs := BracketsStack{}

	for i := 0; i < len(s); i++ {
		v, ok := mp[s[i]]
		if !ok {
			stackDs.Push(s[i])
		} else {
			stackValue, found := stackDs.Pop()
			if !found || v != stackValue {
				return false
			}
		}
	}

	return len(stackDs.items) == 0
}
