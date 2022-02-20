package algo5

func shipWithinDays(weights []int, days int) int {

	left, right := 0, 0
	for _, elem := range weights {
		if left < elem {
			left = elem
		}
		right += elem
	}

	for left < right {
		mid := left + (right-left)/2
		if canHold(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canHold(weights []int, capacity, limits int) bool {
	accum := 0
	days := 1
	for _, elem := range weights {
		if accum+elem > capacity {
			accum = 0
			days++
		}
		accum += elem
	}
	return days <= limits
}
