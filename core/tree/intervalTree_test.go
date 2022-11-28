package tree

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNewIntervalTree(t *testing.T) {
	t.Parallel()

	t.Run("nil config should work", func(t *testing.T) {
		t.Parallel()

		tree := NewIntervalTree(nil)
		assert.False(t, check.IfNil(tree))
		assert.Nil(t, tree.root)
		assert.Equal(t, "", tree.String())
	})
	t.Run("empty config should work", func(t *testing.T) {
		t.Parallel()

		tree := NewIntervalTree(make([]BlocksExceptionInterval, 0))
		assert.False(t, check.IfNil(tree))
		assert.Nil(t, tree.root)
		assert.Equal(t, "", tree.String())
	})
	t.Run("overlapping intervals should work", func(t *testing.T) {
		t.Parallel()

		cfg := []BlocksExceptionInterval{
			{Low: 15, High: 20},
			{Low: 10, High: 30},
			{Low: 17, High: 19},
			{Low: 5, High: 20},
			{Low: 12, High: 15},
			{Low: 30, High: 40},
		}
		//               [15,20]
		//            _____/\____________
		//        [10,30]              [30,40]
		//     _____/\_____         _____/
		// [5,20]       [12,15][17,19]

		tree := NewIntervalTree(cfg)
		assert.False(t, check.IfNil(tree))
		println(tree.String())
		assert.Equal(t, uint64(40), tree.root.max)
		assert.True(t, tree.Contains(16))
		assert.True(t, tree.Contains(11))
		assert.True(t, tree.Contains(35))
		assert.True(t, tree.Contains(6))
		assert.False(t, tree.Contains(4))
		assert.False(t, tree.Contains(41))
	})
	t.Run("non-overlapping intervals should work", func(t *testing.T) {
		t.Parallel()

		cfg := []BlocksExceptionInterval{
			{Low: 1, High: 5},
			{Low: 7, High: 9},
			{Low: 10, High: 11},
			{Low: 12, High: 15},
			{Low: 16, High: 18},
			{Low: 18, High: 20},
			{Low: 22, High: 25},
		}

		//               [12,15]
		//            _____/\____________
		//        [7,9]              [18,20]
		//     _____/\_____         _____/\_____
		// [1,5]       [10,11][16,18]       [22,25]

		tree := NewIntervalTree(cfg)
		assert.False(t, check.IfNil(tree))
		println(tree.String())
		assert.Equal(t, uint64(25), tree.root.max)
		assert.True(t, tree.Contains(12))
		assert.True(t, tree.Contains(16))
		assert.True(t, tree.Contains(8))
		assert.True(t, tree.Contains(19))
		assert.False(t, tree.Contains(0))
		assert.False(t, tree.Contains(21))
	})
	t.Run("multiple offset conflicts should work", func(t *testing.T) {
		t.Parallel()

		cfg := []BlocksExceptionInterval{
			{Low: 22, High: 25},
			{Low: 17, High: 19},
			{Low: 19, High: 20},
			{Low: 14, High: 16},
			{Low: 13, High: 14},
			{Low: 11, High: 13},
			{Low: 20, High: 21},
			{Low: 21, High: 22},
			{Low: 28, High: 30},
			{Low: 26, High: 28},
			{Low: 33, High: 35},
			{Low: 35, High: 37},
			{Low: 35, High: 37},
		}

		//                             [21,22]
		//                          _____/\____________
		//                      [17,19]              [33,35]
		//            ____________/\_____         _____/\____________
		//        [13,14]              [20,21][26,28]              [35,37]
		//     _____/\_____         _____/ _____/\_____         _____/
		// [11,13]       [14,16][19,20][22,25]       [28,30][35,37]

		tree := NewIntervalTree(cfg)
		assert.False(t, check.IfNil(tree))
		println(tree.String())
	})
	t.Run("big tree should not panic", func(t *testing.T) {
		t.Parallel()

		defer func() {
			r := recover()
			if r != nil {
				assert.Fail(t, "should not panic")
			}
		}()

		startPoint := 100
		numNudes := 100
		intervalSize := 3
		cfg := make([]BlocksExceptionInterval, 0, numNudes)
		for i := 0; i < numNudes; i++ {
			cfg = append(cfg, BlocksExceptionInterval{
				Low:  uint64(startPoint + i*intervalSize),
				High: uint64(startPoint + (i+1)*intervalSize),
			})
		}

		tree := NewIntervalTree(cfg)
		assert.False(t, check.IfNil(tree))
		println(tree.String())
	})
}
