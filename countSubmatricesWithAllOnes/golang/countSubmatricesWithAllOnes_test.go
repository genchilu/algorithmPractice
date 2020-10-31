package countSubmatricesWithAllOnes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubmat(t *testing.T) {
	testCases := []struct {
		mat    [][]int
		expect int
	}{
		// {
		// 	[][]int{
		// 		[]int{1, 0, 1},
		// 		[]int{1, 1, 0},
		// 		[]int{1, 1, 0},
		// 	},
		// 	13,
		// },
		// {
		// 	[][]int{
		// 		[]int{1, 0, 1},
		// 		[]int{0, 1, 0},
		// 		[]int{1, 0, 1},
		// 	},
		// 	5,
		// },
		{
			[][]int{
				[]int{0, 1, 1, 0},
				[]int{0, 1, 1, 1},
				[]int{1, 1, 1, 0},
			},
			24,
		},
	}

	for _, testCase := range testCases {
		actual := numSubmat(testCase.mat)

		assert.Equal(t, testCase.expect, actual)
	}
}
