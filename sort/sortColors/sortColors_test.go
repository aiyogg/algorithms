package sortColors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  []int
}{
	{
		[]int{2, 0, 2, 1, 1, 0},
		[]int{0, 0, 1, 1, 2, 2},
	},
	{
		[]int{1, 2, 1, 2, 1, 2},
		[]int{1, 1, 1, 2, 2, 2},
	},
	{
		[]int{0, 1, 2, 0, 1, 2, 0, 1, 2},
		[]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
	},
}

func TestSortColors(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		sortColors(tc.nums)
		ast.Equal(tc.ans, tc.nums)
	}
}
