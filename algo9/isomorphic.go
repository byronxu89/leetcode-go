package algo9

func isIsomorphic(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)
	for i := 0; i < l1; i++ {
		ch1 := s[i]
		ch2 := t[i]
		//[b] => [b]
		// [a] => [a]
		// [d] => [b]

		// [e] => [a]
		// [a] = > [e]
		// [g] => [d]
		// [d] => [g]
		v1, ok1 := map1[ch1]
		v2, ok2 := map2[ch2]
		if ok1 && v1 != ch2 || ok2 && v2 != ch1 {
			return false
		}
		map1[ch1] = ch2
		map2[ch2] = ch1

	}
	return true
}
