package uniqueBinarySearchTrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		num    int
		expect int
	}{
		{
			1,
			1,
		},
		{
			3,
			5,
		},
	}

	for _, testCaste := range testCases {
		actual := numTrees(testCaste.num)
		assert.Equal(t, testCaste.expect, actual)
	}
}
