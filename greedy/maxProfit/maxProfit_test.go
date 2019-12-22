package maxProfit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	prices []int
	res    int
}{
	{
		[]int{7, 2, 1, 5, 3, 6, 4},
		5,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		5,
	},
}

func TestMaxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.res, maxProfit(tc.prices))
	}
}
