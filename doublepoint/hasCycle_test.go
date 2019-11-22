package main

import (
	"algorithms/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	pos  int
	ans  bool
}{
	{
		[]int{3, 2, 0, -4},
		1,
		true,
	},
}

func TestHasCycle(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		head := common.Nums2ListWithCycle(tc.nums, tc.pos)
		ast.Equal(tc.ans, hasCycle(head), "input: %v", tc)
	}
}

func BenchmarkHasCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			head := common.Nums2ListWithCycle(tc.nums, tc.pos)
			hasCycle(head)
		}
	}
}
