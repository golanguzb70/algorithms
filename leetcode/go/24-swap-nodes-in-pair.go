package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	res := SwapPairsHelper(head)
	return res
}

/*
 1. Not to lose pointer to the head of the linked list use temp to implement swapping
 2. Write a recursive helper function that swaps the nodes.
     ||
    \  /
     \/

Swapping logic -> if node.next is not null ->
temp := head
head = head.next
temp.next = temp.next.next
head.next = temp
head.next.next = recursivefunction(head.next.next)

time complexity -> O(n/2)
space complexity -> O(n*(n-2)*(n-4))
*/
func SwapPairsHelper(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	temp := head
	head = head.Next
	temp.Next = temp.Next.Next
	head.Next = temp

	head.Next.Next = SwapPairsHelper(head.Next.Next)
	return head
}
