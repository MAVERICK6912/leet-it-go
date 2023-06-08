package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	testData := []struct {
		ops      []string
		input    []int
		expected []int
	}{{
		ops:      []string{"minstack", "push", "push", "push", "getmin", "pop", "top", "getmin"},
		input:    []int{-1, -2, 0, -3, -1, -1, -1, -1},
		expected: []int{-1, -1, -1, -1, -3, -1, 0, -2},
	}, {
		ops:      []string{"minstack", "push", "push", "push", "getmin", "top", "pop", "getmin"},
		input:    []int{-1, -2, 0, -1, -1, -1, -1, -1},
		expected: []int{-1, -1, -1, -1, -2, -1, -1, -2},
	}, {
		ops:      []string{"minstack", "push", "push", "push", "push", "getmin", "pop", "getmin", "pop", "getmin", "pop", "getmin"},
		input:    []int{-1, 2, 0, 3, 0, -1, -1, -1, -1, -1, -1, -1},
		expected: []int{-1, -1, -1, -1, -1, 0, -1, 0, -1, 0, -1, 2},
	}}
	for _, test := range testData {
		var minStack MinStack
		for index, op := range test.ops {
			switch op {
			case "minstack":
				minStack = Constructor()
			case "push":
				minStack.Push(test.input[index])
			case "pop":
				minStack.Pop()
			case "top":
				assert.Equal(t, test.expected[index], minStack.Top())
			case "getmin":
				assert.Equal(t, test.expected[index], minStack.GetMin())
			}
		}
	}
}
