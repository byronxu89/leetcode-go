package algo9

import "strings"

func reverseWords(s string) string {
	ss := strings.Fields(s)
	var sb strings.Builder
	ln := len(ss) - 1
	sb.WriteString(ss[ln])
	for i := ln - 1; i >= 0; i-- {
		sb.WriteString(" ")
		sb.WriteString(ss[i])

	}

	return sb.String()

}
