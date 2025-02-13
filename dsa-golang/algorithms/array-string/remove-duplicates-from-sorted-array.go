package main

import "fmt"

/*Problem:
Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique
element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be
placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates,
then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first
k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1)
extra memory.
*/

/*Solution:
Notes:
1. Sorted in non-decreasing order
2. Each unique element should occur at most 2 times
3. Return number of slots that contains values.

Algorithm: 2 pointer, p1 is starting index of unique number. p2 is index in the loop
1. Loop through the array.
2. If element equals to last element, continue
3. If element is not equals to last element or p2 is the last element:
  > Add p2-p1 to k and continue, if p2-p1 is less than or equals 2.
  > Delete extra elements and add 2 to k, if p2-p1 is greater than 2.
4. Return k
Time complexity: O(n)
Space complexity: O(1)

Element deleting algorithm:
Input: StartingIndex, SwapCount, Nums, AlreadyDeletedCount
Algorithm:
1. Loop through the array, i=StartingIndex, i< len(Nums)-AlreadyDeletedCount-SwapCount;i++
2. Swap array value: nums[i]=nums[i+SwapCount]

Time complexity: O(n)
Space complexity: O(1)
*/

func RemoveDuplicates(nums []int) int {
	deletedCount := 0
	p2 := 0
	k := 0

	fmt.Println(nums)
	if len(nums) <= 2 {
		return len(nums)
	}

	for i := 0; i < len(nums)-deletedCount; {
		if nums[p2] != nums[i] {
			if i-p2 <= 2 {
				k += i - p2
				p2 = i
			} else {
				k += 2
				SwapArray(p2+2, i-p2-2, deletedCount, nums)
				deletedCount += i - p2 - 2
				i = p2 + 2
				p2 = i
			}
		}
		i++

		if i == len(nums)-deletedCount {
			fmt.Println("here here")

			if i-p2 <= 2 {
				k += i - p2
			} else {
				k += 2
				SwapArray(p2+2, i-p2-2, deletedCount, nums)
				deletedCount += i - p2 - 2
			}
			break
		}

	}

	return k
}

func SwapArray(startingIndex, swapSize, deletedCount int, nums []int) {
	for i := startingIndex; i < len(nums)-deletedCount-swapSize; i++ {
		nums[i] = nums[i+swapSize]
	}
}
