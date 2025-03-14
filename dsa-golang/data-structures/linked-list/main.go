package main

import "fmt"

func main() {
	arr := []int{1, 2, 4}
	arr1 := []int{1, 3, 4}

	list1 := &LinkedList{}
	for _, e := range arr {
		list1.AddEnd(e)
	}

	list2 := &LinkedList{}
	for _, e := range arr1 {
		list2.AddEnd(e)
	}

	result := MergeTwoLists(list1.Head, list2.Head)

	output := LinkedList{Head: result}
	output.PrintList()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) AddEnd(value int) {
	nn := &ListNode{value, nil}

	if l.Head == nil {
		l.Head = nn
	} else {
		cn := l.Head

		for cn.Next != nil {
			cn = cn.Next
		}
		cn.Next = nn
	}
}

func (l *LinkedList) FindAndInsertAfter(target, value int) {
	nn := &ListNode{value, nil}

	if l.Head == nil {
		l.Head = nn
	} else {
		cn := l.Head

		for cn.Next != nil {
			if cn.Val == target {
				break
			}
			cn = cn.Next
		}

		nn.Next = cn.Next
		cn.Next = nn
	}
}

func (l *LinkedList) RemoveNode(value int) {
	if l.Head == nil {
		return
	}

	cn := l.Head

	if cn.Val == value {
		l.Head = l.Head.Next
	}

	for cn.Next != nil && cn.Next.Val != value {
		cn = cn.Next
	}

	if cn.Next == nil {
		return
	}

	cn.Next = cn.Next.Next
}

func (l *LinkedList) RemoveAllAccurances(value int) {
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Val == value {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		return
	}
	cn := l.Head

	for cn.Next != nil {
		if cn.Next.Val == value {
			cn.Next = cn.Next.Next
		} else {
			cn = cn.Next
		}
	}
}

func (l *LinkedList) Reverse() {
	var prev, next *ListNode

	cn := l.Head

	for cn != nil {
		next = cn.Next
		cn.Next = prev
		prev = cn
		cn = next
	}

	l.Head = prev
}

func (l *LinkedList) PrintList() {
	cn := l.Head

	for cn != nil {
		fmt.Printf("Node address: %p\tValue: %d\tNext Node address: %p\n", cn, cn.Val, cn.Next)
		cn = cn.Next
	}
}
