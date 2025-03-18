package main

/*Problem:
Given a binary tree, determine if it is height-balanced.

Height-balanced
-the depth of the two subtrees
of every node never differs by more than one.
*/

/*Solution:
Algorithm:
- DFS
- Check if it is balanced.
- Return false if if only 1 folse exists.
*/

func IsBalanced(root *TreeNode) bool {
	_, res := IsBalancedDFS(root)

	return res
}

func IsBalancedDFS(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	rightLenth, balanced := IsBalancedDFS(root.Right)
	if !balanced {
		return rightLenth, balanced
	}

	leftLenth, balanced := IsBalancedDFS(root.Left)
	if !balanced {
		return leftLenth, balanced
	}

	if leftLenth-rightLenth > 1 || leftLenth-rightLenth < -1 {
		return max(rightLenth, leftLenth) + 1, false
	}

	return max(rightLenth, leftLenth) + 1, true
}
