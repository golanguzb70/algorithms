package leetcode

import "slices"

/*
You are given an integer array matches where matches[i] = [winneri, loseri] indicates that the player winneri defeated player loseri in a match.

Return a list answer of size 2 where:
    answer[0] is a list of all players that have not lost any matches.
    answer[1] is a list of all players that have lost exactly one match.

The values in the two lists should be returned in increasing order.

Note:
    You should only consider the players that have played at least one match.
    The testcases will be generated such that no two matches will have the same outcome.

Algorithms:
	1. Gather all wins of the players
		a) Iterate through the array and gather using map
	2. Gather all losses of the players
		a) Iterate through the array and gather using map
	3. Iterate through the wins map
		a) If the loss is 0 append to the first sub array
	4. Iterate through the loses map
		a) if the loss is 1 append to the second sub array
*/

func FindWinners(matches [][]int) [][]int {
	res := make([][]int, 2)
	wins := map[int]bool{}
	loses := map[int]int{}

	for _, v := range matches {
		wins[v[0]] = true
		loses[v[1]]++
	}

	for k := range wins {
		if loses[k] == 0 {
			res[0] = append(res[0], k)
		}
	}

	for k, v := range loses {
		if v == 1 {
			res[1] = append(res[1], k)
		}
	}

	slices.Sort[[]int](res[0])
	slices.Sort[[]int](res[1])

	return res
}
