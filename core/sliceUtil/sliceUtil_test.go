package sliceUtil_test

import (
	"strconv"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/sliceUtil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrimSliceSliceByte_EmptyInputShouldDoNothing(t *testing.T) {
	t.Parallel()

	input := make([][]byte, 0)
	res := sliceUtil.TrimSliceSliceByte(input)

	assert.Equal(t, input, res)
}

func TestTrimSliceSliceByte_ShouldDecreaseCapacity(t *testing.T) {
	t.Parallel()

	input := make([][]byte, 0, 5)
	input = append(input, []byte("el1"))
	input = append(input, []byte("el2"))

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 5, cap(input))

	// after calling the trim func, the capacity should be equal to the len

	input = sliceUtil.TrimSliceSliceByte(input)
	assert.Equal(t, 2, len(input))
	assert.Equal(t, 2, cap(input))
}

func TestTrimSliceSliceByte_SliceAlreadyOkShouldDoNothing(t *testing.T) {
	t.Parallel()

	input := make([][]byte, 0, 2)
	input = append(input, []byte("el1"))
	input = append(input, []byte("el2"))

	assert.Equal(t, 2, len(input))
	assert.Equal(t, 2, cap(input))

	// after calling the trim func, the capacity should be equal to the len

	input = sliceUtil.TrimSliceSliceByte(input)
	assert.Equal(t, 2, len(input))
	assert.Equal(t, 2, cap(input))
}

func TestIsIndexSetInBitmap(t *testing.T) {
	byte1Map, _ := strconv.ParseInt("11001101", 2, 9)
	byte2Map, _ := strconv.ParseInt("00000101", 2, 9)
	bitmap := []byte{byte(byte1Map), byte(byte2Map)}

	//Byte 1
	require.True(t, sliceUtil.IsIndexSetInBitmap(0, bitmap))
	require.False(t, sliceUtil.IsIndexSetInBitmap(1, bitmap))
	require.True(t, sliceUtil.IsIndexSetInBitmap(2, bitmap))
	require.True(t, sliceUtil.IsIndexSetInBitmap(3, bitmap))
	require.False(t, sliceUtil.IsIndexSetInBitmap(4, bitmap))
	require.False(t, sliceUtil.IsIndexSetInBitmap(5, bitmap))
	require.True(t, sliceUtil.IsIndexSetInBitmap(6, bitmap))
	require.True(t, sliceUtil.IsIndexSetInBitmap(7, bitmap))
	// Byte 2
	require.True(t, sliceUtil.IsIndexSetInBitmap(8, bitmap))
	require.False(t, sliceUtil.IsIndexSetInBitmap(9, bitmap))
	require.True(t, sliceUtil.IsIndexSetInBitmap(10, bitmap))

	for i := uint32(11); i <= 100; i++ {
		require.False(t, sliceUtil.IsIndexSetInBitmap(i, bitmap))
	}
}
