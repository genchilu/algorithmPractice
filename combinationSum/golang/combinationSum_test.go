package combinationSum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func combinationSum(candidates []int, target int) [][]int {

func Test1(t *testing.T) {

	testCases := []struct {
		candidates []int
		target     int
		expect     [][]int
	}{
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := combinationSum(testCase.candidates, testCase.target)

		for i := range actual {
			sort.Ints(actual[i])
		}

		assert.Contains(t, testCase.expect, actual)
	}
}
