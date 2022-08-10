package numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenSliceAndLoadSlice(t *testing.T) {
	gen := NewGenerator(8, "")
	bytes := make([]byte, 8*8)
	copy(bytes, gen.GenSlice())
	act := make([]int64, 8)
	copy(act, LoadSlice(bytes))
	require.Equal(t, gen.buf, act)
}
