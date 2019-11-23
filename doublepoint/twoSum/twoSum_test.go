package twoSum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums   []int
	sum    int
	expect []int
}{
	{
		[]int{2, 7, 11, 15},
		18,
		[]int{1, 2},
	},
}

func TestTwoSum(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.expect, twoSum(tc.nums, tc.sum))
	}
}
