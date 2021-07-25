package numberOfGoodWaysToSplitAString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//numSplits(s string) int

func TestNumSplits(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		{
			"aacaba",
			2,
		},
		{
			"abcd",
			1,
		},
		{
			"aaaaa",
			4,
		},
		{
			"acbadbaada",
			2,
		},
	}

	for _, testCase := range testCases {
		actual := numSplits(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
