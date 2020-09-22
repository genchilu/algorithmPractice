package isGraphBipartite

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestIsBipartite(t *testing.T) {
	testCases := []struct{
		graph [][]int
		expect bool
	}{
		{
			[][]int{
				[]int{1,3},
				[]int{0,2},
				[]int{1,3},
				[]int{0,2},
			},
			true,
		},
		{
			[][]int{
				[]int{1,2,3},
				[]int{0,2},
				[]int{0,1,3},
				[]int{0,2},
			},
			false,
		},
		{
			[][]int{
				[]int{1},
				[]int{0,3},
				[]int{3},
				[]int{1,2},
			},
			true,
		},
		{
			[][]int{
				[]int{4,1},
				[]int{0,2},
				[]int{1,3},
				[]int{2,4},
				[]int{3,0},
			},
			false,
		},
		{
			[][]int{
				[]int{},
				[]int{2,4,6},
				[]int{1,4,8,9},
				[]int{7,8},
				[]int{1,2,8,9},
				[]int{6,9},
				[]int{1,5,7,8,9},
				[]int{3,6,9},
				[]int{2,3,4,6,9},
				[]int{2,4,5,6,7,8},
			},
			false,
		},
	}
	for _, testCase := range testCases {
		actual := isBipartite(testCase.graph)
		assert.Equal(t, testCase.expect, actual)
	}
}



