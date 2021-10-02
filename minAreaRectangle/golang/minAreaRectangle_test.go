package minAreaRectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		points [][]int
		expect int
	}{
		{
			[][]int{
				{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2},
			},
			4,
		},
		{
			[][]int{
				{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3},
			},
			2,
		},
		{
			[][]int{
				{0, 1}, {1, 3}, {3, 3}, {4, 4}, {1, 4}, {2, 3}, {1, 0}, {3, 4},
			},
			2,
		},
		{
			[][]int{
				{3, 2}, {3, 1}, {4, 4}, {1, 1}, {4, 3}, {0, 3}, {0, 2}, {4, 0},
			},
			0,
		},
	}

	for _, testCase := range testCases {
		actual := minAreaRect(testCase.points)
		assert.Equal(t, testCase.expect, actual)
	}
}
