package numberOfDistinctIslands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinctIslands(t *testing.T) {
	testCases := []struct {
		grid   [][]int
		expect int
	}{
		// {
		// 	[][]int{
		// 		[]int{1, 1, 0, 0, 0},
		// 		[]int{1, 1, 0, 0, 0},
		// 		[]int{0, 0, 0, 1, 1},
		// 		[]int{0, 0, 0, 1, 1},
		// 	},
		// 	1,
		// },
		// {
		// 	[][]int{
		// 		[]int{1, 1, 0, 1, 1},
		// 		[]int{1, 0, 0, 0, 0},
		// 		[]int{0, 0, 0, 0, 1},
		// 		[]int{1, 1, 0, 1, 1},
		// 	},
		// 	3,
		// },
		{
			[][]int{
				[]int{0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0},
				[]int{0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0},
				[]int{0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0},
				[]int{1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
			},
			15,
		},
	}
	for _, testCase := range testCases {
		actual := numDistinctIslands(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}

}