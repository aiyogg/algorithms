package canPlaceFlowers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	flowerbeds []int
	n          int
	res        bool
}{
	{
		[]int{1, 0, 0, 0, 1},
		1,
		true,
	},
	{
		[]int{1, 0, 0, 0, 1},
		2,
		false,
	},
}

func TestCanPlaceFlowers(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.res, canPlaceFlowers(tc.flowerbeds, tc.n))
	}
}
