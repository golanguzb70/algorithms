package main

/*Problem:
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.
*/

/*Solution:
My thinking process.
- string to be permutation of another string, it should have the same length and frequncies of characters.
- Use fixed sliding window of size s1 in s2.
- Gather frequncies of s1 to hash map key is charakter value is frequncy

Solution1:
- Every time when window slides check if the window is permutation of the s2.

Time complexity: O(n)
Space complexity: O(n)
*/

func CheckInclusion(s1 string, s2 string) bool {
	mp1 := map[byte]int{}
	mp2 := map[byte]int{}
	windowsize := len(s1)

	for i := 0; i < windowsize; i++ {
		mp1[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		mp2[s2[i]]++
		if i-windowsize+1 < 0 {
			continue
		}
		if isPermutation(mp1, mp2) {
			return true
		}

		mp2[s2[i-windowsize+1]]--
	}

	return false
}

func isPermutation(mp1, mp2 map[byte]int) bool {

	for k, v := range mp1 {
		val, ok := mp2[k]
		if !ok || val != v {
			return false
		}
	}

	return true
}
