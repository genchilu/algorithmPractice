package findFirstAndLastPositionOfElementInSortedArray

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		nums []int
		target int
		expectResult []int
	} {
		{
			[]int{},
			0,
			[]int{-1, -1},
		},
		{
			make([]int, 1000000000, 1000000000),
			0,
			[]int{-1, -1},
		},
		{
			make([]int, 100, 100),
			100000000000,
			[]int{-1, -1},
		},
		{
			make([]int, 100, 100),
			-100000000000,
			[]int{-1, -1},
		},
		{
			[]int{5,7,7,8,8,10},
			8,
			[]int{3,4},
		},
		{
			[]int{5,7,7,8,8,10},
			6,
			[]int{-1, -1},
		},
		{
			[]int{5,7,7,8,9,10},
			8,
			[]int{3,3},
		},
	}

	for _, testCase := range testCases {
		actualResult := searchRange(testCase.nums, testCase.target)
		
		assert.Equal(t, len(testCase.expectResult), len(actualResult))
		for idx := range(testCase.expectResult) {
			assert.Equal(t, testCase.expectResult[idx], actualResult[idx])
		}
	}
}