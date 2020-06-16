package diffWaysToCompute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	input string
	ans   []int
}{
	{
		"2-1-1",
		[]int{0, 2},
	},
	{
		"2*3-4*5",
		[]int{-34, -14, -10, -10, 10},
	},
}

func TestDiffWaysToCompute(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.ElementsMatch(tc.ans, diffWaysToCompute(tc.input))
	}
}
