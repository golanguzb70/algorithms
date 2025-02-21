package main

/*Problem:
Given n non-negative integers representing an elevation map where the width of each bar is 1,
compute how much water it can trap after raining.
*/

/*Solution:
Example input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

Algorithm is finding next and previues greater elements.
Init:
previues=0 index
next=0 index
trapped_water=(nums[previues] * (next-previues) - black_parts)

1. Loop throught the input
2. If element is next greater element calculate water trapped between them.
3. Move previues=pointer to next greater element.

Time complexity: O(n)
Space complexity: O(1)

Example run:
loop1:
previues=0
next=?
loop2:
is_next_greater_element=true
trapped_water=0
previues=1
next=1
loop3:
previues=1
is_next_greater_element=false
loop4:
previues=1
is_next_greater_element=false
loop3:
previues=1
is_next_greater_element=false
loop3:
previues=1
is_next_greater_element=false
*/

func Trap(height []int) int {
	l := len(height)
	maxLeft := height[0]
	maxRight := height[l-1]
	left, right := 1, l-1
	res := 0
	for left <= right {
		if maxLeft <= maxRight {
			if maxLeft-height[left] > 0 {
				res += maxLeft - height[left]
			}

			if height[left] > maxLeft {
				maxLeft = height[left]
			}
			left++
		} else {
			if maxRight-height[right] > 0 {
				res += maxRight - height[right]
			}

			if height[right] > maxRight {
				maxRight = height[right]
			}
			right--
		}
	}
	return res
}
