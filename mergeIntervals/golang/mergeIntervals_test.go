package mergeIntervals

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		intervals [][]int
		expectResult [][]int
	} {
		{
			[][]int{[]int{1,3}, []int{2,6}, []int{8,10}, []int{15,18}},
			[][]int{[]int{1,6}, []int{8,10}, []int{15,18}},
		},
		{
			[][]int{[]int{1,4}, []int{4,5}},
			[][]int{[]int{1,5}},
		},
	}

	for _, testCase := range testCases {
		actualResult := merge(testCase.intervals)
		fmt.Printf("%v\n", actualResult)
		assert.Equal(t, len(testCase.expectResult), len(actualResult))
		for idx := range(testCase.expectResult) {
			assert.Equal(t, testCase.expectResult[idx], actualResult[idx])
		}
	}
}