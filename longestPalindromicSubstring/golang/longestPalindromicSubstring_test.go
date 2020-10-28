package longestPalindromicSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		s      string
		expect string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"a",
			"a",
		},
		{
			"ac",
			"a",
		},
	}

	for _, testCase := range testCases {
		actual := longestPalindrome(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
