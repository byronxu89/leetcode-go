package main

import (
	. "app/algo2"
	"fmt"
)

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	res := SubdomainVisits(cpdomains)
	for _, elem := range res {
		fmt.Println(elem)
	}
	fmt.Println("======================================")
	nums := []int{1, 2, 2, 3, 1, 4, 2}
	res2 := FindShortestSubArray(nums)
	fmt.Println(res2)
}
