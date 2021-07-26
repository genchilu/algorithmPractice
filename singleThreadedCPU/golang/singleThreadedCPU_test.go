package singleThreadedCPU

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//getOrder(tasks [][]int) []int
func TestGetOrder(t *testing.T) {
	testCases := []struct {
		tasks  [][]int
		expect []int
	}{
		{
			[][]int{
				{1, 2},
				{2, 4},
				{3, 2},
				{4, 1},
			},
			[]int{0, 2, 3, 1},
		},
		{
			[][]int{
				{7, 10},
				{7, 12},
				{7, 5},
				{7, 4},
				{7, 2},
			},
			[]int{4, 3, 2, 0, 1},
		},
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{
				{100, 100},
				{10, 100},
			},
			[]int{1, 0},
		},
	}

	for _, testCase := range testCases {
		actual := getOrder(testCase.tasks)
		assert.Equal(t, testCase.expect, actual)
	}
}
