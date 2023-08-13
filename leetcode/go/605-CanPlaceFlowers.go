package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if length == 1 {
		return flowerbed[0] == 0 || n == 0
	}
	for i := 0; i < length; i++ {
		if i > 0 && i < length-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
				i++
			}
			} else if i == 0 {
				if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
					n--
					flowerbed[i] = 1
					i++
				}
		} else {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				n--
				flowerbed[i] = 1
				i++
			}
		}
		if n == 0 {
			return true;
		}
	}
	return n <= 0
}
