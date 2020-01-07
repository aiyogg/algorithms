package maxSubArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		6,
	},
	{
		[]int{-2},
		-2,
	},
}

func TestMaxSubArray(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSubArray(tc.nums))
	}
}

func TestMaxSubArray2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSubArray2(tc.nums))
	}
}
