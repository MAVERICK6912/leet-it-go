package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	testData := []struct {
		nums     []int
		expected int
	}{{
		nums:     []int{100, 4, 200, 1, 3, 2},
		expected: 4,
	}, {
		nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		expected: 9,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, longestConsecutive(test.nums))
	}
}
