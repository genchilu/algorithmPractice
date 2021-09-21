package validIPAddress

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPV4(t *testing.T) {
	testCases := []struct {
		groups []string
		expect bool
	}{
		{
			[]string{"1a", "1b", "1c", "1d"},
			false,
		},
		{
			[]string{"1", "1", "1", ""},
			false,
		},
		{
			[]string{"0", "0", "0", "0"},
			true,
		},
		{
			[]string{"00", "0", "0", "0"},
			false,
		},
		{
			[]string{"192", "168", "1", "1"},
			true,
		},
		{
			[]string{"192", "168", "1", "00"},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := isIPV4(testCase.groups)
		assert.Equal(t, testCase.expect, actual)
	}
}

func TestIPV6(t *testing.T) {
	testCases := []struct {
		groups []string
		expect bool
	}{
		{
			[]string{
				"2001", "0db8", "85a3", "0000", "0000", "8a2e", "0370", "7334",
			},
			true,
		},
		{
			[]string{
				"2001", "db8", "85a3", "0", "0", "8A2E", "0370", "7334",
			},
			true,
		},
		{
			[]string{
				"02001", "0db8", "85a3", "0000", "0000", "8a2e", "0370", "7334",
			},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := isIPV6(testCase.groups)
		assert.Equal(t, testCase.expect, actual)
	}
}
