package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	testData := []struct {
		s        string
		t        string
		expected bool
	}{{
		s:        "anagram",
		t:        "margana",
		expected: true,
	}, {
		s:        "rat",
		t:        "car",
		expected: false,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, isAnagram(test.s, test.t))
	}
}
