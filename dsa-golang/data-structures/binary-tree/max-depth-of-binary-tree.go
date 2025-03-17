package main

/*Problem:
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node
*/

/*Solution:
Brute force solution:
- DFS
- Add max depth
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}
