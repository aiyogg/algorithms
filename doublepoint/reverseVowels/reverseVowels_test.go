package reverseVowels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s   string
	ans string
}{
	{
		"leetcode",
		"leotcede",
	},
}

func TestReverseVowels(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, reverseVowels(tc.s), "test case: %v", tc)
	}
}
