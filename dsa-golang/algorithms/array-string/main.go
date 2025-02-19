package main

import (
	"fmt"
)

func main() {

	// input := []int{1,1,1,2,3} // 4
	input := []int{-1, 0, 1, 2, -1, -4} // 5
	// input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3} // 7
	// slices.Sort(input)
	fmt.Println(ThreeSum(input))
}
