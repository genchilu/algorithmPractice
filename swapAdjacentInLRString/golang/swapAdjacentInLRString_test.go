package swapAdjacentInLRString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//canTransform(start string, end string) bool {

func TestCanTransform(t *testing.T) {
	testCases := []struct {
		start  string
		end    string
		expect bool
	}{
		{
			"RXXLRXRXL",
			"XRLXXRRLX",
			true,
		},
		{
			"X",
			"L",
			false,
		},
		{
			"LLR",
			"RRL",
			false,
		},
		{
			"XL",
			"LX",
			true,
		},
		{
			"XLLR",
			"LXLX",
			false,
		},
		{
			"LXXLXRLXXL",
			"XLLXRXLXLX",
			false,
		},
	}

	for _, testCase := range testCases {
		actual := canTransform(testCase.start, testCase.end)
		assert.Equal(t, testCase.expect, actual)
	}
}
