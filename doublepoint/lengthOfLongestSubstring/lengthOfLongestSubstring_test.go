package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tsc = []struct {
	s      string
	output int
}{
	{"abcbacbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"", 0},
	{"au", 2},
	{"dvdf", 3},
	{"tmmzuxt", 5},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tsc {
		ast.Equal(tc.output, lengthOfLongestSubstring(tc.s), "input:%v", tc)
	}
}
