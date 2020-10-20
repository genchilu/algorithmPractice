package asteroidCollision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	testCases := []struct {
		asteroids []int
		expect    []int
	}{
		// {
		// 	[]int{5, 10, -5},
		// 	[]int{5, 10},
		// },
		// {
		// 	[]int{8, -8},
		// 	[]int{},
		// },
		// {
		// 	[]int{8, -8, 5, 10},
		// 	[]int{5, 10},
		// },
		// {
		// 	[]int{5, 10, 8, -8},
		// 	[]int{5, 10},
		// },
		// {
		// 	[]int{5, 8, -8, 10},
		// 	[]int{5, 10},
		// },
		// {
		// 	[]int{10, 2, -5},
		// 	[]int{10},
		// },
		// {
		// 	[]int{-2, -1, 1, 2},
		// 	[]int{-2, -1, 1, 2},
		// },
		{
			[]int{1, -2, -2, -2},
			[]int{-2, -2, -2},
		},
		{
			[]int{1, 1, -1, -2},
			[]int{-2},
		},
	}

	for _, testCase := range testCases {
		actual := asteroidCollision(testCase.asteroids)
		assert.Equal(t, testCase.expect, actual)
	}
}
