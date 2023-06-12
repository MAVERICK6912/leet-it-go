package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{{
		nums:     []int{3, 4, 5, 1, 2},
		expected: 1,
	}, {
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		expected: 0,
	}, {
		nums:     []int{11, 13, 15, 17},
		expected: 11,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, findMin(test.nums))
	}
}
