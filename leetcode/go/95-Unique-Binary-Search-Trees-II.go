package leetcode

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTree := helper(start, i-1)
		rightTree := helper(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}

func PrintNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	PrintNode(node.Left)
	PrintNode(node.Right)
}
