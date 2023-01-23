package block

import (
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/stretchr/testify/require"
)

func shouldNotPanic(t *testing.T) {
	if r := recover(); r != nil {
		require.Fail(t, "should not have panicked")
	}
}

func createDefaultShardTriggerRegistry() *ShardTriggerRegistry {
	return &ShardTriggerRegistry{
		IsEpochStart:                true,
		NewEpochHeaderReceived:      true,
		Epoch:                       10,
		MetaEpoch:                   11,
		CurrentRoundIndex:           10000,
		EpochStartRound:             10000,
		EpochFinalityAttestingRound: 10002,
		EpochMetaBlockHash:          []byte("metaBlockHash"),
		EpochStartShardHeader:       &Header{},
	}
}

func createDefaultShardTriggerRegistryV2() *ShardTriggerRegistryV2 {
	return &ShardTriggerRegistryV2{
		EpochStartShardHeader:       &HeaderV2{},
		IsEpochStart:                true,
		NewEpochHeaderReceived:      true,
		Epoch:                       10,
		MetaEpoch:                   11,
		CurrentRoundIndex:           10000,
		EpochStartRound:             10000,
		EpochFinalityAttestingRound: 10002,
		EpochMetaBlockHash:          []byte("metaBlockHash"),
	}
}

func TestShardTriggerRegistry_GetEpochStartHeaderHandlerNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	epochStartHeaderHandler := str.GetEpochStartHeaderHandler()
	require.Nil(t, epochStartHeaderHandler)
}

func TestShardTriggerRegistry_GetEpochStartHeaderHandlerOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	header := &Header{Epoch: 15}
	str.EpochStartShardHeader = header
	epochStartHeaderHandler := str.GetEpochStartHeaderHandler()
	require.Equal(t, header, epochStartHeaderHandler)
}

func TestShardTriggerRegistry_SetIsEpochStartNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetIsEpochStart(true)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetIsEpochStartOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.IsEpochStart = false
	err := str.SetIsEpochStart(true)
	require.Nil(t, err)
	require.Equal(t, true, str.IsEpochStart)
}

func TestShardTriggerRegistry_SetNewEpochHeaderReceivedNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetNewEpochHeaderReceived(true)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetNewEpochHeaderReceivedOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.NewEpochHeaderReceived = false
	err := str.SetNewEpochHeaderReceived(true)
	require.Nil(t, err)
	require.Equal(t, true, str.NewEpochHeaderReceived)
}

func TestShardTriggerRegistry_SetEpochNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetEpoch(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetEpochOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.Epoch = 0
	err := str.SetEpoch(20)
	require.Nil(t, err)
	require.Equal(t, uint32(20), str.Epoch)
}

func TestShardTriggerRegistry_SetMetaEpochNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetMetaEpoch(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetMetaEpochOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.MetaEpoch = 0
	err := str.SetMetaEpoch(20)
	require.Nil(t, err)
	require.Equal(t, uint32(20), str.MetaEpoch)
}

func TestShardTriggerRegistry_SetCurrentRoundIndexNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetCurrentRoundIndex(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetCurrentRoundIndexOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.CurrentRoundIndex = 0
	err := str.SetCurrentRoundIndex(20)
	require.Nil(t, err)
	require.Equal(t, int64(20), str.CurrentRoundIndex)
}

func TestShardTriggerRegistry_SetEpochStartRoundNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetEpochStartRound(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetEpochStartRoundOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.EpochStartRound = 0
	err := str.SetEpochStartRound(20)
	require.Nil(t, err)
	require.Equal(t, uint64(20), str.EpochStartRound)
}

func TestShardTriggerRegistry_SetEpochFinalityAttestingRoundNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetEpochFinalityAttestingRound(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetEpochFinalityAttestingRoundOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.EpochFinalityAttestingRound = 0
	err := str.SetEpochFinalityAttestingRound(20)
	require.Nil(t, err)
	require.Equal(t, uint64(20), str.EpochFinalityAttestingRound)
}

func TestShardTriggerRegistry_SetEpochMetaBlockHashNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetEpochMetaBlockHash([]byte("meta block hash"))
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetEpochMetaBlockHashOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.EpochMetaBlockHash = []byte("hash")
	metaBlockHash := []byte("meta block hash")
	err := str.SetEpochMetaBlockHash(metaBlockHash)
	require.Nil(t, err)
	require.Equal(t, metaBlockHash, str.EpochMetaBlockHash)
}

func TestShardTriggerRegistry_SetEpochStartHeaderHandlerNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistry
	err := str.SetEpochStartHeaderHandler(&Header{})
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistry_SetEpochStartHeaderHandlerNilHeaderToSet(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.EpochStartShardHeader = &Header{
		Epoch: 10,
	}
	setHeader := data.HeaderHandler(nil)
	err := str.SetEpochStartHeaderHandler(setHeader)
	require.Equal(t, data.ErrInvalidTypeAssertion, err)
}

func TestShardTriggerRegistry_SetEpochStartHeaderHandlerOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistry()
	str.EpochStartShardHeader = &Header{
		Epoch: 10,
	}
	setHeader := &Header{
		Epoch: 20,
	}
	err := str.SetEpochStartHeaderHandler(setHeader)
	require.Nil(t, err)
	require.Equal(t, setHeader, str.EpochStartShardHeader)
}

func TestShardTriggerRegistryV2_GetEpochStartHeaderHandlerNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	epochStartHeaderHandler := str.GetEpochStartHeaderHandler()
	require.Nil(t, epochStartHeaderHandler)
}

func TestShardTriggerRegistryV2_GetEpochStartHeaderHandlerOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochStartShardHeader = &HeaderV2{
		Header:            &Header{},
		ScheduledRootHash: []byte("scheduledRootHash"),
	}
	epochStartHeaderHandler := str.GetEpochStartHeaderHandler()
	require.Equal(t, str.EpochStartShardHeader, epochStartHeaderHandler)
}

func TestShardTriggerRegistryV2_SetIsEpochStartNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetIsEpochStart(false)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetIsEpochStartOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.IsEpochStart = false
	err := str.SetIsEpochStart(true)
	require.Nil(t, err)
	require.Equal(t, true, str.IsEpochStart)
}

func TestShardTriggerRegistryV2_SetNewEpochHeaderReceivedNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetNewEpochHeaderReceived(false)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetNewEpochHeaderReceivedOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.NewEpochHeaderReceived = false
	err := str.SetNewEpochHeaderReceived(true)
	require.Nil(t, err)
	require.Equal(t, true, str.NewEpochHeaderReceived)
}

func TestShardTriggerRegistryV2_SetEpochNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetEpoch(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetEpochOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.Epoch = 0
	err := str.SetEpoch(20)
	require.Nil(t, err)
	require.Equal(t, uint32(20), str.Epoch)
}

func TestShardTriggerRegistryV2_SetMetaEpochNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetMetaEpoch(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetMetaEpochOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.Epoch = 0
	err := str.SetMetaEpoch(20)
	require.Nil(t, err)
	require.Equal(t, uint32(20), str.MetaEpoch)
}

func TestShardTriggerRegistryV2_SetCurrentRoundIndexNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetCurrentRoundIndex(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetCurrentRoundIndexOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.CurrentRoundIndex = 0
	err := str.SetCurrentRoundIndex(20)
	require.Nil(t, err)
	require.Equal(t, int64(20), str.CurrentRoundIndex)
}

func TestShardTriggerRegistryV2_SetEpochStartRoundNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetEpochStartRound(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}
func TestShardTriggerRegistryV2_SetEpochStartRoundOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochStartRound = 0
	err := str.SetEpochStartRound(20)
	require.Nil(t, err)
	require.Equal(t, uint64(20), str.EpochStartRound)
}

func TestShardTriggerRegistryV2_SetEpochFinalityAttestingRoundNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetEpochFinalityAttestingRound(20)
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetEpochFinalityAttestingRoundOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochFinalityAttestingRound = 0
	err := str.SetEpochFinalityAttestingRound(20)
	require.Nil(t, err)
	require.Equal(t, uint64(20), str.EpochFinalityAttestingRound)
}

func TestShardTriggerRegistryV2_SetEpochMetaBlockHashNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetEpochMetaBlockHash([]byte("epoch meta block hash"))
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetEpochMetaBlockHashOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochMetaBlockHash = []byte("meta hash")
	setMetaBlockHash := []byte("set meta hash")
	err := str.SetEpochMetaBlockHash(setMetaBlockHash)
	require.Nil(t, err)
	require.Equal(t, setMetaBlockHash, str.EpochMetaBlockHash)
}

func TestShardTriggerRegistryV2_SetEpochStartHeaderHandlerNilShardTriggerRegistry(t *testing.T) {
	t.Parallel()

	defer shouldNotPanic(t)

	var str *ShardTriggerRegistryV2
	err := str.SetEpochStartHeaderHandler(&HeaderV2{})
	require.Equal(t, data.ErrNilPointerReceiver, err)
}

func TestShardTriggerRegistryV2_SetEpochStartHeaderHandlerNilHeaderToSet(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochStartShardHeader = &HeaderV2{
		Header:            &Header{},
		ScheduledRootHash: []byte("scheduled root hash"),
	}
	setHeader := data.HeaderHandler(nil)
	err := str.SetEpochStartHeaderHandler(setHeader)
	require.Equal(t, data.ErrInvalidTypeAssertion, err)
}

func TestShardTriggerRegistryV2_SetEpochStartHeaderHandlerOK(t *testing.T) {
	t.Parallel()

	str := createDefaultShardTriggerRegistryV2()
	str.EpochStartShardHeader = &HeaderV2{
		Header:            &Header{Epoch: 1},
		ScheduledRootHash: []byte("scheduled root hash"),
	}

	setHeader := &HeaderV2{
		Header:            &Header{Epoch: 10},
		ScheduledRootHash: []byte("set scheduled root hash"),
	}

	err := str.SetEpochStartHeaderHandler(setHeader)
	require.Nil(t, err)
	require.Equal(t, setHeader, str.EpochStartShardHeader)
}
