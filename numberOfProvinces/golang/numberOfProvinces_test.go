package numberOfProvinces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findCircleNum(isConnected [][]int) int {

func TestFindCircleNum(t *testing.T) {
	testCases := []struct {
		isConnected [][]int
		expect      int
	}{
		{
			[][]int{
				[]int{1, 1, 0},
				[]int{1, 1, 0},
				[]int{0, 0, 1},
			},
			2,
		},
		{
			[][]int{
				[]int{1, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 1},
			},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := findCircleNum(testCase.isConnected)
		assert.Equal(t, testCase.expect, actual)
	}
}
