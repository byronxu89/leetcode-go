package algo2

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	/*
		811. Subdomain Visit Count
		https://leetcode-cn.com/problems/subdomain-visit-count
		time complexity O(n)
		space complexity O(n)
	*/
	cache := make(map[string]int, len(cpdomains))
	for _, val := range cpdomains {
		splited := strings.Split(val, " ")
		count, _ := strconv.Atoi(splited[0])
		domains := strings.Split(splited[1], ".")
		for i := 0; i < len(domains); i++ {
			key := strings.Join(domains[i:], ".")

			if value, ok := cache[key]; ok {
				cache[key] = value + count
			} else {
				cache[key] = count
			}
		}
	}
	ans := make([]string, 0)
	for key, val := range cache {
		ans = append(ans, strconv.Itoa(val)+" "+key)
	}
	return ans
}

var SubdomainVisits = subdomainVisits
