package reconstructItinerary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findItinerary(tickets [][]string) []string {
func TestFindItinerary(t *testing.T) {
	testCases := []struct {
		tickets [][]string
		expect  []string
	}{
		{
			[][]string{
				{"MUC", "LHR"},
				{"JFK", "MUC"},
				{"SFO", "SJC"},
				{"LHR", "SFO"},
			},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			[][]string{
				{"JFK", "SFO"},
				{"JFK", "ATL"},
				{"SFO", "ATL"},
				{"ATL", "JFK"},
				{"ATL", "SFO"},
			},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			[][]string{
				{"JFK", "ATL"},
				{"ATL", "JFK"},
			},
			[]string{"JFK", "ATL", "JFK"},
		},
	}

	for _, testCase := range testCases {
		actual := findItinerary(testCase.tickets)
		assert.Equal(t, testCase.expect, actual)
	}
}
