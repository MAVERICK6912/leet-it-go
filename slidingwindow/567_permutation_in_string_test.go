package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	testData := []struct {
		s1       string
		s2       string
		expected bool
	}{{
		s1:       "ab",
		s2:       "eidbaooo",
		expected: true,
	}, {
		s1:       "ab",
		s2:       "eidboaoo",
		expected: false,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, checkInclusion(test.s1, test.s2))
	}
}
