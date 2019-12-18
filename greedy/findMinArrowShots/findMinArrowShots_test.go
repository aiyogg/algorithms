package findMinArrowShots

import (
	"algorithms/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	points [][]int
	ans    int
}{
	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		2,
	},

	{
		[][]int{{1, 2}, {1, 2}, {1, 2}},
		1,
	},

	{
		[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		2,
	},
}

func TestFindMinArrowShots(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findMinArrowShots(common.Intss2IntervalSlice(tc.points)))
	}
}
