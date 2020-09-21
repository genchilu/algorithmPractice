package accountsMerge

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)


func TestAccountsMerge(t *testing.T) {
	testCases := []struct {
		accounts [][]string
		expect [][]string
	} {
		{
			[][]string{
				[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
				[]string{"John", "johnnybravo@mail.com"},
				[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				[]string{"Mary", "mary@mail.com"},
			},
			[][]string{
				[]string{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				[]string{"John", "johnnybravo@mail.com"},
				[]string{"Mary", "mary@mail.com"},
			},
		},
		{
			[][]string{
				[]string{"Alex","Alex5@m.co","Alex4@m.co","Alex0@m.co"},
				[]string{"Ethan","Ethan3@m.co","Ethan3@m.co","Ethan0@m.co"},
				[]string{"Kevin","Kevin4@m.co","Kevin2@m.co","Kevin2@m.co"},
				[]string{"Gabe","Gabe0@m.co","Gabe3@m.co","Gabe2@m.co"},
				[]string{"Gabe","Gabe3@m.co","Gabe4@m.co","Gabe2@m.co"},
			},
			[][]string{
				[]string{"Alex","Alex0@m.co","Alex4@m.co","Alex5@m.co"},
				[]string{"Ethan","Ethan0@m.co","Ethan3@m.co"},
				[]string{"Kevin","Kevin2@m.co","Kevin4@m.co"},
				[]string{"Gabe","Gabe0@m.co","Gabe2@m.co","Gabe3@m.co","Gabe4@m.co"},
			},
		},
		{
			[][]string{
				[]string{"David","David0@m.co","David1@m.co"},
				[]string{"David","David3@m.co","David4@m.co"},
				[]string{"David","David4@m.co","David5@m.co"},
				[]string{"David","David2@m.co","David3@m.co"},
				[]string{"David","David1@m.co","David2@m.co"},
			},
			[][]string{
				[]string{"David","David0@m.co","David1@m.co","David2@m.co","David3@m.co","David4@m.co","David5@m.co"},
			},
		},
	}

	for _, testCase := range testCases {
		actual := accountsMerge(testCase.accounts)
		
		assert.Equal(t, len(testCase.expect), len(actual))
		equalCount := 0
		for _, a1 := range actual {
			for _, a2 := range testCase.expect {
				if reflect.DeepEqual(a1, a2) {
					equalCount++
					break
				}
			}
		}
		assert.Equal(t, len(testCase.expect), equalCount)
	}

}