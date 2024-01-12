package leetcode

import "strings"

/*
You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.

Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and lowercase letters.

Return true if a and b are alike. Otherwise, return false.
*/

func HalvesAreAlike(s string) bool {
	var res int32

	for i := 0; i < len(s)/2; i++ {
		if strings.Contains("aeiouAEIOU", string(s[i])) {
			res++
		}
		if strings.Contains("aeiouAEIOU", string(s[len(s)-1-i])) {
			res--
		}
	}

	return res == 0
}
