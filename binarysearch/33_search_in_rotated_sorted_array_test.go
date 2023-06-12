package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotated(t *testing.T) {
	testData := []struct {
		nums     []int
		target   int
		expected int
	}{{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
	}, {
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
	}, {
		nums:     []int{1},
		target:   0,
		expected: -1,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, searchRotated(test.nums, test.target))
	}
}
