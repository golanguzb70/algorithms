package main

/*Problem:
Given two integer arrays nums1 and nums2, return an array of their
intersection. Each element in the result must be unique and you may return the result in any order.
*/

/*Solution:
1. Map that values of nums1 to map.
2. Iterate through the nums2, add to response if value exists in the map.

*/

func Intersection(nums1 []int, nums2 []int) []int {
	responseMp := map[int]struct{}{}
	mp := map[int]struct{}{}

	for _, e := range nums1 {
		mp[e] = struct{}{}
	}

	for _, e := range nums2 {
		if _, ok := mp[e]; ok {
			responseMp[e] = struct{}{}
		}
	}

	resp := []int{}
	for k := range responseMp {
		resp = append(resp, k)
	}

	return resp
}
