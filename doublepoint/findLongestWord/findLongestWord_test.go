package findlongestword

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s      string
	d      []string
	output string
}{
	{
		"abpcplea",
		[]string{"ale", "apple", "monkey", "plea"},
		"apple",
	},
	{
		"abpcplea",
		[]string{"a", "b", "c"},
		"a",
	},
}

func TestFindLongestWord(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.output, findLongestWord(tc.s, tc.d))
	}
}
