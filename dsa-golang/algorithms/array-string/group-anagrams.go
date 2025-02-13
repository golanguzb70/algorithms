package main

import (
	"sort"
	"strings"
)

/*Problem:
Given an array of strings strs, group the
anagrams together. You can return the answer in any order.

Anagram: In other words, two strings are considered anagrams if they have the same characters
with the same frequency, but the order of the characters is different.
*/

/*Solution:
Used topics: String, Array, Hash

Algorithm:
1. Loop through the array:
2. Append sorted string element as a key and real value as a value to the map[string][]string{}
3. Loop through the map, append values of map to response.

Time complexity: O(n)
Space complexity: O(n)
*/

func GroupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	response := [][]string{}

	for _, e := range strs {
		sortedStr := strings.Split(e, "")
		sort.Strings(sortedStr)
		mp[strings.Join(sortedStr, "")] = append(mp[strings.Join(sortedStr, "")], e)
	}

	for _, e := range mp {
		response = append(response, e)
	}

	return response
}
