package arrayandhashing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	testData := []struct {
		input          []string
		expextedEncode string
	}{
		{
			input:          []string{"lint", "code", "love", "you"},
			expextedEncode: "4~lint4~code4~love3~you",
		}, {
			input:          []string{"we", "say", "~", "yes"},
			expextedEncode: "2~we3~say1~~3~yes",
		},
		{
			input:          []string{"aquaman", "batman", "superman", "wonderwoman", "hawkgirl", "greenlantern", "the martian"},
			expextedEncode: "7~aquaman6~batman8~superman11~wonderwoman8~hawkgirl12~greenlantern11~the martian",
		}}
	for _, test := range testData {
		assert.Equal(t, test.expextedEncode, encode(test.input))
		assert.Equal(t, test.input, decode(test.expextedEncode))
	}
}
