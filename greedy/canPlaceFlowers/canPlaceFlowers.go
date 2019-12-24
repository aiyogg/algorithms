package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	cnt := 0
	for i := 0; i < length && cnt < n; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		var prev int
		var next int
		if i == 0 {
			prev = 0
		} else {
			prev = flowerbed[i-1]
		}
		if i == length-1 {
			next = 0
		} else {
			next = flowerbed[i+1]
		}
		if prev == 0 && next == 0 {
			cnt++
			flowerbed[i] = 1 // 需要隔一棵
		}
	}
	return cnt >= n
}
