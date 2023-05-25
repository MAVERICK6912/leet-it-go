package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	testData := []struct {
		nums     []int
		k        int
		expected []int
	}{{
		nums:     []int{1, 1, 1, 2, 2, 3},
		k:        2,
		expected: []int{1, 2},
	}, {
		nums:     []int{1},
		k:        1,
		expected: []int{1},
	}}

	for _, test := range testData {
		assert.ElementsMatch(t, test.expected, topKFrequent(test.nums, test.k))
	}
}
