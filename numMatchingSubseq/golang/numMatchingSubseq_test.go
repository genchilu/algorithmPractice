package numMatchingSubseq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//unc numMatchingSubseq(S string, words []string) int {

func TestNumMatchingSubseq(t *testing.T) {
	testCases := []struct {
		S      string
		word   []string
		expect int
	}{
		// {
		// 	"abcde",
		// 	[]string{"a", "bb", "acd", "ace"},
		// 	3,
		// },
		{
			"rwpddkvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvjubjgloeofnpjqlkdsqvruvabjrikfwronbrdyyjnakstqjac",
			[]string{"wpddkvbnn", "lnagtva", "kvbnnuglnagtvamxkqtwhqgwbqgfbvgkwyuqkdwhzudsxvju", "rwpddkvbnnugln", "gloeofnpjqlkdsqvruvabjrikfwronbrdyyj", "vbgeinupkvgmgxeaaiuiyojmoqkahwvbpwugdainxciedbdkos", "mspuhbykmmumtveoighlcgpcapzczomshiblnvhjzqjlfkpina", "rgmliajkiknongrofpugfgajedxicdhxinzjakwnifvxwlokip", "fhepktaipapyrbylskxddypwmuuxyoivcewzrdwwlrlhqwzikq", "qatithxifaaiwyszlkgoljzkkweqkjjzvymedvclfxwcezqebx"},
			5,
		},
	}

	for _, testCase := range testCases {
		actual := numMatchingSubseq(testCase.S, testCase.word)
		assert.Equal(t, testCase.expect, actual)
	}
}
