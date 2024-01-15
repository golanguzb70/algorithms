package main

import (
	"fmt"

	leetcode "github.com/golanguzb70/algorithms/leetcode/go"
)

func main() {
	matches := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}
	fmt.Println(leetcode.FindWinners(matches))
}
