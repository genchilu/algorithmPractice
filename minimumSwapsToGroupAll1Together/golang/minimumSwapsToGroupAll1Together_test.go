package minimumSwapsToGroupAll1Together

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minSwaps(data []int) int {
func TestDev1(t *testing.T) {
	data := []int{1, 0, 1, 0, 1}
	expect := 1

	actual := minSwaps(data)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	data := []int{0, 0, 0, 1, 0}
	expect := 0

	actual := minSwaps(data)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	data := []int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}
	expect := 3

	actual := minSwaps(data)

	assert.Equal(t, expect, actual)
}
