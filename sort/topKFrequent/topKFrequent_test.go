package topKFrequent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	k    int
	ans  []int
}{
	{
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	},
	{
		[]int{1},
		1,
		[]int{1},
	},
}

func TestTopKFrequent(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, topKFrequent(tc.nums, tc.k))
	}
}
