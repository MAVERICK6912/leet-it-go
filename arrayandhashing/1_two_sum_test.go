package arrayandhashing

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
		expected: []int{0, 1},
	}, {
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	}, {
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	}}
	for _, test := range testData {
		assert.ElementsMatch(t, test.expected, twoSum(test.nums, test.target))
	}
}
