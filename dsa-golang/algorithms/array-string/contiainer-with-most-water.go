package main

import "fmt"

/* Problem:
You are given an integer array height of length n. There are n vertical lines drawn
such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.
Notice that you may not slant the container.
*/

/*Solution:
Use 2 pointer algorithms for this too.
initialization:
p1=first-element, moves only forward
p2=last-element, moves only backward
stops when p1=p2

Which pointer should move: a pointer with minimum height should move.
answer=max(answer, minimumPointerHeight*(p2-p1))

Time compexity: O(n)
Space complexity: O(1)
*/

func MaxArea(height []int) int {
	ans := 0
	p1, p2 := 0, len(height)-1

	for p1 < p2 {
		fmt.Printf("P1: %d P2: %d Capacity: %d\n", height[p1], height[p2], min(height[p1], height[p2])*(p2-p1))
		ans = max(ans, min(height[p1], height[p2])*(p2-p1))

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return ans
}
