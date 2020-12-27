package letterCombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func letterCombinations(digits string) []string {

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		digits string
		expect []string
	}{
		// {
		// 	"2",
		// 	[]string{"a", "b", "c"},
		// },
		// {
		// 	"23",
		// 	[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		// },
		// {
		// 	"7",
		// 	[]string{"p", "q", "r", "s"},
		// },
		{
			"234",
			[]string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for _, testCase := range testCases {
		actual := letterCombinations(testCase.digits)
		assert.Equal(t, testCase.expect, actual)
	}
}
