package Q148

// func maxSatisfied(customers []int, grumpy []int, minutes int) int {
// 	l, base, extra := len(customers), 0, 0

// 	// 計算初始滿意度
// 	for i := 0; i < minutes; i++ {
// 		if grumpy[i] == 0 {
// 			base += customers[i]
// 		} else {
// 			extra += customers[i]
// 		}
// 	}

// 	maxExtra := extra

// 	// 使用滑動視窗找出最大提升的區間
// 	for i := 0; i <= l; i++ {
// 		if grumpy[i] == 0 {
// 			base += customers[i]
// 		} else {
// 			extra += customers[i]
// 		}

// 		if grumpy[i-minutes] == 1 {
// 			extra -= customers[i-minutes]
// 		}
// 		if extra > maxExtra {
// 			maxExtra = extra
// 		}
// 	}

// 	// 最大滿意度為基礎滿意度加上最大提升
// 	return base + maxExtra
// }

type record struct {
	GrumpyPoint int
	Value       int
	Locate      int
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var r record
	l, maxV := len(customers), 0

	// 先算出來基本數值
	for i := 0; i < l; i++ {
		if grumpy[i] == 0 {
			maxV += customers[i]
		}
	}

	// 找出最容易生氣的區間，然後記錄座標
	for i := 0; i <= l-minutes; i++ {
		sumG, sumV := sum(grumpy[i:i+minutes]), sum(customers[i:i+minutes])
		// 如果相等就比較值
		if sumG > r.GrumpyPoint || (sumG == r.GrumpyPoint && sumV > r.Value) {
			r = record{
				GrumpyPoint: sumG,
				Value:       sumV,
				Locate:      i,
			}
		}
	}

	for i := r.Locate; i < r.Locate+minutes; i++ {
		if grumpy[i] == 0 {
			maxV -= customers[i]
		}
	}

	return maxV + r.Value
}

func sum(arr []int) int {
	var s int
	for _, v := range arr {
		s += v
	}
	return s
}
