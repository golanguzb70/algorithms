package main

/*Problem:
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

/*Solution:
- Maintain 2 pointers, p1=cn p2 is cn+n
- When p2 becomes nil remove p1
*/

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	cn := head
	var prev *ListNode
	lenth := 0

	for cn.Next != nil {
		lenth++
		cn = cn.Next
	}

	cn = head

	for i := 0; i < lenth; i++ {

		if lenth-n == i {
			if prev == nil {
				prev = cn.Next
				head = prev
			} else {
				prev = cn.Next
			}

		} else {
			if prev == nil {
				prev = cn
				head = prev
			} else {
				prev = cn
			}
			cn = cn.Next
		}

	}

	return head
}
