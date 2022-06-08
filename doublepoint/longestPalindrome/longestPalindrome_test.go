package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	s      string
	output string
}{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"ac", "a"},
}

func TestLongestPalindrome1(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.output, longestPalindrome1(tc.s), "input:%v", tc)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.output, longestPalindrome2(tc.s), "input:%v", tc)
	}
}
