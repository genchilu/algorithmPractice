package cuttingRibbons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxLength(ribbons []int, k int) int {
func TestDev1(t *testing.T) {
	ribbons := []int{9, 7, 5}
	k := 3
	expect := 5

	actual := maxLength(ribbons, k)
	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	ribbons := []int{7, 5, 9}
	k := 4
	expect := 4

	actual := maxLength(ribbons, k)
	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	ribbons := []int{5, 7, 9}
	k := 22
	expect := 0

	actual := maxLength(ribbons, k)
	assert.Equal(t, expect, actual)
}

func TestDev4(t *testing.T) {
	ribbons := []int{100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 1, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000}
	k := 49
	expect := 100000

	actual := maxLength(ribbons, k)
	assert.Equal(t, expect, actual)
}

func TestDev5(t *testing.T) {
	ribbons := []int{100000, 1, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000}
	k := 50
	expect := 50000

	actual := maxLength(ribbons, k)
	assert.Equal(t, expect, actual)
}
