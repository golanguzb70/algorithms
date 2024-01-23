package leetcode

/*
You are given an array of strings arr. A string s is formed by the concatenation of a subsequence
of arr that has unique characters.

Return the maximum possible length of s.

A subsequence is an array that can be derived from another array by deleting some or no elements
without changing the order of the remaining elements.



*/

func getUniqueChars(st string) int {
	var umap [26]int
	for _, ch := range st {
		umap[ch-'a'] += 1
		if umap[ch-'a'] > 1 {
			return 0
		}
	}
	var count int
	for _, val := range umap {
		if val == 1 {
			count += 1
		}
	}
	return count
}

func dfs(index int, arr []string, st string) int {
	if index == len(arr) {
		return getUniqueChars(st)
	}
	taken := dfs(index+1, arr, st+arr[index])
	notTaken := dfs(index+1, arr, st)
	return max(taken, notTaken)
}

func maxLength(arr []string) int {
	return dfs(0, arr, "")
}
