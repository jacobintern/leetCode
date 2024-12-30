package Q148

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	l, base, extra := len(customers), 0, 0

	// 計算初始滿意度
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 0 {
			base += customers[i]
		} else {
			extra += customers[i]
		}
	}

	maxExtra := extra

	// 使用滑動視窗找出最大提升的區間
	for i := minutes; i < l; i++ {
		if grumpy[i] == 0 {
			base += customers[i]
		} else {
			extra += customers[i]
		}

		if grumpy[i-minutes] == 1 {
			extra -= customers[i-minutes]
		}
		if extra > maxExtra {
			maxExtra = extra
		}
	}

	// 最大滿意度為基礎滿意度加上最大提升
	return base + maxExtra
}
