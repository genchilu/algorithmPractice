package dailyTemperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//dailyTemperatures(temperatures []int) []int
func TestDev1(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expect := []int{1, 1, 4, 2, 1, 1, 0, 0}

	actual := dailyTemperatures(temperatures)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	temperatures := []int{30, 40, 50, 60}
	expect := []int{1, 1, 1, 0}

	actual := dailyTemperatures(temperatures)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	temperatures := []int{30, 60, 90}
	expect := []int{1, 1, 0}

	actual := dailyTemperatures(temperatures)

	assert.Equal(t, expect, actual)
}
