package binarysearch

import "math"

func minEatingSpeed(piles []int, h int) int {
	findMax := func(a []int) int {
		max := 0
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	left, right := 1, findMax(piles)
	res := right
	for left <= right {
		eatRate := left + (right-left)/2
		hoursTaken := getHoursToFinishPile(eatRate, piles)

		if hoursTaken <= h {
			res = min(res, eatRate)
			right = eatRate - 1
		} else {
			left = eatRate + 1
		}
	}
	return res
}

func getHoursToFinishPile(eatRate int, piles []int) int {
	hoursTaken := 0
	for _, pile := range piles {
		hoursTaken += int(math.Ceil(float64(pile) / float64(eatRate)))
	}
	return hoursTaken
}
