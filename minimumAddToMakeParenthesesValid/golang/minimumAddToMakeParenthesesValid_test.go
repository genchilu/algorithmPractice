package minimumAddToMakeParenthesesValid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minAddToMakeValid(s string) int {
func TestMinAddToMakeValid(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{
			"())",
			1,
		},
		{
			"(((",
			3,
		},
		{
			"()",
			0,
		},
		{
			"()))((",
			4,
		},
	}

	for _, testCase := range testCases {
		actual := minAddToMakeValid(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
