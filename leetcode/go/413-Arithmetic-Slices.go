package leetcode

/*
An integer array is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

    For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.

Given an integer array nums, return the number of arithmetic subarrays of nums.

A subarray is a contiguous subsequence of the array.

Algorithm:
Let's imagine I have an array [1, 2, 3, 4]
My goal is to get arithmetic subarrays of the array.
lenth of subarray should be at  least 3
I can get:
1. [1, 2, 3] d=1
2. [2, 3, 4] d=1
3. [1, 2, 3, 4] d=1

So the result is 3

First brute force solution is recursive solution

f(arr, last_number, difference, current_length) {
	if arr.lenth == 0 {
		return current_lenth >= 3;
	}


	if abs(arr[0] - last_number) == difference  &&  current_length >= 3 {
		res++
	} else if abs(arr[0] - last_number) != difference  {
		return 0;
	}

	if arr.lenth >= 3 {
		res += f(arr[1:], arr[0], abs(arr[0]-last_number), 1)
	}
	return res + f(arr[1:], arr[0], difference, current_lenth+1)
}

dry run:
arr := [1, 2, 3, 4]: f(arr[1:], arr[0], abs(arr[0]-arr[1]), 1)
res = 0
1:
arr=[2, 3, 4]
lenth == 3
curr = 1
res = 0;
	1.1: arr[3, 4] expect(2, 3, 4)
	lenth == 2
	curr = 1
	res = 0
		1.2: arr[4]
		lenth = 1
		curr = 2
		res = 0
			1.3: arr[]
			lenth = 0;
			curr = 3
			res = 1
		res = 1
	1.4: arr[3, 4]
	lenth 2
	curr = 2;
	res = 0
		1.5: arr[4]
		lenth = 1
		curr = 3
		res = 1
			1.6: arr[]
			lenth=0
			current=4
			res =1
		res = 2
res= 3
*/

func NumberOfArithmeticSlicesMedium(nums []int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}

	res := 0
	for i := 0; i <= length-3; i++ {
		diference := nums[i+1] - nums[i]
		count := 1

		for j := i + 1; j < length; j++ {
			if nums[j]-nums[j-1] != diference {
				break
			}
			count++
			if count >= 3 {
				res++
			}
		}
		
	}

	return res
}
