package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testData := []struct {
		nums     []int
		target   int
		expected []int
	}{{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{1, 2},
	}, {
		nums:     []int{2, 3, 4},
		target:   6,
		expected: []int{1, 3},
	}, {
		nums:     []int{-1, 0},
		target:   -1,
		expected: []int{1, 2},
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, twoSum(test.nums, test.target))
	}
}
