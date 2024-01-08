package leetcode

import "slices"

/*
Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.

Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.

Logic here is that:


*/

func FindContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var pointerG, pointerS, result int

	for pointerG < len(g) && pointerS < len(s) {
		if g[pointerG] <= s[pointerS] {
			result++
			pointerG++
		}
		pointerS++
	}

	return result
}
