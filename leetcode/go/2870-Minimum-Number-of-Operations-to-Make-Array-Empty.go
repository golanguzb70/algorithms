package leetcode

/*
You are given a 0-indexed array nums consisting of positive integers.

There are two types of operations that you can apply on the array any number of times:

    Choose two elements with equal values and delete them from the array.
    Choose three elements with equal values and delete them from the array.

Return the minimum number of operations required to make the array empty, or -1 if it is not possible.

Algorith:
1. store all numbers to map[int]int key=number value=occurance
2. Iterate through the map

Think a case
1. occurance = 1 (return -1)
2. occurance % 3 = 0 (add occurance/3)
3. occurance % 3 = 1 ()
4. occurance % 3 = 2 ()


*/

func MinOperations(nums []int) int {
	res := 0
	mp := map[int]int{}

	for _, e := range nums {
		mp[e]++
	}

	for _, v := range mp {
		if v == 1 {
			return -1
		} else if v%3 == 0 {
			res += v / 3
		} else if v%3 == 1 {
			res += 2 + (v-4)/3
		} else if v%3 == 2 {
			res += (v / 3) + 1
		}
	}
	return res
}
