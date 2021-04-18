package decodeWays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func numDecodings(s string) int {
func TestNumDecodings(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{
			"12",
			2,
		},
		{
			"226",
			3,
		},
		{
			"0",
			0,
		},
		{
			"06",
			0,
		},
		{
			"111111111111111111111111111111111111111111111",
			1836311903,
		},
	}

	for _, testCase := range testCases {
		actual := numDecodings(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
