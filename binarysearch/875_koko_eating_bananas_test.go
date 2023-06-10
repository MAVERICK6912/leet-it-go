package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	testData := []struct {
		piles    []int
		hours    int
		expected int
	}{{
		piles:    []int{3, 6, 7, 11},
		hours:    8,
		expected: 4,
	}, {
		piles:    []int{30, 11, 23, 4, 20},
		hours:    5,
		expected: 30,
	}, {
		piles:    []int{30, 11, 23, 4, 20},
		hours:    6,
		expected: 23,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, minEatingSpeed(test.piles, test.hours))
	}
}
