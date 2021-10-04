package rangeAddition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func getModifiedArray(length int, updates [][]int) []int {
func TestDev(t *testing.T) {
	testCases := []struct {
		length  int
		updates [][]int
		expect  []int
	}{
		{
			5,
			[][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}},
			[]int{-2, 0, 3, 5, 3},
		},
		{
			10,
			[][]int{{2, 4, 6}, {5, 6, 8}, {1, 9, -4}},
			[]int{0, -4, 2, 2, 2, 4, 4, -4, -4, -4},
		},
	}

	for _, testCase := range testCases {
		actual := getModifiedArray(testCase.length, testCase.updates)
		assert.Equal(t, testCase.expect, actual)
	}
}
