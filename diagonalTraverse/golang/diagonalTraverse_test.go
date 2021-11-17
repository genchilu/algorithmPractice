package diagonalTraverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findDiagonalOrder(mat [][]int) []int {
func TestDev1(t *testing.T) {
	mat := [][]int{{1}}
	expect := []int{1}

	actual := findDiagonalOrder(mat)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}

	actual := findDiagonalOrder(mat)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	mat := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := []int{1, 2, 3, 4}

	actual := findDiagonalOrder(mat)

	assert.Equal(t, expect, actual)
}
