package longestSubstringWithoutRepeatingCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func lengthOfLongestSubstring(s string) int {

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbbbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"",
			0,
		},
		{
			" ",
			1,
		},
		{
			"abba",
			2,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLongestSubstring(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
