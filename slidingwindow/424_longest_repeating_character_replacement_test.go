package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	testData := []struct {
		s        string
		k        int
		expected int
	}{{
		s:        "ABAB",
		k:        2,
		expected: 4,
	}, {
		s:        "AABABBA",
		k:        1,
		expected: 4,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, characterReplacement(test.s, test.k))
	}
}
