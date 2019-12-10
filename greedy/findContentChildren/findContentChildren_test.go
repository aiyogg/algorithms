package findContentChildren

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	grid []int
	size []int
	ans  int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 1},
		1,
	},
	{
		[]int{1, 2},
		[]int{1, 2, 3},
		2,
	},
}

func TestFindContentChildren(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findContentChildren(tc.grid, tc.size))
	}
}
