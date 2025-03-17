package main

/*Problem:
Given the root of a binary tree, invert the tree,  and return its root.
*/

/*Solution:
What is inverting?
- inverting is swaping leafes of binary tree.

Algorithm:
- DFS
- Swap the leaves
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	root.Right = InvertTree(root.Right)
	root.Left = InvertTree(root.Left)

	root.Left, root.Right = root.Right, root.Left

	return root
}
