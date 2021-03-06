package kargerMinCut

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestKargerMinCut(t *testing.T) {
	testCases := []struct{
		inputEdges [][]int
		expectMinCut int
	} {
		{
			[][]int{},
			0,
		},
		{
			nil,
			0,
		},
		{
			[][]int{
				[]int{0,1},
				[]int{1,2},
				[]int{2,3},
				[]int{0,4},
				[]int{1,5},
				[]int{2,6},
				[]int{3,7},
				[]int{4,5},
				[]int{5,6},
				[]int{6,7},
				[]int{0,5},
				[]int{1,4},
				[]int{2,7},
				[]int{3,6},
			},
			2,
		},
		{
			[][]int{
				[]int{0,1},
				[]int{0,2},
				[]int{0,3},
				[]int{0,4},
				[]int{1,2},
				[]int{1,3},
				[]int{1,4},
				[]int{2,3},
				[]int{2,4},
				[]int{3,4},
				[]int{5,6},
				[]int{5,7},
				[]int{5,8},
				[]int{5,9},
				[]int{6,7},
				[]int{6,8},
				[]int{6,9},
				[]int{7,8},
				[]int{7,9},
				[]int{8,9},
				[]int{0,5},
				[]int{2,6},
				[]int{4,7},
			},
			3,
		},
	}

	for _, testCase := range testCases {
		acutalMinCut := doMinCut(testCase.inputEdges)
		assert.Equal(t, testCase.expectMinCut, acutalMinCut)
	}
}