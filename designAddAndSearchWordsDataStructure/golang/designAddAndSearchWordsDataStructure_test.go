package designAddAndSearchWordsDataStructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndSearch1(t *testing.T) {
	testCases := []struct {
		words       []string
		searchwords []string
		expect      []bool
	}{
		{
			[]string{"bad", "dad", "mad"},
			[]string{"pad", "bad", ".ad", "b.."},
			[]bool{false, true, true, true},
		},
		{
			[]string{"a", "ab"},
			[]string{"a", "a.", "ab", ".a", ".b", "ab.", ".", ".."},
			[]bool{true, true, true, false, true, false, true, true},
		},
	}

	for _, testCase := range testCases {
		wordDictionary := Constructor()
		for _, w := range testCase.words {
			wordDictionary.AddWord(w)
		}

		for i, w := range testCase.searchwords {
			actual := wordDictionary.Search(w)
			assert.Equal(t, testCase.expect[i], actual)
		}
	}
}
