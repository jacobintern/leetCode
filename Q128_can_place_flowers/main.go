package Q128

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}

		if flowerbed[i] == 1 {
			continue
		}

		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}

		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}

		flowerbed[i] = 1
		n--

	}
	return n <= 0
}
