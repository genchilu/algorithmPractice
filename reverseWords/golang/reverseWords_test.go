package reverseWords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func reverseWords(s string) string {
func TestReverseWords(t *testing.T) {
	testCases := []struct {
		s      string
		expect string
	}{
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"  hello world  ",
			"world hello",
		},
		{
			"a good   example",
			"example good a",
		},
		{
			"  Bob    Loves  Alice   ",
			"Alice Loves Bob",
		},
		{
			"Alice does not even like bob",
			"bob like even not does Alice",
		},
	}

	for _, testCase := range testCases {
		actual := reverseWords(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
