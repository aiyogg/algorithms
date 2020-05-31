package nextGreatestLetter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	letters []byte
	target  byte
	ans     byte
}{
	{
		[]byte{'c', 'f', 'g'},
		'a',
		'c',
	},
	{
		[]byte{'a', 'b', 'c'},
		'z',
		'a',
	},
}

func TestNextGreatestLetter(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, nextGreatestLetter(tc.letters, tc.target))
	}
}

func BenchmarkNextGreatestLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextGreatestLetter(tc.letters, tc.target)
		}
	}
}
