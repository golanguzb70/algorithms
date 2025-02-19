package main

/*Problem:
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
find two numbers such that they add up to a specific target number. Let these two numbers be
numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array
[index1, index2] of length 2.

The tests are generated such that there is exactly one solution.
You may not use the same element twice.

Your solution must use only constant extra space.
*/

/*Solution:
Algorithm:
1. Loop through the input
2. Map the array
3. Loop through the array again
4. If target-element key exists in the array return the indexes adding 1.

Time complexity: O(n)
Space complexity: O(n)

*/

func TwoSum(numbers []int, target int) []int {
	mp := map[int]int{}

	for i, e := range numbers {
		mp[e] = i
	}

	for i, e := range numbers {
		val, exists := mp[target-e]

		if exists {
			return []int{i + 1, val + 1}
		}
	}

	return []int{}
}
