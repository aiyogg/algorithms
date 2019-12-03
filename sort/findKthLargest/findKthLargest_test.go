package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	k    int
	ans  int
}{
	{
		[]int{3, 2, 1, 5, 6, 4},
		2,
		5,
	},
	{
		[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		4,
		4,
	},
}

func TestHeap(t *testing.T) {
	ast := assert.New(t)

	h := new(highHeap)
	i := 5
	h.Push(i)
	ast.Equal(i, h.Pop(), "Pop() after push(%d)", i)
}

func TestFindKthLargest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("test case: %v \n", tc)
		ast.Equal(tc.ans, findKthLargest(tc.nums, tc.k), "Input: %v", tc)
	}
}

func TestFindKthLargest2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("test case: %v \n", tc)
		ast.Equal(tc.ans, findKthLargest2(tc.nums, tc.k), "Input: %v", tc)
	}
}
