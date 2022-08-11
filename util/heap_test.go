package util_test

import (
	"testing"

	"github.com/musenwill/algo/util"
	"github.com/stretchr/testify/require"
)

func TestHeap(t *testing.T) {
	{
		hp := util.NewHeap(0, func(i, j int) bool {
			return i < j
		})
		require.Equal(t, 0, hp.Size())

		hp.BatchPush(4, 6, 9, 2, 6, 1, 7, 8, 3)
		require.Equal(t, 6, hp.Size())
		require.Equal(t, []int{4, 6, 8, 6, 7, 9}, hp.Gets())
	}
	{
		hp := util.NewHeap(0, func(i, j float64) bool {
			return i > j
		})
		require.Equal(t, 0, hp.Size())

		hp.BatchPush(7, 6, 9, 2, 6, 1, 4, 8, 3)
		require.Equal(t, 7, hp.Size())
		require.Equal(t, []float64{7, 6, 4, 6, 1, 2, 3}, hp.Gets())
	}
}
