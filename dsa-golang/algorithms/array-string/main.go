package main

import (
	"fmt"
)

func main() {
	input := "/.../a/../b/c/../d/./"
	fmt.Println(SimplifyPath(input))
}
