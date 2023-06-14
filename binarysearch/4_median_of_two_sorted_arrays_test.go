package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testData := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{{
		nums1:    []int{1, 3},
		nums2:    []int{2},
		expected: float64(2),
	}, {
		nums1:    []int{1, 2},
		nums2:    []int{3, 4},
		expected: float64(2.5),
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, findMedianSortedArrays(test.nums1, test.nums2))
	}
}
