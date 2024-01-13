package leetcode

/*
You are given two strings of the same length s and t. In one step you can choose any character of t and replace it with another character.

Return the minimum number of steps to make t an anagram of s.

An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.


Algorithm:
1. Count frequencies of characters in each string
2. Add the difference to the result.

*/

func MinSteps(s string, t string) (res int) {
	dict := [26]int{}
	for k, v := range s {
		dict[v-'a']++
		dict[t[k]-'a']--
	}

	for _, v := range dict {
		if v > 0 {
			res += v
		}
	}
	return res
}
