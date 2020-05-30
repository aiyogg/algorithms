package sqrt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	x   int
	ans int
}{
	{
		4,
		2,
	},
	{
		8,
		2,
	},
}

func TestMySqrt(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, mySqrt(tc.x))
	}
}
