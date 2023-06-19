package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testData := []struct {
		prices   []int
		expected int
	}{{
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	}, {
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, maxProfit(test.prices))
	}
}
