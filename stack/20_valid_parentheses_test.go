package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testData := []struct {
		s        string
		expected bool
	}{{
		s:        "()",
		expected: true,
	}, {
		s:        "()[]{}",
		expected: true,
	}, {
		s:        "(]",
		expected: false,
	}, {
		s:        "{{{[[[((()))]]]}}}",
		expected: true,
	}}
	for _, test := range testData {
		assert.Equal(t, test.expected, isValid(test.s))
	}
}
