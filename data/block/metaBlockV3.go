//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. metaBlockV3.proto

package block

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// don't break the interface
var _ = data.HeaderHandler(&MetaBlockV3{})
var _ = data.MetaHeaderHandler(&MetaBlockV3{})

// GetExecutionResultsHandlers will return the execution result handlers
func (m *MetaBlockV3) GetExecutionResultsHandlers() []data.BaseExecutionResultHandler {
	if m == nil {
		return nil
	}

	executionResultsHandlers := make([]data.BaseExecutionResultHandler, len(m.ExecutionResults))
	for i, executionResult := range m.ExecutionResults {
		executionResultsHandlers[i] = executionResult
	}

	return executionResultsHandlers
}

// GetLastExecutionResultHandler will return the last execution result handler
func (m *MetaBlockV3) GetLastExecutionResultHandler() data.LastExecutionResultHandler {
	if m == nil {
		return nil
	}

	return m.LastExecutionResult
}

// GetValidatorStatsRootHash returns nil
func (m *MetaBlockV3) GetValidatorStatsRootHash() []byte {
	// TODO should we return the validators statistics root the last notarized execution result ?
	// OR should be have the validatorStatsRootHash as field on MetaBlockV3
	return nil
}

// GetDevFeesInEpoch returns nil
func (m *MetaBlockV3) GetDevFeesInEpoch() *big.Int {
	// TODO is correct to return the DevFeesInEpoch from the last execution result ?
	return nil
}

// GetAccumulatedFeesInEpoch returns nil
func (m *MetaBlockV3) GetAccumulatedFeesInEpoch() *big.Int {
	return nil
}

// GetEpochStartHandler will return the epoch start data
func (m *MetaBlockV3) GetEpochStartHandler() data.EpochStartHandler {
	if m == nil {
		return nil
	}

	return &m.EpochStart
}

// GetShardInfoHandlers gets the shardInfo as an array of ShardDataHandler
func (m *MetaBlockV3) GetShardInfoHandlers() []data.ShardDataHandler {
	if m == nil || m.ShardInfo == nil {
		return nil
	}

	shardInfoHandlers := make([]data.ShardDataHandler, len(m.ShardInfo))
	for i := range m.ShardInfo {
		shardInfoHandlers[i] = &m.ShardInfo[i]
	}

	return shardInfoHandlers
}

// SetShardInfoHandlers will set the provided shard info
func (m *MetaBlockV3) SetShardInfoHandlers(shardInfo []data.ShardDataHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if shardInfo == nil {
		m.ShardInfo = nil
		return nil
	}

	sInfo := make([]ShardData, len(shardInfo))
	for i := range shardInfo {
		shData, ok := shardInfo[i].(*ShardData)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if shData == nil {
			return data.ErrNilPointerDereference
		}
		sInfo[i] = *shData
	}

	m.ShardInfo = sInfo

	return nil
}

// SetValidatorStatsRootHash returns nil
func (m *MetaBlockV3) SetValidatorStatsRootHash(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetDevFeesInEpoch returns nil
func (m *MetaBlockV3) SetDevFeesInEpoch(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// SetAccumulatedFeesInEpoch returns nil
func (m *MetaBlockV3) SetAccumulatedFeesInEpoch(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// GetShardID returns the metachain shard id
func (m *MetaBlockV3) GetShardID() uint32 {
	return core.MetachainShardId
}

// GetRootHash always returns nil
func (m *MetaBlockV3) GetRootHash() []byte {
	return nil
}

// GetPubKeysBitmap always returns nil
func (m *MetaBlockV3) GetPubKeysBitmap() []byte {
	return nil
}

// GetSignature always returns nil
func (m *MetaBlockV3) GetSignature() []byte {
	return nil
}

// GetTimeStamp returns the timestamp
func (m *MetaBlockV3) GetTimeStamp() uint64 {
	if m == nil {
		return 0
	}
	return m.TimestampMs
}

// GetReceiptsHash always returns nil
func (m *MetaBlockV3) GetReceiptsHash() []byte {
	return nil
}

// GetAccumulatedFees always returns 0
func (m *MetaBlockV3) GetAccumulatedFees() *big.Int {
	return nil
}

// GetDeveloperFees always returns 0
func (m *MetaBlockV3) GetDeveloperFees() *big.Int {
	return nil
}

// GetMiniBlockHeadersWithDst as a map of hashes and sender IDs
func (m *MetaBlockV3) GetMiniBlockHeadersWithDst(destID uint32) map[string]uint32 {
	if m == nil {
		return nil
	}

	hashDst := make(map[string]uint32)
	for i := 0; i < len(m.ShardInfo); i++ {
		if m.ShardInfo[i].ShardID == destID {
			continue
		}

		for _, val := range m.ShardInfo[i].ShardMiniBlockHeaders {
			if val.ReceiverShardID == destID && val.SenderShardID != destID {
				hashDst[string(val.Hash)] = val.SenderShardID
			}
		}
	}

	for _, val := range m.MiniBlockHeaders {
		isDestinationShard := (val.ReceiverShardID == destID ||
			val.ReceiverShardID == core.AllShardId) &&
			val.SenderShardID != destID
		if isDestinationShard {
			hashDst[string(val.Hash)] = val.SenderShardID
		}
	}

	return hashDst
}

// GetOrderedCrossMiniblocksWithDst gets all cross miniblocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (m *MetaBlockV3) GetOrderedCrossMiniblocksWithDst(destID uint32) []*data.MiniBlockInfo {
	if m == nil {
		return nil
	}

	miniBlocks := make([]*data.MiniBlockInfo, 0)

	for i := 0; i < len(m.ShardInfo); i++ {
		if m.ShardInfo[i].ShardID == destID {
			continue
		}

		for _, mb := range m.ShardInfo[i].ShardMiniBlockHeaders {
			if mb.ReceiverShardID == destID && mb.SenderShardID != destID {
				miniBlocks = append(miniBlocks, &data.MiniBlockInfo{
					Hash:          mb.Hash,
					SenderShardID: mb.SenderShardID,
					Round:         m.ShardInfo[i].Round,
				})
			}
		}
	}

	for _, mb := range m.MiniBlockHeaders {
		isDestinationShard := (mb.ReceiverShardID == destID ||
			mb.ReceiverShardID == core.AllShardId) &&
			mb.SenderShardID != destID
		if isDestinationShard {
			miniBlocks = append(miniBlocks, &data.MiniBlockInfo{
				Hash:          mb.Hash,
				SenderShardID: mb.SenderShardID,
				Round:         m.Round,
			})
		}
	}

	sort.Slice(miniBlocks, func(i, j int) bool {
		return miniBlocks[i].Round < miniBlocks[j].Round
	})

	return miniBlocks
}

// GetMiniBlockHeadersHashes gets the miniblock hashes
func (m *MetaBlockV3) GetMiniBlockHeadersHashes() [][]byte {
	if m == nil {
		return nil
	}

	result := make([][]byte, 0, len(m.MiniBlockHeaders))
	for _, miniblock := range m.MiniBlockHeaders {
		result = append(result, miniblock.Hash)
	}

	return result
}

// GetMiniBlockHeaderHandlers returns the miniBlock headers as an array of miniBlock header handlers
func (m *MetaBlockV3) GetMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if m == nil {
		return nil
	}

	mbHeaders := m.GetMiniBlockHeaders()
	mbHeaderHandlers := make([]data.MiniBlockHeaderHandler, len(mbHeaders))

	for i := range mbHeaders {
		mbHeaderHandlers[i] = &mbHeaders[i]
	}

	return mbHeaderHandlers
}

// HasScheduledSupport returns false
func (m *MetaBlockV3) HasScheduledSupport() bool {
	return false
}

// GetAdditionalData gets the additional version-related data for the header
func (m *MetaBlockV3) GetAdditionalData() headerVersionData.HeaderAdditionalData {
	return nil
}

// HasScheduledMiniBlocks returns false
func (m *MetaBlockV3) HasScheduledMiniBlocks() bool {
	return false
}

// SetAccumulatedFees will do nothing
func (m *MetaBlockV3) SetAccumulatedFees(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// SetDeveloperFees will do nothing
func (m *MetaBlockV3) SetDeveloperFees(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// SetShardID will do nothing
func (m *MetaBlockV3) SetShardID(_ uint32) error {
	return nil
}

// SetNonce sets header nonce
func (m *MetaBlockV3) SetNonce(n uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Nonce = n

	return nil
}

// SetEpoch sets header epoch
func (m *MetaBlockV3) SetEpoch(e uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Epoch = e

	return nil
}

// SetRound sets header rounds
func (m *MetaBlockV3) SetRound(r uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Round = r

	return nil
}

// SetTimeStamp sets header timestamp
func (m *MetaBlockV3) SetTimeStamp(ts uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.TimestampMs = ts

	return nil
}

// SetRootHash will do nothing
func (m *MetaBlockV3) SetRootHash(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetPrevHash sets prev hash
func (m *MetaBlockV3) SetPrevHash(pvHash []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.PrevHash = pvHash

	return nil
}

// SetPrevRandSeed sets the previous randomness seed
func (m *MetaBlockV3) SetPrevRandSeed(pvRandSeed []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.PrevRandSeed = pvRandSeed

	return nil
}

// SetRandSeed sets the current random seed
func (m *MetaBlockV3) SetRandSeed(randSeed []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.RandSeed = randSeed

	return nil
}

// SetPubKeysBitmap always returns nil
func (m *MetaBlockV3) SetPubKeysBitmap(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetSignature always returns nil
func (m *MetaBlockV3) SetSignature(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetLeaderSignature will set the leader's signature
func (m *MetaBlockV3) SetLeaderSignature(sg []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.LeaderSignature = sg

	return nil
}

// SetChainID sets the chain ID on which this block is valid on
func (m *MetaBlockV3) SetChainID(chainID []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.ChainID = chainID

	return nil
}

// SetSoftwareVersion sets the software version of the block
func (m *MetaBlockV3) SetSoftwareVersion(version []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.SoftwareVersion = version

	return nil
}

// SetTxCount sets the transaction count of the current meta block
func (m *MetaBlockV3) SetTxCount(txCount uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.TxCount = txCount

	return nil
}

// SetMiniBlockHeaderHandlers sets the miniBlock headers from the given miniBlock header handlers
func (m *MetaBlockV3) SetMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if mbHeaderHandlers == nil {
		m.MiniBlockHeaders = nil
		return nil
	}

	mbHeaders := make([]MiniBlockHeader, len(mbHeaderHandlers))
	for i := range mbHeaderHandlers {
		mbHeader, ok := mbHeaderHandlers[i].(*MiniBlockHeader)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if mbHeader == nil {
			return data.ErrNilPointerDereference
		}
		mbHeaders[i] = *mbHeader
	}

	m.MiniBlockHeaders = mbHeaders

	return nil
}

// SetReceiptsHash always returns nil
func (m *MetaBlockV3) SetReceiptsHash(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetScheduledRootHash not supported on the first version of metablock
func (m *MetaBlockV3) SetScheduledRootHash(_ []byte) error {
	return data.ErrScheduledRootHashNotSupported
}

// ValidateHeaderVersion - always valid
func (m *MetaBlockV3) ValidateHeaderVersion() error {
	return nil
}

// SetAdditionalData sets the additional version-related data
func (m *MetaBlockV3) SetAdditionalData(_ headerVersionData.HeaderAdditionalData) error {
	return data.ErrFieldNotSupported
}

// IsStartOfEpochBlock verifies if the block is of type start of epoch
func (m *MetaBlockV3) IsStartOfEpochBlock() bool {
	if m == nil {
		return false
	}

	return len(m.EpochStart.LastFinalizedHeaders) > 0
}

// ShallowClone will return a clone of the object
func (m *MetaBlockV3) ShallowClone() data.HeaderHandler {
	if m == nil {
		return nil
	}

	metaBlockCopy := *m

	return &metaBlockCopy
}

// CheckFieldsForNil checks a predefined set of fields for nil values
func (m *MetaBlockV3) CheckFieldsForNil() error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if m.PrevHash == nil {
		return fmt.Errorf("%w in MetaBlockV3.PrevHash", data.ErrNilValue)
	}
	if m.PrevRandSeed == nil {
		return fmt.Errorf("%w in MetaBlockV3.PrevRandSeed", data.ErrNilValue)
	}
	if m.RandSeed == nil {
		return fmt.Errorf("%w in MetaBlockV3.RandSeed", data.ErrNilValue)
	}
	if m.LeaderSignature == nil {
		return fmt.Errorf("%w in MetaBlockV3.LeaderSignature", data.ErrNilValue)
	}
	if m.ChainID == nil {
		return fmt.Errorf("%w in MetaBlockV3.ChainID", data.ErrNilValue)
	}
	if m.SoftwareVersion == nil {
		return fmt.Errorf("%w in MetaBlockV3.SoftwareVersion", data.ErrNilValue)
	}

	return nil
}

// IsHeaderV3 checks if the header is of type MetaBlockV3
func (m *MetaBlockV3) IsHeaderV3() bool {
	return m != nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (m *MetaBlockV3) IsInterfaceNil() bool {
	return m == nil
}
