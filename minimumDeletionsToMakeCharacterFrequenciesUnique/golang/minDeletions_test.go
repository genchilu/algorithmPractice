package minDeletions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minDeletions(s string) int {
func TestMinDeletions(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{
			"aab",
			0,
		},
		{
			"aaabbbcc",
			2,
		},
		{
			"abcabc",
			3,
		},
		{
			"bbcebab",
			2,
		},
	}

	for _, testCase := range testCases {
		actual := minDeletions(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
