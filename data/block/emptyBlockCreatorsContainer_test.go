package block

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/marshal/factory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyBlockCreatorsContainer(t *testing.T) {
	t.Parallel()

	container := NewEmptyBlockCreatorsContainer()
	require.False(t, check.IfNil(container))
	assert.Equal(t, 0, len(container.blockCreators))
}

func TestEmptyBlockCreatorsContainer_Add(t *testing.T) {
	t.Parallel()

	t.Run("nil block creator should error", func(t *testing.T) {
		container := NewEmptyBlockCreatorsContainer()
		err := container.Add(core.ShardHeaderV1, nil)
		assert.Equal(t, data.ErrNilEmptyBlockCreator, err)
		assert.Equal(t, 0, len(container.blockCreators))
	})
	t.Run("not nil block creator should work", func(t *testing.T) {
		container := NewEmptyBlockCreatorsContainer()
		creator := NewEmptyHeaderCreator()
		err := container.Add(core.ShardHeaderV1, creator)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(container.blockCreators))
	})
}

func TestEmptyBlockCreatorsContainer_Get(t *testing.T) {
	t.Parallel()

	t.Run("missing header creator should error", func(t *testing.T) {
		container := NewEmptyBlockCreatorsContainer()
		creator, err := container.Get(core.ShardHeaderV1)
		assert.Equal(t, data.ErrInvalidHeaderType, err)
		assert.True(t, check.IfNil(creator))
	})
	t.Run("existing header creator should work", func(t *testing.T) {
		container := NewEmptyBlockCreatorsContainer()
		creator := NewEmptyHeaderCreator()
		_ = container.Add(core.ShardHeaderV1, creator)
		recovered, err := container.Get(core.ShardHeaderV1)
		assert.Nil(t, err)
		assert.True(t, recovered == creator) // pointer testing
	})
}

func TestEmptyBlockCreatorsContainer_ConcurrentOperations(t *testing.T) {
	t.Parallel()

	container := NewEmptyBlockCreatorsContainer()
	numOperations := 1000
	wg := &sync.WaitGroup{}
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go func(idx int) {
			time.Sleep(time.Millisecond * 10)
			switch idx {
			case 0:
				_ = container.Add(core.ShardHeaderV1, NewEmptyHeaderCreator())
			case 1:
				_, _ = container.Get(core.ShardHeaderV1)
			default:
				require.Nil(t, fmt.Sprintf("invalid index %d", idx))
			}

			wg.Done()
		}(i % 2)
	}

	wg.Wait()
}

func TestSemiIntegrationUnmarshal(t *testing.T) {
	t.Parallel()

	// setup part
	container := NewEmptyBlockCreatorsContainer()
	err := container.Add(core.ShardHeaderV1, NewEmptyHeaderCreator())
	require.Nil(t, err)
	err = container.Add(core.ShardHeaderV2, NewEmptyHeaderV2Creator())
	require.Nil(t, err)
	err = container.Add(core.MetaHeader, NewEmptyMetaBlockCreator())
	require.Nil(t, err)

	marshaller, _ := factory.NewMarshalizer(factory.GogoProtobuf)
	headerV1 := &Header{Nonce: 1}
	hBytes, _ := marshaller.Marshal(headerV1)

	//usage part
	creator, err := container.Get(core.ShardHeaderV1)
	require.Nil(t, err)

	recoveredHeader, err := GetHeaderFromBytes(marshaller, creator, hBytes)
	assert.Nil(t, err)
	assert.Equal(t, headerV1, recoveredHeader)
	assert.False(t, headerV1 == recoveredHeader) // pointer testing, different objects
}
