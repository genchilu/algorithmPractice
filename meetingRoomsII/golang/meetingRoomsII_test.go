package meetingRoomsII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMeetingRooms(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		expect    int
	}{
		// {
		// 	[][]int{
		// 		[]int{0, 30},
		// 		[]int{5, 10},
		// 		[]int{15, 20},
		// 	},
		// 	2,
		// },
		// {
		// 	[][]int{
		// 		[]int{7, 10},
		// 		[]int{2, 4},
		// 	},
		// 	1,
		// },
		{
			[][]int{
				[]int{13, 15},
				[]int{1, 13},
			},
			1,
		},
	}

	for _, testCase := range testCases {
		actual := minMeetingRooms(testCase.intervals)
		assert.Equal(t, testCase.expect, actual)
	}
}
