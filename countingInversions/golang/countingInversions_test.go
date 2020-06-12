package countingInversion

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCountInversions(t *testing.T) {
	testCases := []struct{
		input []uint
		expect int
	}{
		{nil, 0},
		{[]uint{}, 0},
		{[]uint{1}, 0},
		{[]uint{1,2}, 0},
		{[]uint{2,1}, 1},
		{[]uint{1,1}, 0},
		{[]uint{1,2,3,4,5}, 0},
		{[]uint{5,4,3,2,1}, 10},
		{[]uint{7,5,3,3,9,6,7}, 8},
		{[]uint{23, 42, 32, 64, 12, 4}, 10},
		{[]uint{2,1,2,2,2,2,2}, 1},
		{[]uint{10,12,9,13,8,14,7}, 12},
	}

	for _, testCase := range testCases {
		actualResult := CountInversions(testCase.input)
		assert.Equal(t, testCase.expect, actualResult)
	}
}