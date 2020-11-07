package wordBreak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		s        string
		wordDict []string
		expect   bool
	}{
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"applepen",
			[]string{"apple", "pen"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			false,
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := wordBreak(testCase.s, testCase.wordDict)
		assert.Equal(t, testCase.expect, actual)
	}
}
