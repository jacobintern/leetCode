package Q49

import (
	"math"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return uint32(0)
	}

	bytStr := DecToBin(num)
	res := 0.0
	for i := 0; i < 32; i++ {
		v, _ := strconv.Atoi(string(bytStr[i]))
		res += math.Pow(2, float64(i)) * float64(v)
	}
	return uint32(res)
}

func DecToBin(num uint32) string {
	res := ""
	for num > 0 {
		tmp := int(num % 2)
		res = strconv.Itoa(tmp) + res
		num /= 2
	}

	if strLen := len(res); strLen < 32 {
		for i := 32 - strLen; i > 0; i-- {
			res = "0" + res
		}
	}

	return res
}

// 位移解法
func reverseBits2(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// res 往左移一位後，把 num 的個位數加入
		res = res<<1 + num&1
		// 因為個位數已經加到 res 了所以向右移一格
		num >>= 1
	}
	return res
}
