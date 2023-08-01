package leetcode

/*https://leetcode.com/problems/combinations/*/

func Combine(n int, k int) [][]int {

	result := make([][]int, 0)

	var comb func(start int, curComb []int)

	comb = func(start int, curComb []int) {
		if len(curComb) == k {
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
			return
		}

		for i := start; i <= n; i++ {
			curComb = append(curComb, i)
			comb(i+1, curComb)
			curComb = curComb[:len(curComb)-1]
		}
	}
	comb(1, make([]int, 0))
	return result
}
