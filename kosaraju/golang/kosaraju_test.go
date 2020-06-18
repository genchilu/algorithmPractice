package kosaraju

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKosarajuScc(t *testing.T) {
	testCases := []struct{
		edges [][]int
		expectComponments [][]int
	} {
			// case1
			{
				// edge
				[][]int{
					[]int{6,7},
					[]int{7,0}, 
					[]int{7,4},
					[]int{7,5},
					[]int{5,6},
					[]int{5,4},
					[]int{4,3},
					[]int{3,2},
					[]int{3,4},
					[]int{0,1},
					[]int{0,2},
					[]int{0,3},
					[]int{1,0},
					[]int{1,2}},
				// expect componment
				[][]int{
					[]int{5,6,7},
					[]int{3,4},
					[]int{2},
					[]int{0,1}}},
			{
				// edge
				[][]int{
					[]int{1,0},
					[]int{2,1},
					[]int{0,2},
					[]int{0,3},
					[]int{3,4}},
				// expect componment
				[][]int{
					[]int{0,1,2},
					[]int{3},
					[]int{4}}},
			{
				// edge
				[][]int{
					[]int{1,3},
					[]int{2,1},
					[]int{3,2},
					[]int{3,4},
					[]int{4,5},
					[]int{5,6},
					[]int{6,4},
					[]int{7,8},
					[]int{8,10},
					[]int{9,7},
					[]int{9,6},
					[]int{10,9},
					[]int{10,11}},
				// expect componment
				[][]int{
					[]int{1,2,3},
					[]int{4,5,6},
					[]int{7,8,9,10},
					[]int{11}}},
	}
	
	for _, testCase := range testCases {

		actualResult := DoKosarajuScc(testCase.edges)

		// check componment number 
		assert.Equal(t, len(testCase.expectComponments), len(actualResult))
		// check every componment
		for _, expectComponment := range testCase.expectComponments {
			lowestVertex := expectComponment[0]
			assert.Equal(t, len(expectComponment), len(actualResult[lowestVertex]))
			for i, expectVertex := range expectComponment {
				assert.Equal(t, expectVertex, actualResult[lowestVertex][i])
			}
		}
	}

	

}