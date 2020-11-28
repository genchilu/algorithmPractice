package longestLineOfConsecutiveOneInMatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//longestLine(M [][]int) int {
func TestLongestLine(t *testing.T) {
	testCases := []struct {
		M      [][]int
		expect int
	}{
		// {
		// 	[][]int{
		// 		[]int{0, 1, 1, 0},
		// 		[]int{0, 1, 1, 0},
		// 		[]int{0, 0, 0, 1},
		// 	},
		// 	3,
		// },
		{
			[][]int{
				[]int{0, 0},
				[]int{1, 1},
			},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := longestLine(testCase.M)
		assert.Equal(t, testCase.expect, actual)
	}
}
