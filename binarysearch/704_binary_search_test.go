package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testData := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		}, {
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		}, {
			nums:     []int{1, 5},
			target:   1,
			expected: 0,
		}}
	for _, test := range testData {
		assert.Equal(t, test.expected, search(test.nums, test.target))
	}
}
