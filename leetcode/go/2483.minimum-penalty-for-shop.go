package leetcode

import (
	"fmt"
	"strings"
)

func BestClosingTime(customers string) int {
	y, penalty := 0, 0
	lenth := len(customers)
	totalY := strings.Count(customers, "Y")
	totalN := lenth - totalY
	if totalY == lenth {
		return lenth
	}
	for i := 0; i < lenth; i++ {
		if customers[i] == 'N' {
			fmt.Println(totalY - y)
			if totalN-penalty > totalY-y {
				return i
			}
			penalty++
		} else {
			y++
		}
		if i == lenth-1 && customers[i] == 'Y' {
			return i + 1
		}
	}
	return 0
}
