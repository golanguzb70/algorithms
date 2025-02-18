package main

import (
	"strings"
)

/*Problem:
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing
all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

/*Solution:
Algorithm: 2 pointers
1. initilize pointers. p1=0, p2=len(input)-1
2. decrease p2 if p2 is not letter.
3. increase p1 if p1 is not letter.
4. return false if p1.value!=p2.value
5. Decrease p1 and p2.
6. Stop point p1>p2

Time comlexity: O(n)
Space complexity: O(1)

*/

func isAlphaNumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}

	return false
}

func IsPalindrome(s string) bool {

	p1, p2 := 0, len(s)-1
	hasChanged := false
	for p1 < p2 {
		hasChanged = false
		if !isAlphaNumeric(s[p1]) {
			p1++
			hasChanged = true
		}

		if !isAlphaNumeric(s[p2]) {
			p2--
			hasChanged = true
		}

		if hasChanged {
			continue
		}

		if !strings.EqualFold(strings.ToLower(string(s[p1])), strings.ToLower(string(s[p2]))) {
			return false
		}
		p1++
		p2--
	}

	return true
}
