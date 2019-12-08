package frequencySort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	str string
	ans string
}{
	{
		"tree",
		"eert",
	},
	{
		"cccaaa",
		"aaaccc",
	},
	{
		"afeadedd",
		"dddaaeef",
	},
}

func TestFrequencySort(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, frequencySort(tc.str))
	}
}
