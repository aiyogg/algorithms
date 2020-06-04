package findMin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 4, 5, 1, 2},
		1,
	},
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
	},
}

func TestFindMin(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findMin(tc.nums))
	}
}
