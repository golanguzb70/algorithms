package main

/*Problem:
A string is good if there are no repeated characters.

Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.

Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

A substring is a contiguous sequence of characters in a string.
*/

/*Solution:
- Two pointer approach:
1. Loop through the string. p1=0
2. If difference between index and p1 is 2, check substring for goodness.
3. Increase count if it is good substring

Checking for goodness logic:
1. Save every character into hash map with its last occurance index as a value
2. If character is not duplicated in last 2 indexes, it is good substring. Increase p1++
3. If not good sub substring, p1=last-occurance of value.


Time complexity: O(n)
Space complexity: O(n)

*/

func CountGoodSubstrings(s string) int {
	mp := map[rune]uint8{}
	p1 := uint8(0)
	resp := 0

	for i, v := range s {
		lastOccurance, ok := mp[v]
		if ok && lastOccurance >= p1 {
			p1 = lastOccurance + 1
			mp[v] = uint8(i)
			continue
		}

		if uint8(i)-p1 < 2 {
			mp[v] = uint8(i)
			continue
		}

		mp[v] = uint8(i)
		resp++
		p1++
	}

	return resp
}
