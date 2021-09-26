package arrayOfDoubledPairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		arr    []int
		expect bool
	}{
		// {
		// 	[]int{3, 1, 3, 6},
		// 	false,
		// },
		// {
		// 	[]int{2, 1, 2, 6},
		// 	false,
		// },
		// {
		// 	[]int{4, -2, 2, -4},
		// 	true,
		// },
		// {
		// 	[]int{1, 2, 4, 16, 8, 4},
		// 	false,
		// },
		{
			[]int{2, 4, 0, 0, 8, 1},
			true,
		},
	}

	for _, testCase := range testCases {
		actual := canReorderDoubled(testCase.arr)
		assert.Equal(t, testCase.expect, actual)
	}
}
