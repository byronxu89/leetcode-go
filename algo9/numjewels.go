package algo9

func numJewelsInStones(jewels string, stones string) int {
	len := len(jewels)
	mapping := make(map[rune]struct{}, len)
	for _, ch := range jewels {
		mapping[ch] = struct{}{}
	}
	count := 0
	for _, ch2 := range stones {
		if _, ok := mapping[ch2]; ok {
			count++
		}
	}
	return count
}
