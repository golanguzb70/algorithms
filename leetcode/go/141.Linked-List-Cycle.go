package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

var pos = 1

func hasCycle(head *ListNode) bool {
	var val, index = 0, 0

	for head != nil {
		if index == pos {
			val = head.Val
		} else if index > pos {
			if val == head.Val {
				return true
			}
		}
	}

	return false
}
