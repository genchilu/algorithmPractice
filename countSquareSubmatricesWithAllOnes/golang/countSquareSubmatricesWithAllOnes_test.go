package countSquareSubmatricesWithAllOnes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSquares(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		expect int
	}{
		{
			[][]int{
				[]int{0, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{0, 1, 1, 1},
			},
			15,
		},
		{
			[][]int{
				[]int{1, 0, 1},
				[]int{1, 1, 0},
				[]int{1, 1, 0},
			},
			7,
		},
	}

	for _, testCase := range testCases {
		actual := countSquares(testCase.matrix)
		assert.Equal(t, testCase.expect, actual)
	}
}
