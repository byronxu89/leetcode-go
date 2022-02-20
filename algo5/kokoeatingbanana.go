package algo5

import "math"

func minEatingSpeed(piles []int, h int) int {
	/*
	  二分 1 ... max(N)
	   [3, 6, 7, 11]
	   k >= x
	  possible (k) = true
	   possible(1), possible(2)  false   ..... possible(k) = true possible (k+1) true
	*/
	left, right := 1, math.MaxInt32

	for left < right {
		mid := left + (right-left)/2
		if !possibleK(piles, mid, h) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func possibleK(piles []int, K, H int) bool {

	times := 0
	for i := 0; i < len(piles); i++ {
		times += (piles[i]-1)/K + 1
	}
	return times <= H

}
