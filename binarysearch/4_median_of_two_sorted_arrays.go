package binarysearch

import (
	"math"
)

// The idea is simple:
// 0. we know the median should be at index (lenA+lenB)/2
// 1. odd or even, finding this index solves the problem
// 2. we use binary search for the smaller array
// 3. first iteration we find the half way of nums1 of index i
// 4. to make the half of the overall, we take (m+n)/2-i from nums2 to be j
// 5. define Aleft, Aright, Bleft, Bright to be the left & right of nums1 & nums2 around the i & j
// 6. if Aleft<Bright && Bleft<Aright, then we found the correct indexes in both array
// 7. return the result based of odd or even of the result size
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
	// make sure first array is smaller: do binary search on smaller array
	// lowers the complexity
	if lenA > lenB {
		nums1, nums2 = nums2, nums1
		lenA, lenB = lenB, lenA
	}
	total := lenA + lenB
	half := (total + 1) / 2 // +1 because we want the size include mid point. i.e. math.Ceiling
	// define the left & right index of nums1 to be l, r = 0, len(nums1)-1
	// to calculate half point for binary search
	l, r := 0, lenA-1
	for {
		// next find half index of nums1: i = (l+r)/2
		// so i is either the mid or to the left of mid
		//
		// then we need to take from nums2 from 0 to j to make half of the overall
		// In other words: i+1+j+1 = half or j = half-i-2
		//
		// There is a trap: In go -1/2 is 0 not -1. If l is 0 and r is -1, we want to
		// i to be -1 so that it falls into the Aleft = -math.MaxInt condition. Otherwise
		// it goes into an infinite loop of l=r=h=0
		//
		// we can either use i := int(math.Floor(float64(l+r) / 2)) or below
		i := (l + r) >> 1
		j := half - i - 2                    // -2 because length is last_index+1 from 0 so i+1+j+1 = half
		var Aleft, Aright, Bleft, Bright int // left&right value of nums1&nums2 around i&j
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
		// now we want to see if this mid point is actually the mid point for the final array:
		// cross compare left & right of 1st & 2nd array
		if Aleft <= Bright && Bleft <= Aright {
			if (lenA+lenB)%2 == 0 {
				return float64(max(Aleft, Bleft)+min(Aright, Bright)) / 2
			}
			return float64(max(Aleft, Bleft))
		}
		// otherwise one of two conditions:
		if Aleft > Bright {
			// Aleft is too large, meaning we need to move left of nums1. putting right as half point
			r = i - 1
		} else {
			// Bleft is too large, or Aright is to small, meaning we need to move right of nums1
			l = i + 1
		}

	}
}
