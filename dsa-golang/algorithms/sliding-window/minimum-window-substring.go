package main

/*Problem:
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character
 in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.
*/

/*Solution:
- Use dynamic sliding window algorithm using two pointer to mean start and end of the window.
- Expand right pointer untill all the characters are found in the window.
- After expand left pointer untill only 1 character frequency left.
- During it check if the substring is shortest.
*/

func MinWindow(s string, t string) string {
	response := ""
	p1 := 0
	existingChars, notFoundChars := map[byte]struct{}{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		existingChars[t[i]] = struct{}{}
		notFoundChars[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		_, ok := existingChars[s[i]]
		if ok {
			notFoundChars[s[i]]--
		}

		for IsAllValZero(notFoundChars) {

			if len(response) > len(s[p1:i+1]) || response == "" {
				response = s[p1 : i+1]
			}

			_, ok := existingChars[s[p1]]
			if ok {
				notFoundChars[s[p1]]++
			}
			p1++
		}

	}

	return response
}

func IsAllValZero(mp map[byte]int) bool {
	for _, v := range mp {
		if v > 0 {
			return false
		}
	}
	return true
}
