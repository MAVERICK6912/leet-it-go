package binarysearch

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenA, lenB := len(nums1), len(nums2)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if lenA > lenB {
		nums1, nums2 = nums2, nums1
		lenA, lenB = lenB, lenA
	}
	total := lenA + lenB
	half := (total + 1) / 2
	l, r := 0, lenA-1
	for {
		i := (l + r) >> 1
		j := half - i - 2
		var Aleft, Aright, Bleft, Bright int
		if i >= 0 {
			Aleft = nums1[i]
		} else {
			Aleft = math.MinInt
		}
		if i+1 < lenA {
			Aright = nums1[i+1]
		} else {
			Aright = math.MaxInt
		}
		if j >= 0 {
			Bleft = nums2[j]
		} else {
			Bleft = math.MinInt
		}
		if j+1 < lenB {
			Bright = nums2[j+1]
		} else {
			Bright = math.MaxInt
		}
		// check if partition is correct
		if Aleft <= Bright && Bleft <= Aright {
			if (lenA+lenB)%2 == 0 {
				return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2
			}
			return float64(max(Aleft, Bleft))
		}

		if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}

	}
}
