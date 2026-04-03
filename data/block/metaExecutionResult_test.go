package block

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/data"
)

func TestMetaExecutionResultInfo_GetExecutionResultHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResultInfo)(nil)
		require.Nil(t, mes.GetExecutionResultHandler())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResultInfo{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetExecutionResultHandler())
	})
	t.Run("with execution result", func(t *testing.T) {
		t.Parallel()

		expectedValue := &BaseMetaExecutionResult{}
		mes := &MetaExecutionResultInfo{
			ExecutionResult: expectedValue,
		}
		require.Equal(t, expectedValue, mes.GetExecutionResultHandler())
	})
}

func TestMetaExecutionResultInfo_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResultInfo)(nil)
		require.True(t, mes.IsInterfaceNil())
	})
	t.Run("not nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResultInfo{}
		require.False(t, mes.IsInterfaceNil())
	})
}

func TestBaseMetaExecutionResult_GetHeaderHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := (*BaseMetaExecutionResult)(nil)
		require.Nil(t, bm.GetHeaderHash())
	})
	t.Run("with header hash", func(t *testing.T) {
		t.Parallel()

		expectedValue := []byte("headerHash")
		bm := &BaseMetaExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderHash: expectedValue,
			},
		}
		require.Equal(t, expectedValue, bm.GetHeaderHash())
	})
}

func TestBaseMetaExecutionResult_GetHeaderNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := (*BaseMetaExecutionResult)(nil)
		require.Equal(t, uint64(0), bm.GetHeaderNonce())
	})
	t.Run("with header nonce", func(t *testing.T) {
		t.Parallel()

		expectedValue := uint64(42)
		bm := &BaseMetaExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderNonce: expectedValue,
			},
		}
		require.Equal(t, expectedValue, bm.GetHeaderNonce())
	})
}

func TestBaseMetaExecutionResult_GetHeaderRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := (*BaseMetaExecutionResult)(nil)
		require.Equal(t, uint64(0), bm.GetHeaderRound())
	})
	t.Run("with header round", func(t *testing.T) {
		t.Parallel()

		expectedValue := uint64(100)
		bm := &BaseMetaExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				HeaderRound: expectedValue,
			},
		}
		require.Equal(t, expectedValue, bm.GetHeaderRound())
	})
}

func TestBaseMetaExecutionResult_GetRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := (*BaseMetaExecutionResult)(nil)
		require.Nil(t, bm.GetRootHash())
	})
	t.Run("with root hash", func(t *testing.T) {
		t.Parallel()

		expectedValue := []byte("rootHash")
		bm := &BaseMetaExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				RootHash: expectedValue,
			},
		}
		require.Equal(t, expectedValue, bm.GetRootHash())
	})
}

func TestBaseMetaExecutionResult_GetGasUsed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*BaseMetaExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetGasUsed())
	})

	t.Run("with gas used", func(t *testing.T) {
		t.Parallel()

		er := &BaseMetaExecutionResult{
			BaseExecutionResult: &BaseExecutionResult{
				GasUsed: 50000,
			},
		}
		assert.Equal(t, uint64(50000), er.GetGasUsed())
	})
}

func TestBaseMetaExecutionResult_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := (*BaseMetaExecutionResult)(nil)
		require.True(t, bm.IsInterfaceNil())
	})
	t.Run("not nil receiver", func(t *testing.T) {
		t.Parallel()

		bm := &BaseMetaExecutionResult{}
		require.False(t, bm.IsInterfaceNil())
	})
}

func TestMetaExecutionResult_GetHeaderHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Nil(t, mes.GetHeaderHash())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetHeaderHash())
	})
	t.Run("with header hash", func(t *testing.T) {
		t.Parallel()

		expectedValue := []byte("headerHash")
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					HeaderHash: expectedValue,
				},
			},
		}
		require.Equal(t, expectedValue, mes.GetHeaderHash())
	})
}

func TestMetaExecutionResult_GetHeaderNonce(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Equal(t, uint64(0), mes.GetHeaderNonce())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Equal(t, uint64(0), mes.GetHeaderNonce())
	})
	t.Run("with header nonce", func(t *testing.T) {
		t.Parallel()

		expectedValue := uint64(42)
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					HeaderNonce: expectedValue,
				},
			},
		}
		require.Equal(t, expectedValue, mes.GetHeaderNonce())
	})
}

func TestMetaExecutionResult_GetHeaderRound(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Equal(t, uint64(0), mes.GetHeaderRound())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Equal(t, uint64(0), mes.GetHeaderRound())
	})
	t.Run("with header round", func(t *testing.T) {
		t.Parallel()

		expectedValue := uint64(100)
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					HeaderRound: expectedValue,
				},
			},
		}
		require.Equal(t, expectedValue, mes.GetHeaderRound())
	})
}

func TestMetaExecutionResult_GetHeaderEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*MetaExecutionResult)(nil)
		assert.Equal(t, uint32(0), er.GetHeaderEpoch())
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					HeaderEpoch: 10,
				},
			},
		}
		assert.Equal(t, uint32(10), er.GetHeaderEpoch())
	})
}

func TestMetaExecutionResult_GetRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Nil(t, mes.GetRootHash())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetRootHash())
	})
	t.Run("with root hash", func(t *testing.T) {
		t.Parallel()

		expectedValue := []byte("rootHash")
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					RootHash: expectedValue,
				},
			},
		}
		require.Equal(t, expectedValue, mes.GetRootHash())
	})
}

func TestMetaExecutionResult_GetValidatorStatsRootHash(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Nil(t, mes.GetValidatorStatsRootHash())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetValidatorStatsRootHash())
	})
	t.Run("with validator stats root hash", func(t *testing.T) {
		t.Parallel()

		expectedValue := []byte("validatorStatsRootHash")
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				ValidatorStatsRootHash: expectedValue,
			},
		}
		require.Equal(t, expectedValue, mes.GetValidatorStatsRootHash())
	})
}

func TestMetaExecutionResult_GetDevFeesInEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Nil(t, mes.GetDevFeesInEpoch())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetDevFeesInEpoch())
	})
	t.Run("with dev fees", func(t *testing.T) {
		t.Parallel()

		expectedValue := big.NewInt(100)
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				DevFeesInEpoch: expectedValue,
			},
		}
		require.Equal(t, expectedValue, mes.GetDevFeesInEpoch())
	})
}

func TestMetaExecutionResult_GetMiniBlockHeadersHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		er := (*MetaExecutionResult)(nil)
		assert.Nil(t, er.GetMiniBlockHeadersHandlers())
	})

	t.Run("no miniblocks", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{
			MiniBlockHeaders: nil,
		}
		assert.Empty(t, er.GetMiniBlockHeadersHandlers())
	})

	t.Run("with miniblocks", func(t *testing.T) {
		t.Parallel()

		mb1 := MiniBlockHeader{Hash: []byte("hash1")}
		mb2 := MiniBlockHeader{Hash: []byte("hash2")}
		er := &MetaExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{mb1, mb2},
		}

		handlers := er.GetMiniBlockHeadersHandlers()
		assert.Len(t, handlers, 2)
		assert.Equal(t, mb1.Hash, handlers[0].GetHash())
		assert.Equal(t, mb2.Hash, handlers[1].GetHash())
	})
}

func TestMetaExecutionResult_SetMiniBlockHeadersHandlers(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*MetaExecutionResult)(nil)
		err := er.SetMiniBlockHeadersHandlers(nil)
		assert.Equal(t, data.ErrNilPointerReceiver, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{{Hash: []byte("old")}},
		}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{})
		assert.NoError(t, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with nil slice", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{
			MiniBlockHeaders: []MiniBlockHeader{{Hash: []byte("old")}},
		}
		err := er.SetMiniBlockHeadersHandlers(nil)
		assert.NoError(t, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with invalid mini block header handlers", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{nil})
		assert.Equal(t, data.ErrInvalidTypeAssertion, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with nil mini block header handler", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers([]data.MiniBlockHeaderHandler{(*MiniBlockHeader)(nil)})
		assert.Equal(t, data.ErrNilPointerDereference, err)
		assert.Nil(t, er.MiniBlockHeaders)
	})
	t.Run("with valid miniblocks", func(t *testing.T) {
		t.Parallel()

		mb1 := &MiniBlockHeader{Hash: []byte("hash1")}
		mb2 := &MiniBlockHeader{Hash: []byte("hash2")}
		handlers := []data.MiniBlockHeaderHandler{mb1, mb2}

		er := &MetaExecutionResult{}
		err := er.SetMiniBlockHeadersHandlers(handlers)
		assert.NoError(t, err)
		assert.Len(t, er.MiniBlockHeaders, 2)
		assert.Equal(t, mb1.Hash, er.MiniBlockHeaders[0].Hash)
		assert.Equal(t, mb2.Hash, er.MiniBlockHeaders[1].Hash)
	})
}

func TestMetaExecutionResult_GetAccumulatedFeesInEpoch(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.Nil(t, mes.GetAccumulatedFeesInEpoch())
	})
	t.Run("nil execution result", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{
			ExecutionResult: nil,
		}
		require.Nil(t, mes.GetAccumulatedFeesInEpoch())
	})
	t.Run("with accumulated fees", func(t *testing.T) {
		t.Parallel()

		expectedValue := big.NewInt(100)
		mes := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				AccumulatedFeesInEpoch: expectedValue,
			},
		}
		require.Equal(t, expectedValue, mes.GetAccumulatedFeesInEpoch())
	})
}

func TestMetaExecutionResult_GetGasUsed(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		er := (*MetaExecutionResult)(nil)
		assert.Equal(t, uint64(0), er.GetGasUsed())
	})

	t.Run("with gas used", func(t *testing.T) {
		t.Parallel()

		er := &MetaExecutionResult{
			ExecutionResult: &BaseMetaExecutionResult{
				BaseExecutionResult: &BaseExecutionResult{
					GasUsed: 50000,
				},
			},
		}
		assert.Equal(t, uint64(50000), er.GetGasUsed())
	})
}

func TestMetaExecutionResult_IsInterfaceNil(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := (*MetaExecutionResult)(nil)
		require.True(t, mes.IsInterfaceNil())
	})
	t.Run("not nil receiver", func(t *testing.T) {
		t.Parallel()

		mes := &MetaExecutionResult{}
		require.False(t, mes.IsInterfaceNil())
	})
}
