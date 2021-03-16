package maximumAreaOfAPieceOfCakeAfterHorizontalAndVerticalCuts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
		expect         int
	}{
		{
			5,
			4,
			[]int{1, 2, 4},
			[]int{1, 3},
			4,
		},
		{
			5,
			4,
			[]int{3},
			[]int{3},
			9,
		},
	}

	for _, testCase := range testCases {
		actual := maxArea(testCase.h, testCase.w, testCase.horizontalCuts, testCase.verticalCuts)
		assert.Equal(t, testCase.expect, actual)
	}
}
