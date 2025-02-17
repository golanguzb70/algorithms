package main

import "fmt"

/*Problem:
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type Node struct {
	Val  int
	Prev *Node
}

type MinStack struct {
	Tail    *Node
	MinNode *Node
}

func Constructor() MinStack {
	return MinStack{
		Tail: nil,
	}
}

func (this *MinStack) Push(val int) {
	newNode := &Node{
		Val:  val,
		Prev: nil,
	}

	if this.Tail == nil {

		this.Tail = newNode
	} else {
		newNode.Prev = this.Tail

		this.Tail = newNode
	}

	if this.MinNode == nil {
		this.MinNode = newNode
	} else if this.MinNode.Val > newNode.Val {
		this.MinNode = newNode
	}
}

func (this *MinStack) Pop() {
	if this.Tail != nil {
		if this.Tail.Prev == nil {
			this.Tail = nil
			this.MinNode = nil
		} else {
			fmt.Printf("Tail: %p; MinNode: %p \n", this.Tail, this.MinNode)
			if fmt.Sprintf("%p", this.Tail) == fmt.Sprintf("%p", this.MinNode) {
				this.Tail = this.Tail.Prev
				this.MinNode = this.Tail
				this.InitMin()
			} else {
				this.Tail = this.Tail.Prev
			}
		}
	}
}

func (this *MinStack) InitMin() {
	current := this.Tail

	for current != nil {
		if current.Val < this.MinNode.Val {
			this.MinNode = current
		}
		current = current.Prev
	}
}

func (this *MinStack) Top() int {
	if this.Tail != nil {
		return this.Tail.Val
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if this.MinNode != nil {
		return this.MinNode.Val
	}

	return 0
}

func TestStack() {
	stack := Constructor()
	stack.Push(10)
	stack.Push(0)
	stack.Push(-3)

	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println("Minimum is", stack.GetMin())

	stack.Pop()
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())
}
