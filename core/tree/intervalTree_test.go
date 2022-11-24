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
			{
				Low:  15,
				High: 20,
			},
			{
				Low:  10,
				High: 30,
			},
			{
				Low:  17,
				High: 19,
			},
			{
				Low:  5,
				High: 20,
			},
			{
				Low:  12,
				High: 15,
			},
			{
				Low:  30,
				High: 40,
			},
		}
		//               [15,20]
		//            _____/\____
		//        [10,30]       [17,19]
		//     _____/\____         \____
		// [5,20]       [12,15]       [30,40]

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
			{
				Low:  1,
				High: 5,
			},
			{
				Low:  7,
				High: 9,
			},
			{
				Low:  10,
				High: 11,
			},
			{
				Low:  12,
				High: 15,
			},
			{
				Low:  16,
				High: 18,
			},
			{
				Low:  18,
				High: 20,
			},
			{
				Low:  22,
				High: 25,
			},
		}

		//				   [12,15]
		//				_____/\____
		//		    [7,9]       [16,18]
		//		_____/\____         \____
		//	[1,5]       [10,11]       [18,20]
		//		                           \____
		//		                               [22,25]

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

	t.Run("offset conflict should work", func(t *testing.T) {
		t.Parallel()

		cfg := []BlocksExceptionInterval{
			{
				Low:  1,
				High: 5,
			},
			{
				Low:  7,
				High: 9,
			},
			{
				Low:  10,
				High: 11,
			},
			{
				Low:  12,
				High: 15,
			},
			{
				Low:  16,
				High: 18,
			},
			{
				Low:  18,
				High: 20,
			},
			{
				Low:  18,
				High: 21,
			},
			{
				Low:  23,
				High: 25,
			},
		}

		//               [16,18]
		//            _____/\___________
		//        [7,9]              [18,21]
		//     _____/\____          _____/\____
		// [1,5]       [10,11][18,20]       [23,25]
		//                  \____
		//                      [12,15]

		tree := NewIntervalTree(cfg)
		assert.False(t, check.IfNil(tree))
		println(tree.String())
		assert.Equal(t, uint64(25), tree.root.max)
		assert.True(t, tree.Contains(12))
		assert.True(t, tree.Contains(16))
		assert.True(t, tree.Contains(8))
		assert.True(t, tree.Contains(19))
		assert.False(t, tree.Contains(0))
		assert.False(t, tree.Contains(22))
	})
}
