package quicksort


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct{
		input []int
		expect []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1,2}, []int{1,2}},
		{[]int{2,1}, []int{1,2}},
		{[]int{1,1}, []int{1,1}},
		{[]int{1,2,3,4,5}, []int{1,2,3,4,5}},
		{[]int{5,4,3,2,1}, []int{1,2,3,4,5}},
		{[]int{7,5,3,3,9,6,7}, []int{3,3,5,6,7,7,9}},
		{[]int{23, 42, 32, 64, 12, 4}, []int{4, 12, 23, 32, 42, 64}},
		{[]int{2,1,2,2,2,2,2}, []int{1,2,2,2,2,2,2}},
		{[]int{10,12,9,13,8,14,7}, []int{7,8,9,10,12,13,14}},
	}

	for _, testCase := range testCases {
		Sort(testCase.input)

		assert.Equal(t, testCase.expect, testCase.input);
	}
}