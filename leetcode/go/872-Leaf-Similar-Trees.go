package leetcode

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := DFS(root1)
	arr2 := DFS(root2)
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func DFS(root *TreeNode) []int {
	res := []int{}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	if root.Left != nil {
		res = append(res, DFS(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, DFS(root.Right)...)
	}

	return res
}
