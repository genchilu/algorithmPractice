package kClosestPointsToOrigin

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	//"fmt"
)

func TestKClosest(t *testing.T) {
	testCases := []struct{
		points [][]int
		k int
		expect [][]int
	} {
		{
			[][]int{
				[]int{1,3},
				[]int{-2,2},
			},
			1,
			[][]int{
				[]int{-2,2},
			},
		},
		{
			[][]int{
				[]int{3,3},
				[]int{5,-1},
				[]int{-2,4},
			},
			2,
			[][]int{
				[]int{3,3},
				[]int{-2,4},
			},
		},
		{
			[][]int{
				[]int{0,1},
				[]int{1,0},
			},
			2,
			[][]int{
				[]int{0,1},
				[]int{1,0},
			},
		},
	}

	for _, testCase := range testCases {
		actual := KClosest(testCase.points, testCase.k)
		assert.Equal(t, len(testCase.expect), len(actual))
		matchCount:=0
		for _,v := range actual {
			for _,vv :=range testCase.expect {
				if reflect.DeepEqual(v,vv) {
					matchCount++
				}
			}
		}
		assert.Equal(t, len(actual), matchCount)
	}
}