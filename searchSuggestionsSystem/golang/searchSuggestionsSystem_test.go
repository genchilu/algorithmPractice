package searchSuggestionsSystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggestedProducts(t *testing.T) {
	testCases := []struct {
		products   []string
		searchWord string
		expect     [][]string
	}{
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{
				[]string{"mobile", "moneypot", "monitor"},
				[]string{"mobile", "moneypot", "monitor"},
				[]string{"mouse", "mousepad"},
				[]string{"mouse", "mousepad"},
				[]string{"mouse", "mousepad"},
			},
		},
		{
			[]string{"havana"},
			"havana",
			[][]string{
				[]string{"havana"},
				[]string{"havana"},
				[]string{"havana"},
				[]string{"havana"},
				[]string{"havana"},
				[]string{"havana"},
			},
		},
		{
			[]string{"bags", "baggage", "banner", "box", "cloths"},
			"bags",
			[][]string{
				[]string{"baggage", "bags", "banner"},
				[]string{"baggage", "bags", "banner"},
				[]string{"baggage", "bags"},
				[]string{"bags"},
			},
		},
		{
			[]string{"havana"},
			"tatiana",
			[][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{},
			},
		},
	}

	for _, testCase := range testCases {
		actual := suggestedProducts(testCase.products, testCase.searchWord)
		assert.Equal(t, testCase.expect, actual)
	}
}
