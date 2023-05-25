package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testData := []struct {
		input    []int
		expected bool
	}{{
		input:    []int{1, 2, 3, 1},
		expected: true,
	}, {
		input:    []int{1, 2, 3},
		expected: false,
	}, {
		input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		expected: true,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, containsDuplicate(test.input))
	}
}
