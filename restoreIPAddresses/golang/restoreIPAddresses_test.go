package restoreIPAddresses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		s      string
		expect []string
	}{
		{
			"25525511135",
			[]string{
				"255.255.11.135",
				"255.255.111.35",
			},
		},
		{
			"0000",
			[]string{
				"0.0.0.0",
			},
		},
		{
			"1111",
			[]string{
				"1.1.1.1",
			},
		},
		{
			"010010",
			[]string{
				"0.10.0.10",
				"0.100.1.0",
			},
		},
		{
			"101023",
			[]string{
				"1.0.10.23",
				"1.0.102.3",
				"10.1.0.23",
				"10.10.2.3",
				"101.0.2.3",
			},
		},
	}

	for _, testCase := range testCases {
		actual := restoreIpAddresses(testCase.s)
		assert.Equal(t, len(testCase.expect), len(actual))

		for _, s := range actual {
			assert.Contains(t, testCase.expect, s)
		}
	}
}
