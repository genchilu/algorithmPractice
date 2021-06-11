package buildingsWithAnOceanView

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBuildings(t *testing.T) {
	testCases := []struct {
		heights []int
		expect  []int
	}{
		{
			[]int{4, 2, 3, 1},
			[]int{0, 2, 3},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{0, 1, 2, 3},
		},
		{
			[]int{1, 3, 2, 4},
			[]int{3},
		},
		{
			[]int{2, 2, 2, 2},
			[]int{3},
		},
	}

	for _, testCase := range testCases {
		actual := findBuildings(testCase.heights)
		assert.Equal(t, testCase.expect, actual)
	}
}
