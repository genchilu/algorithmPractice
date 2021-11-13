package minimumMovesToEqualArrayElementsII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minMoves2(nums []int) int {
func TestDev1(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 2

	actual := minMoves2(nums)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	nums := []int{1, 10, 2, 9}
	expect := 16

	actual := minMoves2(nums)

	assert.Equal(t, expect, actual)
}
