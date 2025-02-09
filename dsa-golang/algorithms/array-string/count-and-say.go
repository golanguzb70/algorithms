package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*Problem:
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the run-length encoding of countAndSay(n - 1).
Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

Given a positive integer n, return the nth element of the count-and-say sequence.
*/

/*Solution:
1. Write a function that maps string to int array
2. Write a function that maps array to string
3. Call them recursively n-1 times.

Time complexity: O(n)
Space complexity: O(n)

*/

func StringToArray(str string) []string {
	answer := []string{}
	startingIndexOfValue := 0
	lastValueCounting := rune(str[0])
	for i, v := range str {
		if lastValueCounting != v {
			intValue, _ := strconv.Atoi(string(lastValueCounting))
			answer = append(answer, fmt.Sprintf("%d%d", i-startingIndexOfValue, intValue))
			lastValueCounting = v
			startingIndexOfValue = i
			if i == len(str)-1 {
				answer = append(answer, fmt.Sprintf("1%s", string(lastValueCounting)))
			}
		} else if i == len(str)-1 {
			intValue, _ := strconv.Atoi(string(lastValueCounting))
			answer = append(answer, fmt.Sprintf("%d%d", i-startingIndexOfValue+1, intValue))
		}
	}

	return answer
}

func CountAndSayRecursive(n int, str string) string {
	if n == 0 {
		return str
	}

	str = CountAndSayRecursive(n-1, str)
	str = strings.Join(StringToArray(str), "")

	// fmt.Printf("N: %d, Str: %s\n", n, str)

	return str
}

func countAndSayLoop(n int) string {
	initString := "1"

	for i := 1; i < n; i++ {
		initString = strings.Join(StringToArray(initString), "")
	}

	return initString
}

func CountAndSay(n int) string {
	answer := countAndSayLoop(n - 1)

	return answer
}
