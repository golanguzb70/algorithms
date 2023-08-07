package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	var start, end = 0, len(matrix) - 1
	var mid, mid1, high, low = 0, 0, 0, 0
	for start <= end {
		if start == len(matrix)-1 {
			if matrix[start][len(matrix[start])-1] < target {
				return false
			}
		}
		mid = (start + end) / 2
		if matrix[mid][0] > target {
			end = mid - 1
		} else if matrix[mid][len(matrix[mid])-1] < target {
			start = mid + 1
		} else {
			low, high = 0, len(matrix[mid])-1
			for low <= high {
				mid1 = (low + high) / 2
				if matrix[mid][mid1] == target {
					return true
				} else if matrix[mid][mid1] > target {
					high = mid1 - 1
				} else {
					low = mid1 + 1
				}
			}
			return false
		}

	}
	return false
}
