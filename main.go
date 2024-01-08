package main

import (
	"fmt"

	leetcode "github.com/golanguzb70/algorithms/leetcode/go"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(leetcode.FindContentChildren(g, s))
}
