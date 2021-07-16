package shortestPathToGetFood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func getFood(grid [][]byte) int
func TestGetFood(t *testing.T) {
	testCases := []struct {
		grid   [][]byte
		expect int
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'O', 'O', 'O', 'X'},
				{'X', 'O', 'O', '#', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X'},
			},
			3,
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'X', 'O', 'X'},
				{'X', 'O', 'X', '#', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
			-1,
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
			},
			6,
		},
		{
			[][]byte{
				{'O', '*'},
				{'#', 'O'},
			},
			2,
		},
		{
			[][]byte{
				{'X', '*'},
				{'#', 'X'},
			},
			-1,
		},
	}

	for _, testCase := range testCases {
		actual := getFood(testCase.grid)
		assert.Equal(t, testCase.expect, actual)
	}
}
