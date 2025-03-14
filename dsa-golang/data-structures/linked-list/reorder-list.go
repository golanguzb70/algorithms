package main

/*Problem:
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/

/*Solution:
- Get last half of the list to array
- Merge them together.
*/

func ReorderList(head *ListNode) {
	arr := []*ListNode{}
	temp := head
	for temp != nil {
		arr = append(arr, temp)
		temp = temp.Next
	}

	cn := head
	head = cn

	for i := len(arr) - 1; i >= (len(arr)-1)/2; i-- {
		nn := arr[i]
		nn.Next = cn.Next
		cn.Next = nn
		cn = cn.Next.Next
	}

	cn.Next = nil
}
