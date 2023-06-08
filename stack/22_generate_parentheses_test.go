package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	testData := []struct {
		n        int
		expected []string
	}{{
		n:        3,
		expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	}, {
		n:        1,
		expected: []string{"()"},
	}}
	for _, test := range testData {
		assert.ElementsMatch(t, test.expected, generateParenthesis(test.n))
	}
}
