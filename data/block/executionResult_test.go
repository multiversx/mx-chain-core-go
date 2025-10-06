package block

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/multiversx/mx-chain-core-go/data"
)

func TestExecutionResult_GetHeaderHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetHeaderHash())
	})

	t.Run("with header hash", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderHash: []byte("hash"),
			},
		}
		assert.Equal(t, []byte("hash"), er.GetHeaderHash())
	})
}

func TestExecutionResult_GetHeaderNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetHeaderNonce())
	})

	t.Run("with nonce", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderNonce: 42,
			},
		}
		assert.Equal(t, uint64(42), er.GetHeaderNonce())
	})
}

func TestExecutionResult_GetHeaderRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetHeaderRound())
	})

	t.Run("with round", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderRound: 10,
			},
		}
		assert.Equal(t, uint64(10), er.GetHeaderRound())
	})
}

func TestExecutionResult_GetHeaderEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Equal(t, uint32(0), er.GetHeaderEpoch())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderEpoch: 10,
			},
		}
		assert.Equal(t, uint32(10), er.GetHeaderEpoch())
	})
}

func TestExecutionResult_GetRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetRootHash())
	})

	t.Run("with root hash", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				RootHash: []byte("root"),
			},
		}
		assert.Equal(t, []byte("root"), er.GetRootHash())
	})
}

func TestExecutionResult_GetMiniBlockHeadersHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetMiniBlockHeadersHandlers())
	})

	t.Run("no miniblocks", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			MiniBlockHeaders: nil,
		}
		assert.Empty(t, er.GetMiniBlockHeadersHandlers())
	})

	t.Run("with miniblocks", func(t *testing.T) {
		t.Parallel()

		mb1 := MiniBlockHeader{Hash: []byte("hash1")}
		mb2 := MiniBlockHeader{Hash: []byte("hash2")}
		er := &ExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{mb1, mb2},
		}

		handlers := er.GetMiniBlockHeadersHandlers()
		assert.Len(t, handlers, 2)
		assert.Equal(t, mb1.Hash, handlers[0].GetHash())
		assert.Equal(t, mb2.Hash, handlers[1].GetHash())
	})
}

func TestExecutionResult_SetMiniBlockHeadersHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		err := er.SetMiniBlockHeadersHandlers(nil)
		assert.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{{Hash: []byte("old")}},
		}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{})
		assert.NoError(t, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with nil slice", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{{Hash: []byte("old")}},
		}
		err := er.SetMiniBlockHeadersHandlers(nil)
		assert.NoError(t, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with invalid mini block header handlers", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{nil})
		assert.Equal(t, data.ErrInvalidTypeAssertion, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with nil mini block header handler", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{(*MiniBlockHeader)(nil)})
		assert.Equal(t, data.ErrNilPointerDereference, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with valid miniblocks", func(t *testing.T) {
		t.Parallel()

		mb1 := &MiniBlockHeader{Hash: []byte("hash1")}
		mb2 := &MiniBlockHeader{Hash: []byte("hash2")}
		handlers := []data.MiniBlockHeaderHandler{mb1, mb2}

		er := &ExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers(handlers)
		assert.NoError(t, err)
		assert.Len(t, er.MiniBlockHeaders, 2)
		assert.Equal(t, mb1.Hash, er.MiniBlockHeaders[0].Hash)
		assert.Equal(t, mb2.Hash, er.MiniBlockHeaders[1].Hash)
	})
}

func TestExecutionResult_GetReceiptsHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetReceiptsHash())
	})

	t.Run("with receipts hash", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			ReceiptsHash: []byte("receipts"),
		}
		assert.Equal(t, []byte("receipts"), er.GetReceiptsHash())
	})
}

func TestExecutionResult_GetDeveloperFees(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetDeveloperFees())
	})

	t.Run("with developer fees", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			DeveloperFees: big.NewInt(100),
		}
		assert.Equal(t, big.NewInt(100), er.GetDeveloperFees())
	})
}

func TestExecutionResult_GetAccumulatedFees(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Nil(t, er.GetAccumulatedFees())
	})

	t.Run("with accumulated fees", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			AccumulatedFees: big.NewInt(1000),
		}
		assert.Equal(t, big.NewInt(1000), er.GetAccumulatedFees())
	})
}

func TestExecutionResult_GetGasUsed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetGasUsed())
	})

	t.Run("with gas used", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				GasUsed: 50000,
			},
		}
		assert.Equal(t, uint64(50000), er.GetGasUsed())
	})
}

func TestExecutionResult_GetExecutedTxCount(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		er := (*ExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetExecutedTxCount())
	})

	t.Run("with executed tx count", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{
			ExecutedTxCount: 42,
		}
		assert.Equal(t, uint64(42), er.GetExecutedTxCount())
	})
}

func TestExecutionResult_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*ExecutionResult)(nil)
		assert.True(t, er.IsInterfaceNil())
	})

	t.Run("non-nil receiver", func(t *testing.T) {
		t.Parallel()

		er := &ExecutionResult{}
		assert.False(t, er.IsInterfaceNil())
	})
}

func TestExecutionResultInfo_GetExecutionResultHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		eri := (*ExecutionResultInfo)(nil)
		assert.Nil(t, eri.GetExecutionResultHandler())
	})

	t.Run("with execution result", func(t *testing.T) {
		t.Parallel()

		ber := &BaseExecutionResult{
			HeaderHash: []byte("test"),
		}
		eri := &ExecutionResultInfo{
			ExecutionResult: ber,
		}

		result := eri.GetExecutionResultHandler()
		assert.Equal(t, ber, result)
	})
}

func TestExecutionResultInfo_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		eri := (*ExecutionResultInfo)(nil)
		assert.True(t, eri.IsInterfaceNil())
	})

	t.Run("non-nil receiver", func(t *testing.T) {
		t.Parallel()

		eri := &ExecutionResultInfo{}
		assert.False(t, eri.IsInterfaceNil())
	})
}

func TestBaseExecutionResult_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		ber := (*BaseExecutionResult)(nil)
		assert.True(t, ber.IsInterfaceNil())
	})

	t.Run("non-nil receiver", func(t *testing.T) {
		t.Parallel()

		ber := &BaseExecutionResult{}
		assert.False(t, ber.IsInterfaceNil())
	})
}
