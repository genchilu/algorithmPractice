package wordLadder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	testCases := []struct {
		begin    string
		end      string
		wordList []string
		expect   int
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			5,
		},
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			0,
		},
	}

	for _, testCase := range testCases {
		actual := ladderLength(testCase.begin, testCase.end, testCase.wordList)
		assert.Equal(t, testCase.expect, actual)
	}
}
