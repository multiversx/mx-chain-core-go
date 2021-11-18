package sliceUtil

// TrimSliceSliceByte creates a copy of the provided slice without the excess capacity
func TrimSliceSliceByte(in [][]byte) [][]byte {
	if len(in) == 0 {
		return [][]byte{}
	}
	ret := make([][]byte, len(in))
	copy(ret, in)
	return ret
}

// IsIndexSetInBitmap - checks if a bit is set(1) in the given bitmap
func IsIndexSetInBitmap(index uint32, bitmap []byte) bool {
	indexOutOfBounds := index >= uint32(len(bitmap))*8
	if indexOutOfBounds {
		return false
	}

	bytePos := index / 8
	byteInMap := bitmap[bytePos]
	bitPos := index % 8
	mask := uint8(1 << bitPos)
	return (byteInMap & mask) != 0
}

func SetIndexInBitmap(index uint32, bitmap []byte) {
	indexOutOfBounds := index >= uint32(len(bitmap))*8
	if indexOutOfBounds {
		return
	}

	bytePos := index / 8
	bitPos := index % 8
	mask := uint8(1 << bitPos)
	bitmap[bytePos] |= mask
}
