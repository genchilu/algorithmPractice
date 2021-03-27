package findClosestElements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findClosestElements(arr []int, k int, x int) []int {
func TestFindClosestElements(t *testing.T) {
	testCases := []struct {
		arr    []int
		k      int
		x      int
		expect []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			4,
			3,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
			-1,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
			3,
			4,
			[]int{3, 3, 4},
		},
	}

	for _, testCase := range testCases {
		actual := findClosestElements(testCase.arr, testCase.k, testCase.x)
		assert.Equal(t, testCase.expect, actual)
	}
}
