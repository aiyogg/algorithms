package singleelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		2,
	},
	{
		[]int{2, 2, 3, 3, 4, 4, 5},
		5,
	},
}

func TestSingleNonDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, singleNonDuplicate(tc.nums))
	}
}
func BenchmarkSingleNonDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNonDuplicate(tc.nums)
		}
	}
}
