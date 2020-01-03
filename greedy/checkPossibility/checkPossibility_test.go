package checkPossibility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{4, 2, 3},
		true,
	},
	{
		[]int{4, 2, 1},
		false,
	},
	{
		[]int{3, 4, 2, 3},
		false,
	},
}

func TestCheckPossibility(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, checkPossibility(tc.nums))
	}
}
