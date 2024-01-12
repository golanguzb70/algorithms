package leetcode

/*
You are given the root of a binary tree with unique values, and an integer start. At minute 0, an infection starts from the node with value start.

Each minute, a node becomes infected if:

	The node is currently uninfected.
	The node is adjacent to an infected node.

Return the number of minutes needed for the entire tree to be infected.
*/
var ans = 0

func AmountOfTime(root *TreeNode, start int) int {
	ans = 0
	DepthFirstSearch(root, start)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DepthFirstSearch(root *TreeNode, start int) (int, bool) {
	if root == nil {
		return 0, false
	}

	left, foundInLeft := DepthFirstSearch(root.Left, start)
	right, foundInRight := DepthFirstSearch(root.Right, start)

	if root.Val == start {
		ans = max(ans, max(left, right))
		return 1, true
	}

	if foundInLeft {
		ans = max(ans, right+left)
		return left + 1, true

	} else if foundInRight {
		ans = max(ans, right+left)
		return right + 1, true

	} else {
		return max(left, right) + 1, false
	}
}

