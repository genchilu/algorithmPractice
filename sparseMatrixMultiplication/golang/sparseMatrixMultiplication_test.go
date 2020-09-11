package sparseMatrixMultiplication

import (
	"testing"
	"github.com/stretchr/testify/assert"
	// "math"
	//"fmt"
)

func TestMultiply(t *testing.T) {
	testCases := []struct{
		A [][]int
		B [][]int
		expect [][]int
	}{
		{
			[][]int{[]int{1, 0, 0}, []int{-1, 0, 3}},
			[][]int{[]int{7,0,0}, []int{0,0,0}, []int{0,0,1}},
			[][]int{[]int{7,0,0}, []int{-7,0,3}},
		},
	}

	for _, testCase := range testCases {
		actual := multiply(testCase.A, testCase.B)
		assert.Equal(t, testCase.expect, actual)
	}
}
