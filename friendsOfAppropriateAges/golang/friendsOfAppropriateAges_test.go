package friendsOfAppropriateAges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumFriendRequests(t *testing.T) {
	testCases := []struct {
		ages   []int
		expect int
	}{
		{
			[]int{16, 16},
			2,
		},
		{
			[]int{16, 17, 18},
			2,
		},
		{
			[]int{20, 30, 100, 110, 120},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := numFriendRequests(testCase.ages)
		assert.Equal(t, testCase.expect, actual)
	}
}
