package algo9

func lengthOfLastWord(s string) int {
	n := len(s) - 1
	begin, end := 0, n+1
	for i := n; i >= 0; i-- {
		if i-1 >= 0 && s[i] == ' ' && s[i-1] != ' ' {
			end = i
		}
		if i-1 >= 0 && s[i] != ' ' && s[i-1] == ' ' {
			begin = i
			break
		}
	}

	return end - begin
}
