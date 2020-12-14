package sequenceReconstruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func sequenceReconstruction(org []int, seqs [][]int) bool {

func TestSequenceReconstruction(t *testing.T) {
	testCases := []struct {
		org    []int
		seqs   [][]int
		expect bool
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2},
				[]int{1, 3},
			},
			false,
		},
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2},
			},
			false,
		},
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2},
				[]int{2, 3},
				[]int{1, 3},
			},
			true,
		},
		{
			[]int{4, 1, 5, 2, 6, 3},
			[][]int{
				[]int{5, 2, 6, 3},
				[]int{4, 1, 5, 2},
			},
			true,
		},
		{
			[]int{1},
			[][]int{
				[]int{1, 1},
			},
			false,
		},
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{1, 2},
				[]int{2, 1},
				[]int{1, 3},
				[]int{2, 3},
			},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := sequenceReconstruction(testCase.org, testCase.seqs)
		assert.Equal(t, testCase.expect, actual)
	}
}
