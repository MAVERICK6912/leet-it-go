package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testData := []struct {
		height   []int
		expected int
	}{{
		height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	}, {
		height:   []int{1, 1},
		expected: 1,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, maxArea(test.height))
	}
}
