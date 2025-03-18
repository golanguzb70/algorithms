package main

/*Problem:
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/

/*Solution:
1. DFS
2. Add the thenth of last left and right leaves.
3. Get max length if length of right or left leaves are longer.
*/

func DiameterOfBinaryTree(root *TreeNode) int {

	_, diametr := DFS(root)

	return diametr
}

func DFS(root *TreeNode) (int, int) { // (deepestL, diametr)
	if root == nil {
		return 0, 0
	}

	leftL, leftD := DFS(root.Left)
	rightL, rightD := DFS(root.Right)

	diametr := leftL + rightL

	return max(leftL+1, rightL+1), max(leftD, max(rightD, diametr))
}
