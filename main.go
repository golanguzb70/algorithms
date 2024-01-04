package main

import (
	"fmt"

	leetcode "github.com/golanguzb70/algorithms/leetcode/go"
)

func main() {
	arr := []int{14,12,14,14,12,14,14,12,12,12,12,14,14,12,14,14,14,12,12}

	fmt.Println(leetcode.MinOperations(arr) == 7)
}
