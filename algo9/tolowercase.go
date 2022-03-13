package algo9

func toLowerCase(s string) string {
	converted := make([]rune, len(s))

	for i, ch := range s {

		if ch >= 'A' && ch <= 'Z' {
			converted[i] = ch - 'A' + 'a'
		} else {
			converted[i] = ch
		}
	}

	return string(converted)
}
