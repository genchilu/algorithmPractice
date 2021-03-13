package longestSubstringWithAtMostKDistinctCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	testCases := []struct {
		s      string
		k      int
		expect int
	}{
		{
			"eceba",
			2,
			3,
		},
		{
			"aa",
			1,
			2,
		},
		{
			"a",
			2,
			1,
		},
		{
			"abaccc",
			2,
			4,
		},
		{
			"bacc",
			2,
			3,
		},
		{
			"ababacccccc",
			2,
			7,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLongestSubstringKDistinct(testCase.s, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
