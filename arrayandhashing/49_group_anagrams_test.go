package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	testData := []struct {
		data     []string
		expected [][]string
	}{{
		data:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		expected: [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
	}, {
		data:     []string{""},
		expected: [][]string{{""}},
	}, {
		data:     []string{"a"},
		expected: [][]string{{"a"}},
	}}

	for _, test := range testData {
		assert.ElementsMatch(t, test.expected, groupAnagrams(test.data))
	}
}
