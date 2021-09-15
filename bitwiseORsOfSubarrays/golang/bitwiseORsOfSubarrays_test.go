package bitwiseORsOfSubarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		arr    []int
		expect int
	}{
		{
			[]int{0},
			1,
		},
		{
			[]int{1, 1, 2},
			3,
		},
		{
			[]int{1, 2, 4},
			6,
		},
	}

	for _, testCase := range testCases {
		actual := subarrayBitwiseORs(testCase.arr)
		assert.Equal(t, testCase.expect, actual)
	}
}
