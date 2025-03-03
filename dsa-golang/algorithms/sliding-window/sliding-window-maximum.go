package main

/*Problem:
You are given an array of integers nums, there is a sliding window of size k which is moving from the
very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding
window moves right by one position.

Return the max sliding window.
*/

/*
Solution:
Window size is static: k

Algrothm:
- Loop through the array saving latest max number index.
- When window slides check if last max number is in the window.
- If latest max number is in the window add it to response and continue sliding.
- If not search for max number in the window.

Time complexity: O(n)
Space complexity: O(1)
*/
type Node struct {
	val  int
	next *Node
	prev *Node
}

// Dequeue struct
type Dequeue struct {
	head *Node
	tail *Node
	size int
}

// PushBack adds an element to the back of the deque
func (dq *Dequeue) PushBack(val int) {
	newNode := &Node{val: val}

	if dq.tail == nil { // If deque is empty
		dq.head, dq.tail = newNode, newNode
	} else {
		dq.tail.next = newNode
		newNode.prev = dq.tail
		dq.tail = newNode
	}
	dq.size++
}

// PushFront adds an element to the front of the deque
func (dq *Dequeue) PushFront(val int) {
	newNode := &Node{val: val}

	if dq.head == nil { // If deque is empty
		dq.head, dq.tail = newNode, newNode
	} else {
		newNode.next = dq.head
		dq.head.prev = newNode
		dq.head = newNode
	}
	dq.size++
}

// Back returns the last element in the deque
func (dq *Dequeue) Back() int {
	if dq.tail == nil {
		return 0
	}
	return dq.tail.val
}

// Front returns the first element in the deque
func (dq *Dequeue) Front() int {
	if dq.head == nil {
		return 0
	}
	return dq.head.val
}

// PopBack removes an element from the back
func (dq *Dequeue) PopBack() {
	if dq.tail == nil {
		return
	}

	if dq.tail == dq.head { // If only one element
		dq.head, dq.tail = nil, nil
	} else {
		dq.tail = dq.tail.prev
		dq.tail.next = nil
	}
	dq.size--
}

// PopFront removes an element from the front
func (dq *Dequeue) PopFront() {
	if dq.head == nil {
		return
	}

	if dq.head == dq.tail { // If only one element
		dq.head, dq.tail = nil, nil
	} else {
		dq.head = dq.head.next
		dq.head.prev = nil
	}
	dq.size--
}

// Size returns the current size of the deque
func (dq *Dequeue) Size() int {
	return dq.size
}

// IsEmpty checks if the deque is empty
func (dq *Dequeue) IsEmpty() bool {
	return dq.size == 0
}

func MaxSlidingWindow(nums []int, k int) []int {
	response := make([]int, len(nums)-k+1)
	deque := Dequeue{}
	for i := range nums {
		for deque.size > 0 && deque.Back() < nums[i] {
			deque.PopBack()
		}

		deque.PushBack(nums[i])

		if i-k+1 >= 0 {
			response[i-k+1] = deque.Front()
			if nums[i-k+1] >= deque.Front() {
				deque.PopFront()
			}
		}

	}

	return response
}
