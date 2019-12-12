package eraseOverlapIntervals

import (
	"algorithms/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	intervals [][]int
	ans       int
}{
	{
		[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		1,
	},

	{
		[][]int{{1, 2}, {1, 2}, {1, 2}},
		2,
	},

	{
		[][]int{{2, 3}},
		0,
	},
}

func TestEraseOverlapIntervals(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, eraseOverlapIntervals(common.Intss2IntervalSlice(tc.intervals)))
	}
}
