package groupShiftedStrings

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	//"fmt"
)

func TestGroupStrings(t *testing.T) {

	testCases := []struct{
		strings []string
		expect [][]string
	} {
		{
			[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			[][]string{[]string{"abc","bcd","xyz"}, []string{"az","ba"}, []string{"acef"}, []string{"a","z"}},
		},
	}

	for _, testCase := range testCases {
		actualResult := groupStrings(testCase.strings)
		assert.Equal(t, len(testCase.expect), len(actualResult))

		for _, v := range actualResult {
			found := false
			for _, e := range testCase.expect {
				if (reflect.DeepEqual(v, e)) {
					found = true
					break
				}
			}
			assert.Equal(t, true, found)
		}
	}

}