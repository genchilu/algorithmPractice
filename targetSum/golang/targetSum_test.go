package targetSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev1(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	expect := 5

	actual := findTargetSumWays(nums, target)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	nums := []int{1}
	target := 1
	expect := 1

	actual := findTargetSumWays(nums, target)

	assert.Equal(t, expect, actual)
}
