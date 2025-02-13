package main

/*Problem:
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all
the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

/*Solution:

Algorithm:
1. Open new array with the same size as an input.
2. Loop through the input from start to end, put product of all left side element to the index of the new array.
3. Loop through the input from end to start, put product of all right side element to the index of the new array.

Time complexity: O(n)
Space complexity: O(1)
*/

func ProductExceptSelf(nums []int) []int {
	response := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {
		response[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		response[i] *= product
		product *= nums[i]
	}

	return response
}
