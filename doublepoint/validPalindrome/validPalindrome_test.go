package validPalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s      string
	except bool
}{
	{
		"aba",
		true,
	},
	{
		"abca",
		true,
	},
	{
		"abxca",
		false,
	},
}

func TestValidPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.except, validPalindrome(tc.s), "test case: ", tc)
		ast.Equal(tc.except, validPalindromeRecursive(tc.s), "test case: ", tc)
	}
}
