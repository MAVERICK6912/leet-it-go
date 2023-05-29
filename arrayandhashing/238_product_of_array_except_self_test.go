package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	testData := []struct {
		nums    []int
		expectd []int
	}{{
		nums:    []int{1, 2, 3, 4},
		expectd: []int{24, 12, 8, 6},
	}, {
		nums:    []int{-1, 1, 0, -3, 3},
		expectd: []int{0, 0, 9, 0, 0},
	}}
	for _, test := range testData {
		assert.ElementsMatch(t, test.expectd, productExceptSelf(test.nums))
	}
}
