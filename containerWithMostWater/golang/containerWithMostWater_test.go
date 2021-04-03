package containerWithMostWater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxArea(height []int) int {

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		height []int
		expect int
	}{
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{4, 3, 2, 1, 4},
			16,
		},
		{
			[]int{1, 2, 1},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := maxArea(testCase.height)
		assert.Equal(t, testCase.expect, actual)
	}
}
