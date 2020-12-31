package maximumLengthOfAConcatenatedStringWithUniqueCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxLength(arr []string) int {
func TestMaxLength(t *testing.T) {
	testCases := []struct {
		arr    []string
		expect int
	}{
		{
			[]string{"un", "iq", "ue"},
			4,
		},
		{
			[]string{"cha", "r", "act", "ers"},
			6,
		},
		{
			[]string{"abcdefghijklmnopqrstuvwxyz"},
			26,
		},
		{
			[]string{"a", "abc", "d", "de", "def"},
			6,
		},
	}

	for _, testCase := range testCases {
		actual := maxLength(testCase.arr)
		assert.Equal(t, testCase.expect, actual)
	}
}
