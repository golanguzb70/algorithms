package main

import (
	"fmt"
)

func main() {

	// input := []int{1,1,1,2,3} // 4
	input := []int{1,1,1,2,2,3} // 5
	// input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3} // 7
	fmt.Println(RemoveDuplicates(input))
	fmt.Println(input)
}
