package numberOfConnectedComponentsInAnUndirectedGraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func countComponents(n int, edges [][]int) int {
func TestCountComponents(t *testing.T) {
	testCases := []struct {
		n      int
		edges  [][]int
		expect int
	}{
		{
			5,
			[][]int{
				[]int{0, 1},
				[]int{1, 2},
				[]int{3, 4},
			},
			2,
		},
		{
			5,
			[][]int{
				[]int{0, 1},
				[]int{1, 2},
				[]int{2, 3},
				[]int{3, 4},
			},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := countComponents(testCase.n, testCase.edges)
		assert.Equal(t, testCase.expect, actual)
	}
}
