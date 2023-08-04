package Q93

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := duration
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] >= duration {
			res += duration
		} else {
			res += timeSeries[i] - timeSeries[i-1]
		}

	}
	return res
}
