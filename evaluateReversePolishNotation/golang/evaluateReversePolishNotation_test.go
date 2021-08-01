package evaluateReversePolishNotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//evalRPN(tokens []string) int {
func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		tokens []string
		expect int
	}{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for _, testCase := range testCases {
		actual := evalRPN(testCase.tokens)
		assert.Equal(t, testCase.expect, actual)
	}
}
