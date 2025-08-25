//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. blockV3.proto

package block

import (
	"fmt"
	"math/big"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/headerVersionData"
)

// GetRootHash always returns nil
func (hv3 *HeaderV3) GetRootHash() []byte {
	return nil
}

// GetPubKeysBitmap always returns nil
func (hv3 *HeaderV3) GetPubKeysBitmap() []byte {
	return nil
}

// GetSignature always returns nil
func (hv3 *HeaderV3) GetSignature() []byte {
	return nil
}

// GetTimeStamp returns the timestamp
func (hv3 *HeaderV3) GetTimeStamp() uint64 {
	if hv3 == nil {
		return 0
	}

	return hv3.TimestampMs
}

// GetMiniBlockHeadersWithDst as a map of hashes and sender IDs
func (hv3 *HeaderV3) GetMiniBlockHeadersWithDst(destId uint32) map[string]uint32 {
	if hv3 == nil {
		return nil
	}

	hashDst := make(map[string]uint32)
	for _, val := range hv3.MiniBlockHeaders {
		if val.ReceiverShardID == destId && val.SenderShardID != destId {
			hashDst[string(val.Hash)] = val.SenderShardID
		}
	}
	return hashDst
}

// GetOrderedCrossMiniblocksWithDst gets all cross miniblocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (hv3 *HeaderV3) GetOrderedCrossMiniblocksWithDst(destId uint32) []*data.MiniBlockInfo {
	if hv3 == nil {
		return nil
	}

	miniBlocks := make([]*data.MiniBlockInfo, 0)

	for _, mb := range hv3.MiniBlockHeaders {
		if mb.ReceiverShardID == destId && mb.SenderShardID != destId {
			miniBlocks = append(miniBlocks, &data.MiniBlockInfo{
				Hash:          mb.Hash,
				SenderShardID: mb.SenderShardID,
				Round:         hv3.Round,
			})
		}
	}

	return miniBlocks
}

// GetMiniBlockHeadersHashes gets the miniblock hashes
func (hv3 *HeaderV3) GetMiniBlockHeadersHashes() [][]byte {
	if hv3 == nil {
		return nil
	}

	result := make([][]byte, 0, len(hv3.MiniBlockHeaders))
	for _, miniblock := range hv3.MiniBlockHeaders {
		result = append(result, miniblock.Hash)
	}
	return result
}

// GetMiniBlockHeaderHandlers returns the miniBlock headers as an array of miniBlock header handlers
func (hv3 *HeaderV3) GetMiniBlockHeaderHandlers() []data.MiniBlockHeaderHandler {
	if hv3 == nil {
		return nil
	}

	mbHeaders := hv3.GetMiniBlockHeaders()
	mbHeaderHandlers := make([]data.MiniBlockHeaderHandler, len(mbHeaders))

	for i := range mbHeaders {
		mbHeaderHandlers[i] = &mbHeaders[i]
	}

	return mbHeaderHandlers
}

// HasScheduledSupport always returns false
func (hv3 *HeaderV3) HasScheduledSupport() bool {
	return false
}

// GetAdditionalData always returns nil
func (hv3 *HeaderV3) GetAdditionalData() headerVersionData.HeaderAdditionalData {
	return nil
}

// HasScheduledMiniBlocks always returns false
func (hv3 *HeaderV3) HasScheduledMiniBlocks() bool {
	return false
}

// SetAccumulatedFees always returns nil
func (hv3 *HeaderV3) SetAccumulatedFees(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// SetDeveloperFees always returns nil
func (hv3 *HeaderV3) SetDeveloperFees(_ *big.Int) error {
	return data.ErrFieldNotSupported
}

// SetShardID sets header shard ID
func (hv3 *HeaderV3) SetShardID(shId uint32) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.ShardID = shId
	return nil
}

// SetNonce sets header nonce
func (hv3 *HeaderV3) SetNonce(n uint64) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.Nonce = n
	return nil
}

// SetEpoch sets header epoch
func (hv3 *HeaderV3) SetEpoch(e uint32) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.Epoch = e
	return nil
}

// SetRound sets header round
func (hv3 *HeaderV3) SetRound(r uint64) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.Round = r
	return nil
}

// SetTimeStamp sets header timestamp
func (hv3 *HeaderV3) SetTimeStamp(ts uint64) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.TimestampMs = ts
	return nil
}

// SetRootHash always returns nil
func (hv3 *HeaderV3) SetRootHash(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetPrevHash sets prev hash
func (hv3 *HeaderV3) SetPrevHash(pvHash []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.PrevHash = pvHash
	return nil
}

// SetPrevRandSeed sets previous random seed
func (hv3 *HeaderV3) SetPrevRandSeed(pvRandSeed []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.PrevRandSeed = pvRandSeed
	return nil
}

// SetRandSeed sets previous random seed
func (hv3 *HeaderV3) SetRandSeed(randSeed []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.RandSeed = randSeed
	return nil
}

// SetPubKeysBitmap always returns nil
func (hv3 *HeaderV3) SetPubKeysBitmap(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetSignature always returns nil
func (hv3 *HeaderV3) SetSignature(_ []byte) error {
	return data.ErrFieldNotSupported
}

// SetLeaderSignature sets the leader's signature
func (hv3 *HeaderV3) SetLeaderSignature(sg []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.LeaderSignature = sg
	return nil
}

// SetChainID sets the chain ID on which this block is valid on
func (hv3 *HeaderV3) SetChainID(chainID []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.ChainID = chainID
	return nil
}

// SetSoftwareVersion sets the software version of the header
func (hv3 *HeaderV3) SetSoftwareVersion(version []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.SoftwareVersion = version
	return nil
}

// SetTxCount sets the transaction count of the proposed block associated with this header
func (hv3 *HeaderV3) SetTxCount(txCount uint32) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.TxCount = txCount
	return nil
}

// SetMiniBlockHeaderHandlers sets the miniBlock headers from the given miniBlock header handlers
func (hv3 *HeaderV3) SetMiniBlockHeaderHandlers(mbHeaderHandlers []data.MiniBlockHeaderHandler) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}
	if len(mbHeaderHandlers) == 0 {
		hv3.MiniBlockHeaders = nil
		return nil
	}

	miniBlockHeaders := make([]MiniBlockHeader, len(mbHeaderHandlers))
	for i, mbHeaderHandler := range mbHeaderHandlers {
		mbHeader, ok := mbHeaderHandler.(*MiniBlockHeader)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if mbHeader == nil {
			return data.ErrNilPointerDereference
		}
		miniBlockHeaders[i] = *mbHeader
	}

	hv3.MiniBlockHeaders = miniBlockHeaders
	return nil
}

// SetReceiptsHash sets the receipts hash
func (hv3 *HeaderV3) SetReceiptsHash(hash []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.ReceiptsHash = hash
	return nil
}

// SetScheduledRootHash always returns nil
func (hv3 *HeaderV3) SetScheduledRootHash(_ []byte) error {
	return data.ErrFieldNotSupported
}

// ValidateHeaderVersion always returns nil
func (hv3 *HeaderV3) ValidateHeaderVersion() error {
	return nil
}

// SetAdditionalData always returns nil
func (hv3 *HeaderV3) SetAdditionalData(_ headerVersionData.HeaderAdditionalData) error {
	return data.ErrFieldNotSupported
}

// IsStartOfEpochBlock verifies if the block is of type start of epoch
func (hv3 *HeaderV3) IsStartOfEpochBlock() bool {
	if hv3 == nil {
		return false
	}

	return len(hv3.EpochStartMetaHash) > 0
}

// ShallowClone returns a clone of the object
func (hv3 *HeaderV3) ShallowClone() data.HeaderHandler {
	if hv3 == nil {
		return nil
	}

	headerCopy := *hv3
	return &headerCopy
}

// CheckFieldsForNil checks a predefined set of fields for nil values
func (hv3 *HeaderV3) CheckFieldsForNil() error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}
	if hv3.PrevHash == nil {
		return fmt.Errorf("%w in Header.PrevHash", data.ErrNilValue)
	}
	if hv3.PrevRandSeed == nil {
		return fmt.Errorf("%w in Header.PrevRandSeed", data.ErrNilValue)
	}
	if hv3.RandSeed == nil {
		return fmt.Errorf("%w in Header.RandSeed", data.ErrNilValue)
	}
	if hv3.LeaderSignature == nil {
		return fmt.Errorf("%w in Header.LeaderSignature", data.ErrNilValue)
	}
	if hv3.SoftwareVersion == nil {
		return fmt.Errorf("%w in Header.SoftwareVersion", data.ErrNilValue)
	}
	if hv3.LastExecutionResult == nil {
		return fmt.Errorf("%w in Header.LastExecutionResult", data.ErrNilValue)
	}

	return nil
}

// GetAccumulatedFees always returns 0
func (hv3 *HeaderV3) GetAccumulatedFees() *big.Int {
	return nil
}

// GetDeveloperFees always returns 0
func (hv3 *HeaderV3) GetDeveloperFees() *big.Int {
	return nil
}

// SetEpochStartMetaHash sets the epoch start metaBlock hash
func (hv3 *HeaderV3) SetEpochStartMetaHash(hash []byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}
	hv3.EpochStartMetaHash = hash
	return nil
}

// GetBlockBodyTypeInt32 returns the block body type as int32
func (hv3 *HeaderV3) GetBlockBodyTypeInt32() int32 {
	if hv3 == nil {
		return -1
	}

	return int32(hv3.GetBlockBodyType())
}

// SetMetaBlockHashes sets the metaBlock hashes
func (hv3 *HeaderV3) SetMetaBlockHashes(hashes [][]byte) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}
	hv3.MetaBlockHashes = hashes
	return nil
}

// MapMiniBlockHashesToShards is a map of mini block hashes and sender IDs
func (hv3 *HeaderV3) MapMiniBlockHashesToShards() map[string]uint32 {
	if hv3 == nil {
		return nil
	}

	hashDst := make(map[string]uint32)
	for _, val := range hv3.MiniBlockHeaders {
		hashDst[string(val.Hash)] = val.SenderShardID
	}
	return hashDst
}

// SetBlockBodyTypeInt32 sets the blockBodyType in the header
func (hv3 *HeaderV3) SetBlockBodyTypeInt32(blockBodyType int32) error {
	if hv3 == nil {
		return data.ErrNilPointerReceiver
	}

	hv3.BlockBodyType = Type(blockBodyType)

	return nil
}

// GetLastExecutionResultHandler returns the last execution result
func (hv3 *HeaderV3) GetLastExecutionResultHandler() data.LastExecutionResultHandler {
	if hv3 == nil {
		return nil
	}

	return hv3.LastExecutionResult
}

// GetExecutionResultsHandlers returns the execution results
func (hv3 *HeaderV3) GetExecutionResultsHandlers() []data.BaseExecutionResultHandler {
	if hv3 == nil {
		return nil
	}

	executionResultsHandlers := make([]data.BaseExecutionResultHandler, 0, len(hv3.GetExecutionResults()))
	for _, execResult := range hv3.GetExecutionResults() {
		executionResultsHandlers = append(executionResultsHandlers, execResult)
	}

	return executionResultsHandlers
}

// IsHeaderV3 checks if the header is of type HeaderV3
func (hv3 *HeaderV3) IsHeaderV3() bool {
	return hv3 != nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (hv3 *HeaderV3) IsInterfaceNil() bool {
	return hv3 == nil
}
