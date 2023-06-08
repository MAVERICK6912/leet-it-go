package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	testData := []struct {
		tokens   []string
		expected int
	}{{
		tokens:   []string{"2", "1", "+", "3", "*"},
		expected: 9,
	}, {
		tokens:   []string{"4", "13", "5", "/", "+"},
		expected: 6,
	}, {
		tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		expected: 22,
	}}

	for _, test := range testData {
		assert.Equal(t, test.expected, evalRPN(test.tokens))
	}
}
