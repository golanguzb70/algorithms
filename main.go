package main

import (
	"fmt"

	leetcode "github.com/golanguzb70/algorithms/leetcode/go"
)

func main() {
	nodes := &leetcode.ListNode{
		Val: 1,
		Next: &leetcode.ListNode{
			Val: 2,
			Next: &leetcode.ListNode{
				Val: 3,
				Next: &leetcode.ListNode{
					Val: 4,
				},
			},
		},
	}

	response := leetcode.SwapPairs(nodes)

	for response != nil {
		fmt.Println(response.Val)
		response = response.Next
	}
}
