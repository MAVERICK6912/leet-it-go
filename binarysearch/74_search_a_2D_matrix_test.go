package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	testData := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{{
		matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		target:   3,
		expected: true,
	}, {
		matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		target:   13,
		expected: false,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, searchMatrix(test.matrix, test.target))
	}
}
