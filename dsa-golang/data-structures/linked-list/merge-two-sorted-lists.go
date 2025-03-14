package main

/*Problem:
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list
*/

/*Solution:
- Initilize new linked list
- Loop through input the lists
- Insert the value to output list of node with not greater value.
- Return head the reponse
- Loop break point list1 and list2 is nil
*/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		response, prev *ListNode
	)

	for list1 != nil || list2 != nil {
		current := &ListNode{}
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				current.Val = list1.Val
				list1 = list1.Next
			} else {
				current.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 == nil {
			current.Val = list2.Val
			list2 = list2.Next
		} else {
			current.Val = list1.Val
			list1 = list1.Next
		}
		if response == nil {
			response, prev = current, current
		} else {
			prev.Next = current
			prev = current
		}
	}

	return response
}
