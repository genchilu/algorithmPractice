package continuousSubarraySum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckSubarraySum(t *testing.T) {
	testCases := []struct{
		nums []int
		k int
		expect bool
	}{
		{
			[]int{23, 2, 4, 6, 7},
			6,
			true,
		},
		{
			[]int{23, 2, 6, 4, 7},
			6,
			true,
		},
		{
			[]int{0, 0, 0, 0},
			6,
			true,
		},
		{
			[]int{23,2,6,4,7},
			0,
			false,
		},
		{
			[]int{0,0},
			0,
			true,
		},
		{
			[]int{0,0,0,0,0},
			1,
			true,
		},
	}

	for _, testCase := range testCases {
		actual:=checkSubarraySum(testCase.nums, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
	
}