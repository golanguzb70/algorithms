package leetcode

/*
Given an integer array nums, return the length of the longest strictly increasing
subsequence.

My way of thinking:
I am going to use recursive algorithm to solve the problem.
Let's imagine we have input [10,9,2,5,3,7,101,18]

10

*/

func LengthOfLIS(nums []int) int {
	temp := make([]int, len(nums))
	
    for i := range temp {
        temp[i] = 1
		if i > 0 {
			for j := 0; j < i; j++ {
				if nums[j] < nums[i] && temp[j]+1 > temp[i] {
					temp[i]++;
				}
			}
		}
    }

	res := 0
	for _, e := range temp {
		if e > res {
			res = e
		}
	}
	
	return res
}
