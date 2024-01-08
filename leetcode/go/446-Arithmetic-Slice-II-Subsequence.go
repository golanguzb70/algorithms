package leetcode

import (
	"fmt"
	"math/rand"
)

/*
Given an integer array nums, return the number of all the arithmetic subsequences of nums.

A sequence of numbers is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

    For example, [1, 3, 5, 7, 9], [7, 7, 7, 7], and [3, -1, -5, -9] are arithmetic sequences.
    For example, [1, 1, 2, 5, 7] is not an arithmetic sequence.

A subsequence of an array is a sequence that can be formed by removing some elements (possibly none) of the array.

    For example, [2,5,10] is a subsequence of [1,2,1,2,4,1,5,10].

The test cases are generated so that the answer fits in 32-bit integer.

Dry run:
input: [2,4,6,8,10]
[2, 4, 6]
[2, 6, 8]


*/

func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	return NumberOfArithmeticSlicesHelper(nums[1:], nums[0], nums[1]-nums[0], 1, "main")
}

func NumberOfArithmeticSlicesHelper(nums []int, previusNumber, difference, currentLenth int, key string) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]-previusNumber == difference {
			if currentLenth >= 2 {
				res += 1
			}
			res += NumberOfArithmeticSlicesHelper(nums[i+1:], nums[i], difference, currentLenth+1, fmt.Sprintf("%d:%d-index", rand.Int() % 5, i))
		}
	}

	if len(nums) >= 3 {
		res += NumberOfArithmeticSlicesHelper(nums[1:], nums[0], nums[1]-nums[0], 1, "array > 3")
	}

	fmt.Println("Order:", key, res, nums, currentLenth)
	return res
}

/*
dry run the code
[7, 7, 7, 7]
 1  2  3  4
1. f([2, 3, 4], 1, 0, 1)
c = 1
	2.0: f([3, 4], 2, 0, 2)
	c = 2
	res = 1
		3.0: f([4],3, 0, 3)
		c = 3
		res = 1
			4.0: f([], 3, 0, 4)
			res = 0
		res = 1
	res = 2

	2.1: f([4],3, 0, 3)
		c = 3
		res = 1
			4.0: f([], 3, 0, 4)
			res = 0
		res = 1
	res = 3

	2.2: f([3, 4], 2, 0, 1)
		c = 1
		res = 0
			3.0: f([4], 3, 0, 2)
			res = 1
			c = 2
		res = 1
	res = 4
*/
