package analyzeUserWebsiteVisitPattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostVisitedPattern(t *testing.T) {

	testCases := []struct {
		username  []string
		timestamp []int
		website   []string
		expect    []string
	}{
		{
			[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"},
			[]string{"home", "about", "career"},
		},
		{
			[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"},
			[]string{"home", "about", "career"},
		},
		{
			[]string{"u1", "u1", "u1", "u2", "u2", "u2"},
			[]int{1, 2, 3, 4, 5, 6},
			[]string{"a", "b", "a", "a", "b", "c"},
			[]string{"a", "b", "a"},
		},
		{
			[]string{"zkiikgv", "zkiikgv", "zkiikgv", "zkiikgv"},
			[]int{436363475, 710406388, 386655081, 797150921},
			[]string{"wnaaxbfhxp", "mryxsjc", "oz", "wlarkzzqht"},
			[]string{"oz", "mryxsjc", "wlarkzzqht"},
		},
		{
			[]string{"h", "eiy", "cq", "h", "cq", "txldsscx", "cq", "txldsscx", "h", "cq", "cq"},
			[]int{527896567, 334462937, 517687281, 134127993, 859112386, 159548699, 51100299, 444082139, 926837079, 317455832, 411747930},
			[]string{"hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "yljmntrclw", "hibympufi", "yljmntrclw"},
			[]string{"hibympufi", "hibympufi", "yljmntrclw"},
		},
	}

	for _, testCase := range testCases {
		actual := mostVisitedPattern(testCase.username, testCase.timestamp, testCase.website)
		assert.Equal(t, testCase.expect, actual)
	}
}
