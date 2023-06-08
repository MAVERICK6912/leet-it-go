package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	testData := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{{
		target:   12,
		position: []int{10, 8, 0, 5, 3},
		speed:    []int{2, 4, 1, 1, 3},
		expected: 3,
	}, {
		target:   10,
		position: []int{3},
		speed:    []int{3},
		expected: 1,
	}, {
		target:   100,
		position: []int{0, 2, 4},
		speed:    []int{4, 2, 1},
		expected: 1,
	}, {
		target:   10,
		position: []int{6, 8},
		speed:    []int{3, 2},
		expected: 2,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, carFleet(test.target, test.position, test.speed))
	}
}
