package main

import (
	"fmt"
	"slices"
)

/*Problem:
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.
*/

/*
Solution:

Algorithm:
 1. Loop through the array.
    > Append unique element to extra array
    > Count frequency of elements using map
 2. Sort extra array in decreasing order of frequency of elements.

Time complexity: O(n+nlog(n))
Space complexity: O(n)
*/
func TopKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	newArray := []int{}

	for _, e := range nums {
		_, ok := mp[e]
		if !ok {
			newArray = append(newArray, e)
		}

		mp[e]++
	}

	slices.SortFunc(newArray, func(a int, b int) int {
		if mp[a] > mp[b] {
			return -1
		}
		return 1
	})
	fmt.Println(newArray)

	return newArray[:k]
}
