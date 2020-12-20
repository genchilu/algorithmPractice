package TimeNeededToInformAllEmployees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//numOfMinutes(n int, headID int, manager []int, informTime []int) int {

func TestNumOfMinutes(t *testing.T) {
	testCases := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		expect     int
	}{
		{
			1,
			0,
			[]int{-1},
			[]int{0},
			0,
		},
		{
			6,
			2,
			[]int{2, 2, -1, 2, 2, 2},
			[]int{0, 0, 1, 0, 0, 0},
			1,
		},
		{
			7,
			6,
			[]int{1, 2, 3, 4, 5, 6, -1},
			[]int{0, 6, 5, 4, 3, 2, 1},
			21,
		},
		{
			15,
			0,
			[]int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6},
			[]int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := numOfMinutes(testCase.n, testCase.headID, testCase.manager, testCase.informTime)
		assert.Equal(t, testCase.expect, actual)
	}
}
