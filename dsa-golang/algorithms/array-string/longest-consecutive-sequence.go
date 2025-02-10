package main

import (
	"fmt"
	"sort"
)

/*Problem:
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/

/*Solution:
1. Just sort the array in ascending order.
2. Iterate through the array untill finding the difference.

Time complexity: O(n+nlog(n))
Space complexity: O(1)
*/

func LongestConsecutive(nums []int) int {
	ans := 0

	sort.Ints(nums)
	p1 := 0
	fmt.Println(nums)
	if len(nums) == 1 {
		return 1
	}

	duplicatesCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			duplicatesCount++
		}

		if nums[i] != nums[i-1]+1 && nums[i] != nums[i-1] {
			if i-p1-duplicatesCount > ans {
				ans = i - p1 - duplicatesCount
			}
			duplicatesCount = 0
			p1 = i
		} else if i == len(nums)-1 {
			fmt.Println(i, p1)
			if i-p1+1-duplicatesCount > ans {
				ans = i - p1 + 1 - duplicatesCount
			}
		}
	}

	return ans
}
