package main

import (
	"strconv"
)

/*Problem:
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/

/*Solution:

Reverse Polish Notation - pop 2 items from stack and implement the arithmetic and push 1 result to the stack.

Algorithm:
1. Loop through the input array.
2. Push to the stack if element is number.
3. Pop last to value from stack and push the calculation if element is math operation.
*/

type RPNStack struct {
	items []int
}

func (stack *RPNStack) Push(val int) {
	stack.items = append(stack.items, val)
}

func (stack *RPNStack) Pop() int {
	if len(stack.items) != 0 {
		response := stack.items[len(stack.items)-1]
		stack.items = stack.items[:len(stack.items)-1]
		return response
	}

	return 0
}

func EvalRPN(tokens []string) int {

	stack := RPNStack{}
	for _, e := range tokens {
		intValue, err := strconv.Atoi(e)
		if err == nil {
			stack.Push(intValue)
			continue
		}

		num2 := stack.Pop()
		num1 := stack.Pop()
		switch e {
		case "+":
			stack.Push(num1 + num2)
		case "-":
			stack.Push(num1 - num2)
		case "*":
			stack.Push(num1 * num2)
		case "/":
			stack.Push(num1 / num2)
		}
	}

	return stack.Pop()
}
