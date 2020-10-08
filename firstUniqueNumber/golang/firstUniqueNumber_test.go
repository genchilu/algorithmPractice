package firstUniqueNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqueNumber1(t *testing.T) {
	firstUnique := Constructor([]int{2, 3, 5})
	actual := firstUnique.ShowFirstUnique()
	assert.Equal(t, 2, actual)

	testCases := []struct {
		value  int
		expect int
	}{
		{
			5,
			2,
		},
		{
			2,
			3,
		},
		{
			3,
			-1,
		},
	}

	for _, testCase := range testCases {
		firstUnique.Add(testCase.value)
		actual = firstUnique.ShowFirstUnique()
		assert.Equal(t, testCase.expect, actual)
	}
}

func TestFirstUniqueNumber2(t *testing.T) {
	firstUnique := Constructor([]int{7, 7, 7, 7, 7})
	actual := firstUnique.ShowFirstUnique()
	assert.Equal(t, -1, actual)

	firstUnique.Add(7)
	firstUnique.Add(3)
	firstUnique.Add(3)
	firstUnique.Add(7)
	firstUnique.Add(17)

	actual = firstUnique.ShowFirstUnique()
	assert.Equal(t, 17, actual)
}

func TestFirstUniqueNumber3(t *testing.T) {
	firstUnique := Constructor([]int{233})
	actual := firstUnique.ShowFirstUnique()
	assert.Equal(t, 233, actual)

	firstUnique.Add(11)

	actual = firstUnique.ShowFirstUnique()
	assert.Equal(t, 233, actual)
}
