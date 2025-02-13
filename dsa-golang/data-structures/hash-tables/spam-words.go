package main

/*Problem:
You are given an array of strings message and an array of strings bannedWords.

An array of words is considered spam if there are at least two words in it that exactly match any word in bannedWords.

Return true if the array message is spam, and false otherwise.
*/

/*
Solution:
0. Return true if len(bannedWords) < 2
1. Hash banned words
2. Iterate throught the message.
3. Icrement spam count if word is banned.
4. If spam count is equals to 2 return false.

Time complexity: O(n)
Space complexity: O(n)
*/

func ReportSpam(message []string, bannedWords []string) bool {
	mp := map[string]struct{}{}
	for _, e := range bannedWords {
		mp[e] = struct{}{}
	}

	spamCount := 0

	for _, e := range message {
		_, ok := mp[e]
		if ok {
			spamCount++
		}
		if spamCount == 2 {
			return true
		}
	}

	return false
}
