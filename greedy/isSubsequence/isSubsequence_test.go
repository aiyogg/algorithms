package isSubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s   string
	t   string
	ans bool
}{
	{
		"abc",
		"ahbgdc",
		true,
	},
	{
		"axc",
		"ahbgdc",
		false,
	},
}

func TestIsSubsequence(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isSubsequence(tc.s, tc.t))
	}
}
