package sentenceScreenFitting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// wordsTyping(sentence []string, rows int, cols int) int {
func TestWordsTyping(t *testing.T) {
	testCases := []struct {
		sentence []string
		row      int
		cols     int
		expect   int
	}{
		{
			[]string{"hello", "world"},
			2,
			8,
			1,
		},
		{
			[]string{"a", "bcd", "e"},
			3,
			6,
			2,
		},
		{
			[]string{"I", "had", "apple", "pie"},
			4,
			5,
			1,
		},
		{
			[]string{"f", "p", "a"},
			8,
			7,
			10,
		},
		{
			[]string{"hello"},
			10000,
			1,
			0,
		},
	}

	for _, testCase := range testCases {
		actual := wordsTyping(testCase.sentence, testCase.row, testCase.cols)
		assert.Equal(t, testCase.expect, actual)
	}
}
