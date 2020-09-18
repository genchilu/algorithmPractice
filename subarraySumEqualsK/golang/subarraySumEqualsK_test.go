package subarraySumEqualsK

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	testCases := []struct{
		nums []int
		k int
		expect int
	}{
		{
			[]int{1,1,1},
			2,
			2,
		},
		{
			[]int{0,0,0},
			0,
			6,
		},
		{
			[]int{0,0,0},
			1,
			0,
		},
		{
			[]int{},
			1,
			0,
		},
		{
			[]int{1,2,3,4,5},
			0,
			0,
		},
		{			 
			[]int{0,0,0,0},
			0,
			10,
		},
	}

	for _, testCase := range testCases {
		actual := subarraySum(testCase.nums, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}