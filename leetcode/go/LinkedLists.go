package leetcode

import "fmt"

type Nodee struct {
	Value int
	Next  *Nodee
}

type LinkedList struct {
	Head *Nodee
}

func (l *LinkedList) AddEnd(value int) {
	nn := &Nodee{value, nil}

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
	nn := &Nodee{value, nil}

	if l.Head == nil {
		l.Head = nn
	} else {
		cn := l.Head

		for cn.Next != nil {
			if cn.Value == target {
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

	if cn.Value == value {
		l.Head = l.Head.Next
	}

	for cn.Next != nil && cn.Next.Value != value {
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

	for l.Head != nil && l.Head.Value == value {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		return
	}
	cn := l.Head

	for cn.Next != nil {
		if cn.Next.Value == value {
			cn.Next = cn.Next.Next
		} else {
			cn = cn.Next
		}
	}
}

func (l *LinkedList) Reverse() {
	var prev, next *Nodee

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
		fmt.Printf("Node address: %p\tValue: %d\tNext Node address: %p\n", cn, cn.Value, cn.Next)
		cn = cn.Next
	}
}

func RunLinkedList() {

	list := LinkedList{}
	
	list.AddEnd(1)
	list.AddEnd(2)
	list.AddEnd(3)
	list.AddEnd(4)

	list.FindAndInsertAfter(1, 9)
	list.FindAndInsertAfter(1, 9)
	list.FindAndInsertAfter(3, 2)

	list.PrintList()
	list.Reverse()
	list.PrintList()

	list.RemoveNode(4)
	list.RemoveNode(1)

	list.RemoveAllAccurances(9)
	// list.RemoveAllAccurances(2)
	list.RemoveAllAccurances(3)

	list.PrintList()
}
