package productOfArrayExceptSelf

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		nums []int
		expect []int
	} {
		{
			[]int{1,2,3,4},
			[]int{24,12,8,6},
		},
		{
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
		},
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			[]int{1, 1},
			[]int{1, 1},
		},
		{
			[]int{9, 0, -2},
			[]int{0, -18, 0},
		},
		{
			[]int{0, 4, 0},
			[]int{0, 0, 0},
		},
		{
			[]int{5,9,2,-9,-9,-7,-8,7,-9,10},
			[]int{-51438240,-28576800,-128595600,28576800,28576800,36741600,32148900,-36741600,28576800,-25719120},
		},
	}

	for _, testCase := range testCases {
		actualResult := productExceptSelf(testCase.nums)
		
		assert.Equal(t, testCase.expect, actualResult)
	}
}