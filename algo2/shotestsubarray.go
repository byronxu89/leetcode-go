package algo2

func findShortestSubArray(nums []int) int {
	/*
		697. degree of array
		https://leetcode-cn.com/problems/degree-of-an-array
		time complexity O(n)
		space complexity O(n)
	*/
	pairs := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if arrayval, ok := pairs[nums[i]]; ok {
			arrayval[0]++
			arrayval[2] = i
			pairs[nums[i]] = arrayval
		} else {
			pairs[nums[i]] = []int{1, i, i}
		}

	}
	maxnum, minlen := 0, 0
	for _, val := range pairs {
		if val[0] > maxnum {
			maxnum = val[0]
			minlen = val[2] - val[1] + 1
		} else if val[0] == maxnum {
			tempmin := val[2] - val[1] + 1
			if minlen > tempmin {
				minlen = tempmin
			}

		}
	}
	return minlen
}
