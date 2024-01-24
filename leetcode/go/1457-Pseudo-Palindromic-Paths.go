package leetcode

/*
Given a binary tree where node values are digits from 1 to 9. A path in the binary tree is said to be
pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.

Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

Lets first remind the hints given by the leetcode.
1. Note that the node values of a path form a palindrome if at most one digit has an odd frequency (parity).
2. Use a Depth First Search (DFS) keeping the frequency (parity) of the digits. Once you are in a leaf node
 check if at most one digit has an odd frequency (parity).



*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func PseudoPalindromicPaths(root *TreeNode) int {
	return dpsMemo(map[int]uint8{}, root)
}

func dps(arr []int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
		mp := map[int]uint32{}
		oddCount := 0
		for _, e := range arr {
			mp[e]++
		}
		for _, v := range mp {
			if v%2 != 0 {
				oddCount++
			}
		}
		if oddCount <= 1 {
			return 1
		}
		return 0
	}

	arr1 := append(arr, root.Val)
	return dps(arr1, root.Left) + dps(arr1, root.Right)
}

func dpsMemo(mp map[int]uint8, root *TreeNode) int {
	oddCount := 0
	mp[root.Val]++
	if root.Left == nil && root.Right == nil {
		for _, v := range mp {
			if v%2 != 0 {
				oddCount++
			}
		}

		if oddCount <= 1 {
			return 1
		}
		return 0
	}
	newMap1 := make(map[int]uint8, len(mp))
	newMap2 := make(map[int]uint8, len(mp))

	for k, v := range mp {
		newMap1[k] = v
		newMap2[k] = v
	}
	res := 0
	if root.Left != nil {
		res += dpsMemo(newMap1, root.Left)
	}
	if root.Right != nil {
		res += dpsMemo(newMap2, root.Right)
	}

	return res
}
