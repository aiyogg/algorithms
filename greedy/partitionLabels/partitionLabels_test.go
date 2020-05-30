package partitionLabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	S   string
	ans []int
}{
	{
		"abcdefg",
		[]int{1, 1, 1, 1, 1, 1, 1},
	},
	{
		"ababcbacadefegdehijhklij",
		[]int{9, 7, 8},
	},
}

func TestPartitionLabels(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, partitionLabels(tc.S))
	}
}

func BenchmarkPartitionLabels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			partitionLabels(tc.S)
		}
	}
}
