package kthLargestElement

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		nums []int
		k int
		expect int
	} {
		{
			[]int{3,2,1,5,6,4},
			2,
			5,
		},
		{
			[]int{3,2,3,1,2,4,5,5,6},
			4,
			4,
		},
	}

	for _, testCase := range testCases {
		actualResult := findKthLargest(testCase.nums, testCase.k)
		
		assert.Equal(t, testCase.expect, actualResult)
	}
}