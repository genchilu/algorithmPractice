package quicksort


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct{
		input []uint
		expect []uint
	}{
		{nil, nil},
		{[]uint{}, []uint{}},
		{[]uint{1}, []uint{1}},
		{[]uint{1,2}, []uint{1,2}},
		{[]uint{2,1}, []uint{1,2}},
		{[]uint{1,1}, []uint{1,1}},
		{[]uint{1,2,3,4,5}, []uint{1,2,3,4,5}},
		{[]uint{5,4,3,2,1}, []uint{1,2,3,4,5}},
		{[]uint{7,5,3,3,9,6,7}, []uint{3,3,5,6,7,7,9}},
		{[]uint{23, 42, 32, 64, 12, 4}, []uint{4, 12, 23, 32, 42, 64}},
		{[]uint{2,1,2,2,2,2,2}, []uint{1,2,2,2,2,2,2}},
		{[]uint{10,12,9,13,8,14,7}, []uint{7,8,9,10,12,13,14}},
	}

	for _, testCase := range testCases {
		Sort(testCase.input)

		assert.Equal(t, testCase.expect, testCase.input);
	}
}