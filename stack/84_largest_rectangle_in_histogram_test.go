package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	testData := []struct {
		heights  []int
		expected int
	}{{
		heights:  []int{2, 1, 5, 6, 2, 3},
		expected: 10,
	}, {
		heights:  []int{2, 4},
		expected: 4,
	}, {
		heights:  []int{1},
		expected: 1,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, largestRectangleArea(test.heights))
	}
}
