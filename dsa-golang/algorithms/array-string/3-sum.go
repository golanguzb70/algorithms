package main

import (
	"fmt"
	"slices"
)

/*Problem:
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that
	i != j, i != k, and j != k,
and
	nums[i] + nums[j] + nums[k] == 0

Notice that the solution set must not contain duplicate triplets.
*/

/*Solution:
Just facts:
1. Need three numbers x, y and z
2. If we get any number as x, we can find y and z using 2 sum problem.
3. So the we don't use the same element twice use hash map.


Algorithm:
1. Sort input array.
2. Write 2 sum function func(numbers, target) that returns all the possible unique combinations
that sums up to target.
 > How to know if combination is unique?
  - Search the next number after the current element because input is sorted.
3. Loop through the input array
4. Get 2 sums after the element.
5. If there are results from 2sum function append to the response
6. Return the response


Time complexity: O(n^2)
Space complexity: O(n)
*/

func TwoSumHelper(numbers []int, target int, startNumber int, mp map[string]struct{}) [][]int {
	response := [][]int{}

	for i := 0; i < len(numbers)-1; i++ {
		index, found := slices.BinarySearch(numbers[i+1:], target-numbers[i])
		if found {
			index += i + 1

			key := fmt.Sprintf("%d%d%d", startNumber, numbers[i], numbers[index])
			_, ok := mp[key]
			if !ok {
				response = append(response, []int{numbers[i], numbers[index]})
				mp[key] = struct{}{}
			}
		}
	}

	return response
}

func ThreeSum(nums []int) [][]int {
	response := [][]int{}
	mp := map[string]struct{}{}
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		twosums := TwoSumHelper(nums[i+1:], 0-nums[i], nums[i], mp)
		if len(twosums) > 0 {
			for _, e := range twosums {
				temp := append([]int{nums[i]}, e...)

				response = append(response, temp)
			}
		}
	}

	return response
}
