package quickSort

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

var tcs = []struct {
	origin []int
	result []int
}{
	{
		[]int{2, 8, 7, 1, 3, 5, 6, 4},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
	},
}

func TestQuickSort(t *testing.T) {
	// ast := assert.New(t)
	for _, tc := range tcs {
		t.Logf("~~%v~~", tc)
		quickSort(tc.origin, 0, len(tc.origin))
	}
}
