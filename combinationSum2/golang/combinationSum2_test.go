package combinationSum2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func combinationSum2(candidates []int, target int) [][]int {
//1 1 2 5 6 7 10
func TestDev(t *testing.T) {
	candidates := []int{3, 1, 3, 5, 1, 1}
	target := 8

	expect := [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}

	actual := combinationSum2(candidates, target)

	fmt.Printf("111 %v\n", actual)
	assert.Equal(t, expect, actual)
}
