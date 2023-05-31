package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	testData := []struct {
		nums     []int
		expected [][]int
	}{{
		nums:     []int{0, 1, 1},
		expected: [][]int{},
	}, {
		nums:     []int{-1, 0, 1, 2, -1, -4},
		expected: [][]int{{-1, 2, -1}, {0, 1, -1}},
	}, {
		nums:     []int{0, 0, 0},
		expected: [][]int{{0, 0, 0}},
	}}

	for _, test := range testData {
		assert.ElementsMatch(t, test.expected, threeSum(test.nums))
	}
}
