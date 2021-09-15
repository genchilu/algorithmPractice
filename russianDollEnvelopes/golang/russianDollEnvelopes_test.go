package russianDollEnvelopes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		envelopes [][]int
		expect    int
	}{
		{
			[][]int{
				{5, 4}, {6, 4}, {6, 7}, {2, 3},
			},
			3,
		},
		{
			[][]int{
				{1, 1}, {1, 1}, {1, 1},
			},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := maxEnvelopes(testCase.envelopes)
		assert.Equal(t, testCase.expect, actual)
	}
}
