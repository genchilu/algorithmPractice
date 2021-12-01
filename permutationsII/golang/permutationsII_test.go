package permutationsII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func permuteUnique(nums []int) [][]int {
func TestDev(t *testing.T) {
	nums := []int{1, 2, 3}

	expect := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	actual := permuteUnique(nums)

	assert.Equal(t, expect, actual)
}
