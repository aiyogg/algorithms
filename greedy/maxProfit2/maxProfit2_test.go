package maxProfit2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	prices []int
	res    int
}{
	{
		[]int{7, 1, 5, 3, 6, 4},
		7,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
	},
}

func TestMaxProfit2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.res, maxProfit2(tc.prices))
	}
}
