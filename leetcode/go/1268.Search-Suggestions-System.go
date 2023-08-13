package leetcode

import (
	"sort"
	"strings"
)

func SuggestedProducts(products []string, searchWord string) [][]string {
	response := [][]string{}
	var subArr []string
	leng := len(searchWord)
	for i := 0; i < leng; i++ {
		prefix := searchWord[:i+1]
		subArr = []string{}
		for _, e := range products {
			if strings.HasPrefix(e, prefix) {
				subArr = append(subArr, e)
			}
		}
		sort.StringSlice.Sort(subArr)
		if len(subArr) >= 3 {
			response = append(response, subArr[:3])
		} else {
			response = append(response, subArr)
		}
	}

	return response
}
