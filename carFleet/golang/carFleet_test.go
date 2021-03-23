package carFleet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func carFleet(target int, position []int, speed []int) int {
func TestCarFleet(t *testing.T) {
	testCases := []struct {
		target   int
		position []int
		speed    []int
		expect   int
	}{
		{
			12,
			[]int{10, 8, 0, 5, 3},
			[]int{2, 4, 1, 1, 3},
			3,
		},
		{
			16,
			[]int{11, 14, 13, 6},
			[]int{2, 2, 6, 7},
			2,
		},
		{
			21,
			[]int{1, 15, 6, 8, 18, 14, 16, 2, 19, 17, 3, 20, 5},
			[]int{8, 5, 5, 7, 10, 10, 7, 9, 3, 4, 4, 10, 2},
			7,
		},
	}

	for _, testCase := range testCases {
		actual := carFleet(testCase.target, testCase.position, testCase.speed)
		assert.Equal(t, testCase.expect, actual)
	}

}
