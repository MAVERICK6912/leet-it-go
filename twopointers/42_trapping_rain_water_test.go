package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	testData := []struct {
		height   []int
		expected int
	}{{
		height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		expected: 6,
	}, {
		height:   []int{4, 2, 0, 3, 2, 5},
		expected: 9,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, trap(test.height))
	}
}
