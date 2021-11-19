package capacityToShipPackagesWithinDDays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func shipWithinDays(weights []int, days int) int {
func TestDev1(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	expect := 15

	actual := shipWithinDays(weights, days)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	weights := []int{3, 2, 2, 4, 1, 4}
	days := 3
	expect := 6

	actual := shipWithinDays(weights, days)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	weights := []int{1, 2, 3, 1, 1}
	days := 4
	expect := 3

	actual := shipWithinDays(weights, days)

	assert.Equal(t, expect, actual)
}
