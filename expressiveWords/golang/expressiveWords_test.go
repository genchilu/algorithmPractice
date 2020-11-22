package expressiveWords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressiveWords(t *testing.T) {
	testCases := []struct {
		S      string
		words  []string
		expect int
	}{
		{
			"heeellooo",
			[]string{"hello", "hi", "helo"},
			1,
		},
		{
			"dddiiiinnssssssoooo",
			[]string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := expressiveWords(testCase.S, testCase.words)
		assert.Equal(t, testCase.expect, actual)
	}
}
