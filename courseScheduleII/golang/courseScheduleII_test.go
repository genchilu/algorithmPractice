package courseScheduleII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	testCases := []struct {
		numCourses    int
		prerequisites [][]int
		expect        []int
	}{
		{
			2,
			[][]int{
				[]int{1, 0},
			},
			[]int{0, 1},
		},
		{
			4,
			[][]int{
				[]int{1, 0},
				[]int{2, 0},
				[]int{3, 0},
				[]int{3, 2},
			},
			[]int{0, 1, 2, 3},
		},
		{
			4,
			[][]int{
				[]int{0, 1},
				[]int{3, 1},
				[]int{1, 3},
				[]int{3, 2},
			},
			[]int{},
		},
	}

	for _, testCase := range testCases {
		actual := findOrder(testCase.numCourses, testCase.prerequisites)
		assert.Equal(t, testCase.expect, actual)
	}
}
