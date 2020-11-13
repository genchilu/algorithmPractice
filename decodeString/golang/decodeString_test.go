package decodeString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		s      string
		expect string
	}{
		{
			"3[a]2[bc]",
			"aaabcbc",
		},
		{
			"3[a2[c]]",
			"accaccacc",
		},
		{
			"2[abc]3[cd]ef",
			"abcabccdcdcdef",
		},
		{
			"abc3[cd]xyz",
			"abccdcdcdxyz",
		},
		{
			"3[a]2[b4[F]c]",
			"aaabFFFFcbFFFFc",
		},
	}

	for _, testCase := range testCases {
		actual := decodeString(testCase.s)

		assert.Equal(t, testCase.expect, actual)
	}
}
