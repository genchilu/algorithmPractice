package breakAPalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func breakPalindrome(palindrome string) string {
func TestBreakPalindrome(t *testing.T) {
	testCases := []struct {
		palindrome string
		expect     string
	}{
		{
			"abccba",
			"aaccba",
		},
		{
			"a",
			"",
		},
		{
			"aa",
			"ab",
		},
		{
			"aba",
			"abb",
		},
	}

	for _, testCase := range testCases {
		actual := breakPalindrome(testCase.palindrome)
		assert.Equal(t, testCase.expect, actual)
	}
}
