package reconstructQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	org [][]int
	res [][]int
}{
	{
		[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
		[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
	},
}

func TestReconstructQueue(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.res, reconstructQueue(tc.org))
	}
}
