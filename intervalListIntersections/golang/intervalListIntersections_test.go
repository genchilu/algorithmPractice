package intervalListIntersections

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestIntervalIntersection(t *testing.T) {
	testCases := []struct {
		A      [][]int
		B      [][]int
		expect [][]int
	}{
		{
			[][]int{
				[]int{0, 2},
				[]int{5, 10},
				[]int{13, 23},
				[]int{24, 25},
			},
			[][]int{
				[]int{1, 5},
				[]int{8, 12},
				[]int{15, 24},
				[]int{25, 26},
			},
			[][]int{
				[]int{1, 2},
				[]int{5, 5},
				[]int{8, 10},
				[]int{15, 23},
				[]int{24, 24},
				[]int{25, 25},
			},
		},
	}

	for _, testCase := range testCases {
		actual := intervalIntersection(testCase.A, testCase.B)
		assert.Equal(t, testCase.expect, actual)
	}
}
