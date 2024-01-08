package leetcode

func FindMatrix(nums []int) [][]int {
	result := [][]int{}
	mp := map[int]int{}

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for k, v := range mp {
		for i := 0; i < v; i++ {
			if len(result) >= i+1 {
				result[i] = append(result[i], k)
			} else {
				result = append(result, []int{k})
			}
		}
	}

	return result
}