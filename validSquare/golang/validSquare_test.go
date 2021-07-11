package validSquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool
func TestValidSquare(t *testing.T) {
	testCases := []struct {
		p1     []int
		p2     []int
		p3     []int
		p4     []int
		expect bool
	}{
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{1, 0},
			[]int{0, 1},
			true,
		},
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{1, 0},
			[]int{0, 12},
			false,
		},
		{
			[]int{1, 0},
			[]int{-1, 0},
			[]int{0, 1},
			[]int{0, -1},
			true,
		},
		{
			[]int{0, 0},
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, 1},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := validSquare(testCase.p1, testCase.p2, testCase.p3, testCase.p4)
		assert.Equal(t, testCase.expect, actual)
	}
}
