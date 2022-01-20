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
	if isIndexOutOfBounds(index, bitmap) {
		return false
	}

	byteInMap := getByteAtIndex(index, bitmap)
	mask := calcByteMask(index)
	return (*byteInMap & mask) != 0
}

func SetIndexInBitmap(index uint32, bitmap []byte) {
	if isIndexOutOfBounds(index, bitmap) {
		return
	}

	byteInMap := getByteAtIndex(index, bitmap)
	mask := calcByteMask(index)
	*byteInMap |= mask
}

// Could panic. Do not call this unless you check before if isIndexOutOfBounds == false
func getByteAtIndex(index uint32, bitmap []byte) *byte {
	bytePos := index / 8
	return &bitmap[bytePos]
}

func calcByteMask(index uint32) byte {
	bitPos := index % 8
	return 1 << bitPos
}

func isIndexOutOfBounds(index uint32, bitmap []byte) bool {
	return index >= uint32(len(bitmap))*8
}
