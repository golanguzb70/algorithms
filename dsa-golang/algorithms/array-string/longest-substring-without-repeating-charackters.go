package main

import "fmt"

/*
Given a string s, find the length of the longest substring without repeating characters.

I am going to use 2 pointer algorithm to solve:
input: abcabcbb

start point & checks:  p1 <= p2(is index in iteration)
substring is elements between [p1, p2]
Need to memorize if the character is used or not in last sub string: map[char]index
When you reach an element that was used in last substring:
 1. Label as maximum substring if lenth is longer that previus
 2. Move p1 to an element that is after last use case of the charackter.
When to stop: p2 = len(input)

Time complexity: O(n)
Space complexity: O(n)

*/

/*
abcabcbb

loop1:
	p1=0
	p2=0
	substr=a
	max=1
loop2:
	p1=0
	p2=1
	substr=ab
	max=2
loop3:
	p1=0
	p2=2
	substr=abc
	max=3
loop4:
	p1=1
	p2=3
	substr=bca
	max=3
loop5:
	p1=2
	p2=4
	substr=cab
	max=3
...
*/



func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	p1 := 0
	hashMap := map[rune]int{}

	for i, e := range s {
		index, ok := hashMap[e]
		if !ok || (ok && index < p1) {
			maxLen = max(maxLen, i-p1+1)
			fmt.Println(maxLen)
			hashMap[e] = i
			continue
		}

		p1 = index + 1
		hashMap[e] = i
	}

	return maxLen
}
