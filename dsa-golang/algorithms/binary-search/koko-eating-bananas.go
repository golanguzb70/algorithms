package main

import (
	"math"
	"slices"
)

/*Problem:
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of
bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all
of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
*/

/*Solution:
Let's break questin down to more understandable notes.
1. k is speed of eating bananas.
2. Koko can eat 1 banan in 1 hour, and 1.1 banana in 2 hours. ceil the numer
3. Time taken to finish i'th pile: ceil(number of babanas/speed)

Algorithm:
1. Find the range of ks. 1 to (max banana counts in single pile)
2. Do binary search to calculate to get minimum time taken to eat all of the bananas.
3. Return minimum time.

Time complexity: O(nLog(n))
Space complexity: O(1)
*/

func MinEatingSpeed(piles []int, h int) int {
	low := 1
	high := slices.Max(piles)
	mid, hourCalc := 0, 0
	res := math.MaxInt

	for low <= high {
		mid = (low + high) / 2
		hourCalc = calculateTime(piles, mid)

		if hourCalc <= h {
			if mid <= res {
				res = mid
			}
			high = (low+high)/2 - 1
		} else {
			low = (low+high)/2 + 1
		}
	}

	return res
}

func calculateTime(piles []int, speed int) int {
	res := 0.0

	for _, e := range piles {
		res += math.Ceil(float64(e) / float64(speed))
	}

	return int(res)
}
