package algo3

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	permutation := []int{}
	selection := make([]bool, n)

	var backtrack func(int)
	backtrack = func(index int) {

		if index == n {
			ans = append(ans, append([]int(nil), permutation...))
			return
		}
		for i, v := range nums {
			if selection[i] || i > 0 && !selection[i-1] && v == nums[i-1] {
				continue
			}
			permutation = append(permutation, v)
			selection[i] = true
			backtrack(index + 1)
			selection[i] = false
			permutation = permutation[:len(permutation)-1]

		}

	}
	backtrack(0)
	return ans
}

var PermuteUnique = permuteUnique
