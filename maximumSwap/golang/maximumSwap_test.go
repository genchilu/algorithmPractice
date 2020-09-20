package maximumSwap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaximumSwap(t *testing.T) {
	testCases := []struct{
		num int
		expect int
	}{
		{
			5566,
			6565,
		},
		{
			2736,
			7236,
		},
		{
			9973,
			9973,
		},
		{
			98368,
			98863,
		},
	}

	for _, testCase := range testCases {
		actual:= maximumSwap(testCase.num)
		assert.Equal(t, testCase.expect, actual)
	}
}