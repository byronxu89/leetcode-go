package algo

func PlusOne(digits []int) []int {
	len := len(digits)
	if len <= 0 {
		return nil
	}

	var startpos = len - 1

	if startpos == 0 && digits[0] != 9 {
		return []int{digits[0] + 1}
	}

	for digits[startpos] == 9 {
		digits[startpos] = 0
		if startpos == 0 {
			break
		}
		startpos -= 1

	}
	if startpos == 0 && digits[0] == 0 {
		digits = append([]int{1}, digits...)
	} else {
		digits[startpos] += 1
	}
	return digits
}
