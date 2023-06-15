package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []struct {
		s        string
		expected int
	}{{
		s:        "abcabcbb",
		expected: 3,
	}, {
		s:        "bbbbb",
		expected: 1,
	}, {
		s:        "pwwkew",
		expected: 3,
	}, {
		s:        " ",
		expected: 1,
	}, {
		s:        "dvdf",
		expected: 3,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, lengthOfLongestSubstring(test.s))
	}
}
