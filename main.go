package main

import (
	"fmt"

	leetcode "github.com/golanguzb70/algorithms/leetcode/go"
)

func main() {
	res := leetcode.GenerateTrees(3)
	fmt.Println(res)
	for _, e := range res {
		leetcode.PrintNode(e)
		fmt.Println()
	}
}
