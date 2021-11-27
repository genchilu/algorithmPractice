package threeSumClosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func threeSumClosest(nums []int, target int) int {
func TestDev1(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1

	expect := 2

	actual := threeSumClosest(nums, target)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	nums := []int{0, 0, 0}
	target := 1

	expect := 0

	actual := threeSumClosest(nums, target)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	nums := []int{-111, -111, 3, 6, 7, 16, 17, 18, 19}
	target := 13

	expect := 16

	actual := threeSumClosest(nums, target)

	assert.Equal(t, expect, actual)
}
