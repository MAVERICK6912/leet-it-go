package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		s        string
		expected bool
	}{{
		s:        "A man, a plan, a canal: Panama",
		expected: true,
	}, {
		s:        "race a car",
		expected: false,
	}, {
		s:        "  ",
		expected: true,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, isPalindrome(test.s))
	}
}
