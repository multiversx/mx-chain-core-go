package websocketOutportDriver_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubSlice(t *testing.T) {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, []int{0, 1, 2, 3}, sl[:4])
	require.Equal(t, []int{4, 5, 6, 7, 8, 9}, sl[4:])
}

func TestUint64Overflow(t *testing.T) {
	value := uint64(math.MaxUint64)
	value += 1
	require.Equal(t, uint64(0), value)
}

func TestMapSaveEmptyString(t *testing.T) {
	key := ""
	val := "value"

	mp := map[string]string{}

	mp[key] = val

	fmt.Println(mp)
}

func TestRemoveFromInnerMap(t *testing.T) {
	mp := make(map[string]map[uint64]struct{})
	mp["1"] = make(map[uint64]struct{})

	mp["1"][0] = struct{}{}
	mp["1"][2] = struct{}{}
	require.Equal(t, 1, len(mp))
	require.Equal(t, 2, len(mp["1"]))

	oneVals := mp["1"]
	delete(oneVals, 2)
	require.Equal(t, 1, len(mp["1"]))
}
