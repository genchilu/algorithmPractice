package textJustification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func fullJustify(words []string, maxWidth int) []string {
func TestDev(t *testing.T) {

	testCases := []struct {
		words    []string
		maxWidth int
		expect   []string
	}{
		// {
		// 	[]string{"This", "is", "an", "example", "of", "text", "justification."},
		// 	16,
		// 	[]string{
		// 		"This    is    an",
		// 		"example  of text",
		// 		"justification.  ",
		// 	},
		// },
		// {
		// 	[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
		// 	16,
		// 	[]string{
		// 		"What   must   be",
		// 		"acknowledgment  ",
		// 		"shall be        ",
		// 	},
		// },
		{
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, testCase := range testCases {
		actual := fullJustify(testCase.words, testCase.maxWidth)
		assert.Equal(t, testCase.expect, actual)
	}
}
