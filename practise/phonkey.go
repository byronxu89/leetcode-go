package main

import (
	"fmt"
)

var phoneKeys = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}
func backtrack(digits string, idx int, combination string) {
	if idx == len(digits) {
		combinations = append(combinations, combination)
		return
	} else {
		digit := string(digits[idx])
		letters := phoneKeys[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, idx+1, combination+string(letters[i]))
		}

	}

}

func main() {
	str := "23"
	ret := letterCombinations(str)
	fmt.Println(ret)

}
