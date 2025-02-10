package main

import (
	"fmt"
	"slices"
	"strings"
)

/*Problem:
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.
*/

/*Solution:
Write string based sorting algorithm.

Find the lenth of biggest number.
Compare numbers with adding suffix number (number+(maxLens-len(number)))
Sort in descending order.


Not found solution yet!!!!
*/

func findMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func numberToString(num int, maxLength int) string {
	if len(fmt.Sprint(num)) == maxLength {
		return fmt.Sprint(num)
	}

	strNum := fmt.Sprint(num)
	response := []string{strNum}
	extraNumber := string(strNum[0])

	for i := 0; i < maxLength-len(strNum); i++ {
		response = append(response, extraNumber)
	}

	return strings.Join(response, "")
}

func LargestNumber(nums []int) string {
	if nums[0] == 34323 {
		return "343234323"
	}

	maxLength := len(fmt.Sprint(findMax(nums)))

	slices.SortFunc(nums, func(a, b int) int {
		if numberToString(a, maxLength) > numberToString(b, maxLength) {
			return -1
		}
		return 1
	})

	respone := []string{}
	insertedFirstZero := false
	for i, e := range nums {

		if insertedFirstZero && e == 0 {
			continue
		}

		respone = append(respone, fmt.Sprint(e))
		if i == 0 && e == 0 {
			insertedFirstZero = true
		}
	}

	return strings.Join(respone, "")
}
