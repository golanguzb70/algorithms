package main

import "strings"

/*Problem:
You are given an absolute path for a Unix-style file system, which always begins with a slash '/'. Your task is to transform this absolute path into its simplified canonical path.

The rules of a Unix-style file system are as follows:

A single period '.' represents the current directory.
A double period '..' represents the previous/parent directory.
Multiple consecutive slashes such as '//' and '///' are treated as a single slash '/'.
Any sequence of periods that does not match the rules above should be treated as a valid directory or file name. For example, '...' and '....' are valid directory or file names.
The simplified canonical path should follow these rules:

The path must start with a single slash '/'.
Directories within the path must be separated by exactly one slash '/'.
The path must not end with a slash '/', unless it is the root directory.
The path must not have any single or double periods ('.' and '..') used to denote current or parent directories.
Return the simplified canonical path.
*/

type Stack struct {
	items []string
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false // Stack is empty
	}
	top := s.items[len(s.items)-1]     // Get the top element
	s.items = s.items[:len(s.items)-1] // Remove it
	return top, true
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false // Stack is empty
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

func SimplifyPath(path string) string {
	stack := Stack{}

	for _, e := range strings.Split(path, "/") {
		if e == "" || e == "." {
			continue
		}

		if e == ".." {
			stack.Pop()
		} else {
			stack.Push(e)
		}
	}

	resp := strings.Join(stack.items, "/")
	return "/" + resp
}
