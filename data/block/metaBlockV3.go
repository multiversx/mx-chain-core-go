//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. metaBlockV3.proto

package block

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
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

// SetExecutionResultsHandlers will set the provided execution results
func (m *MetaBlockV3) SetExecutionResultsHandlers(execResults []data.BaseExecutionResultHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if len(execResults) == 0 {
		m.ExecutionResults = nil
		return nil
	}

	executionResults := make([]*MetaExecutionResult, len(execResults))
	for i, execResult := range execResults {
		execResultV3, ok := execResult.(*MetaExecutionResult)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if execResultV3 == nil {
			return data.ErrNilPointerDereference
		}
		executionResults[i] = execResultV3
	}

	m.ExecutionResults = executionResults
	return nil
}

// GetLastExecutionResultHandler will return the last execution result handler
func (m *MetaBlockV3) GetLastExecutionResultHandler() data.LastExecutionResultHandler {
	if m == nil {
		return nil
	}

	return m.LastExecutionResult
}

// GetAccumulatedFeesInEpoch returns nil
func (m *MetaBlockV3) GetAccumulatedFeesInEpoch() *big.Int {
	return nil
}

// SetLastExecutionResultHandler will set the provided last execution result
func (m *MetaBlockV3) SetLastExecutionResultHandler(lastExecResult data.LastExecutionResultHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if lastExecResult == nil {
		return data.ErrNilPointerDereference
	}

	lastExecResultV3, ok := lastExecResult.(*MetaExecutionResultInfo)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}

	m.LastExecutionResult = lastExecResultV3
	return nil
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

// GetShardInfoProposalHandlers gets the shardInfoProposal as an array of ShardDataProposalHandler
func (m *MetaBlockV3) GetShardInfoProposalHandlers() []data.ShardDataProposalHandler {
	if m == nil || m.ShardInfoProposal == nil {
		return nil
	}

	shardInfoProposalHandlers := make([]data.ShardDataProposalHandler, len(m.ShardInfoProposal))
	for i := range m.ShardInfoProposal {
		shardInfoProposalHandlers[i] = &m.ShardInfoProposal[i]
	}

	return shardInfoProposalHandlers
}

// SetShardInfoProposalHandlers will set the provided shard info proposal
func (m *MetaBlockV3) SetShardInfoProposalHandlers(shardInfo []data.ShardDataProposalHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if shardInfo == nil {
		m.ShardInfoProposal = nil
		return nil
	}

	sInfo := make([]ShardDataProposal, len(shardInfo))
	for i := range shardInfo {
		shData, ok := shardInfo[i].(*ShardDataProposal)
		if !ok {
			return data.ErrInvalidTypeAssertion
		}
		if shData == nil {
			return data.ErrNilPointerDereference
		}
		sInfo[i] = *shData
	}

	m.ShardInfoProposal = sInfo

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

// GetMiniBlockHeadersWithDst returns a map of hashes and sender IDs
func (m *MetaBlockV3) GetMiniBlockHeadersWithDst(destID uint32) map[string]uint32 {
	if m == nil {
		return nil
	}

	hashDst := make(map[string]uint32)
	for i := 0; i < len(m.ShardInfo); i++ {
		if m.ShardInfo[i].ShardID == destID {
			continue
		}

		addShardMBHeadersMBToDestMap(m.ShardInfo[i].ShardMiniBlockHeaders, hashDst, destID)
	}

	for _, execResults := range m.ExecutionResults {
		addMetaMBHeadersMBToDestMap(execResults.MiniBlockHeaders, hashDst, destID)
	}

	return hashDst
}

func addMetaMBHeadersMBToDestMap(miniBlockHeaders []MiniBlockHeader, hashDst map[string]uint32, destID uint32) {
	for _, mbHeader := range miniBlockHeaders {
		isDestinationShard := (mbHeader.ReceiverShardID == destID ||
			mbHeader.ReceiverShardID == core.AllShardId) &&
			mbHeader.SenderShardID != destID
		if isDestinationShard {
			hashDst[string(mbHeader.Hash)] = mbHeader.SenderShardID
		}
	}
}

// GetProposedMiniBlockHeadersWithDst returns a map of hashes and sender IDs for proposed mini blocks
func (m *MetaBlockV3) GetProposedMiniBlockHeadersWithDst(destID uint32) map[string]uint32 {
	if m == nil {
		return nil
	}

	hashDst := make(map[string]uint32)
	addMetaMBHeadersMBToDestMap(m.MiniBlockHeaders, hashDst, destID)

	return hashDst
}

// GetOrderedCrossMiniblocksWithDst gets all cross miniblocks with the given destination shard ID, ordered in a
// chronological way, taking into consideration the round in which they were created/executed in the sender shard
func (m *MetaBlockV3) GetOrderedCrossMiniblocksWithDst(destID uint32) []*data.MiniBlockInfo {
	if m == nil {
		return nil
	}

	miniBlocks := getCrossMiniBlocksFromShardInfo(m.ShardInfo, destID)
	miniBlockHeaders := m.getMiniBlocksWithDstFromMetaExecutionResults(destID)
	miniBlocks = append(miniBlocks, miniBlockHeaders...)

	sort.Slice(miniBlocks, func(i, j int) bool {
		return miniBlocks[i].Round < miniBlocks[j].Round
	})

	return miniBlocks
}

func (m *MetaBlockV3) getMiniBlocksWithDstFromMetaExecutionResults(destID uint32) []*data.MiniBlockInfo {
	miniBlocks := make([]*data.MiniBlockInfo, 0)

	for _, execResults := range m.ExecutionResults {
		mbs := getCrossMiniBlocksFromMiniBlockHeaders(execResults.MiniBlockHeaders, destID, execResults.GetHeaderRound())
		miniBlocks = append(miniBlocks, mbs...)
	}

	return miniBlocks
}
func getCrossMiniBlocksFromShardInfo(shardInfo []ShardData, destId uint32) []*data.MiniBlockInfo {
	miniBlocks := make([]*data.MiniBlockInfo, 0)
	for i := 0; i < len(shardInfo); i++ {
		if shardInfo[i].ShardID == destId {
			continue
		}
		mbs := getCrossMiniBlocksFromMiniBlockHeaders(shardInfo[i].ShardMiniBlockHeaders, destId, shardInfo[i].Round)
		mbs = removeMiniBlocksFromShard(mbs, core.MetachainShardId)
		miniBlocks = append(miniBlocks, mbs...)
	}

	return miniBlocks
}

func removeMiniBlocksFromShard(miniBlocks []*data.MiniBlockInfo, shardID uint32) []*data.MiniBlockInfo {
	filteredMiniBlocks := make([]*data.MiniBlockInfo, 0)
	for _, mb := range miniBlocks {
		if mb.SenderShardID != shardID {
			filteredMiniBlocks = append(filteredMiniBlocks, mb)
		}
	}
	return filteredMiniBlocks
}

func getCrossMiniBlocksFromMiniBlockHeaders(miniBlockHeaders []MiniBlockHeader, destId uint32, round uint64) []*data.MiniBlockInfo {
	miniBlocks := make([]*data.MiniBlockInfo, 0)

	for _, mb := range miniBlockHeaders {
		isDestinationShard := (mb.ReceiverShardID == destId ||
			mb.ReceiverShardID == core.AllShardId) &&
			mb.SenderShardID != destId
		if isDestinationShard {
			miniBlocks = append(miniBlocks, &data.MiniBlockInfo{
				Hash:          mb.Hash,
				SenderShardID: mb.SenderShardID,
				Round:         round,
			})
		}
	}

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
	if m.LastExecutionResult == nil {
		return fmt.Errorf("%w in MetaBlockV3.LastExecutionResult", data.ErrNilValue)
	}

	return nil
}

// CheckFieldsIntegrity checks the integrity of the fields
// It checks a predefined set of fields for nil values or invalid values.
// It also checks the integrity of LastExecutionResult and ExecutionResults against the header
func (m *MetaBlockV3) CheckFieldsIntegrity() error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if len(m.Reserved) != 0 {
		return data.ErrNotNilValue
	}
	if len(m.ShardInfo) != 0 && len(m.ShardInfoProposal) == 0 {
		return fmt.Errorf("MetaBlockV3.ShardInfoProposal cannot be nil when MetaBlockV3.ShardInfo is not nil")
	}

	isGenesisRound := m.GetNonce() == 0
	if isGenesisRound {
		return nil
	}

	err := m.checkLastExecutionResultIntegrity()
	if err != nil {
		return err
	}
	if m.Round < m.LastExecutionResult.NotarizedInRound {
		return fmt.Errorf("MetaBlockV3.Round (%d) must be greater than or equal to LastExecutionResult.NotarizedInRound (%d)", m.Round, m.LastExecutionResult.NotarizedInRound)
	}

	if len(m.ExecutionResults) > 0 {
		err := m.checkExecutionResultsIntegrity()
		if err != nil {
			return err
		}
	}

	return nil
}

// checkLastExecutionResultIntegrity checks the integrity of the last execution result against the header it is associated with
func (m *MetaBlockV3) checkLastExecutionResultIntegrity() error {
	if m.LastExecutionResult == nil {
		return fmt.Errorf("%w in Header.LastExecutionResult", data.ErrNilValue)
	}

	return m.checkBaseMetaExecutionResultIntegrity(m.LastExecutionResult.ExecutionResult)
}

// checkExecutionResultsIntegrity checks the integrity of the execution results against the header they are associated with
func (m *MetaBlockV3) checkExecutionResultsIntegrity() error {

	for i, execResult := range m.ExecutionResults {
		if execResult == nil || check.IfNil(execResult) {
			return fmt.Errorf("%w in MetaBlockV3.ExecutionResults at index %d", data.ErrNilValue, i)
		}

		if len(execResult.ReceiptsHash) == 0 {
			return fmt.Errorf("%w in MetaExecutionResult.ReceiptsHash at index %d", data.ErrNilValue, i)
		}
		if execResult.AccumulatedFees == nil {
			return fmt.Errorf("%w in MetaExecutionResult.AccumulatedFees at index %d", data.ErrNilValue, i)
		}
		if execResult.AccumulatedFees.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("%w: MetaExecutionResult.AccumulatedFees cannot be negative at index %d", data.ErrInvalidValue, i)
		}
		if execResult.DeveloperFees == nil {
			return fmt.Errorf("%w in MetaExecutionResult.DeveloperFees at index %d", data.ErrNilValue, i)
		}
		if execResult.DeveloperFees.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("%w: MetaExecutionResult.DeveloperFees cannot be negative at index %d", data.ErrInvalidValue, i)
		}

		err := m.checkBaseMetaExecutionResultIntegrity(execResult.ExecutionResult)
		if err != nil {
			return fmt.Errorf("execution result integrity check failed at index %d: %w", i, err)
		}
	}

	return nil
}

func (m *MetaBlockV3) checkBaseMetaExecutionResultIntegrity(ownBaseMetaExecutionResult *BaseMetaExecutionResult) error {
	if ownBaseMetaExecutionResult == nil || check.IfNil(ownBaseMetaExecutionResult) {
		return data.ErrNilValue
	}

	if len(ownBaseMetaExecutionResult.GetValidatorStatsRootHash()) == 0 {
		return fmt.Errorf("%w in BaseMetaExecutionResult.ValidatorStatsRootHash", data.ErrNilValue)
	}
	if ownBaseMetaExecutionResult.AccumulatedFeesInEpoch == nil {
		return fmt.Errorf("%w in BaseMetaExecutionResult.AccumulatedFeesInEpoch", data.ErrNilValue)
	}
	if ownBaseMetaExecutionResult.AccumulatedFeesInEpoch.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("%w: BaseMetaExecutionResult.AccumulatedFeesInEpoch cannot be negative", data.ErrInvalidValue)
	}
	if ownBaseMetaExecutionResult.DevFeesInEpoch == nil {
		return fmt.Errorf("%w in BaseMetaExecutionResult.DevFeesInEpoch", data.ErrNilValue)
	}
	if ownBaseMetaExecutionResult.DevFeesInEpoch.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("%w: BaseMetaExecutionResult.DevFeesInEpoch cannot be negative", data.ErrInvalidValue)
	}

	err := m.checkBaseExecutionResultIntegrity(ownBaseMetaExecutionResult.BaseExecutionResult)
	if err != nil {
		return err
	}

	return nil
}

// checkBaseExecutionResultIntegrity checks the integrity of a base execution result against the header it is associated with
func (m *MetaBlockV3) checkBaseExecutionResultIntegrity(ownBaseExecutionResult data.BaseExecutionResultHandler) error {
	if ownBaseExecutionResult == nil || check.IfNil(ownBaseExecutionResult) {
		return data.ErrNilValue
	}

	if len(ownBaseExecutionResult.GetHeaderHash()) == 0 {
		return fmt.Errorf("%w in BaseExecutionResult.HeaderHash", data.ErrNilValue)
	}

	if ownBaseExecutionResult.GetHeaderNonce() >= m.Nonce {
		return fmt.Errorf("BaseExecutionResult.HeaderNonce (%d) must be less than Header.Nonce (%d)", ownBaseExecutionResult.GetHeaderNonce(), m.Nonce)
	}

	if ownBaseExecutionResult.GetHeaderRound() >= m.Round {
		return fmt.Errorf("BaseExecutionResult.HeaderRound (%d) must be less than Header.Round (%d)", ownBaseExecutionResult.GetHeaderRound(), m.Round)
	}

	if ownBaseExecutionResult.GetHeaderEpoch() > m.Epoch {
		return fmt.Errorf("BaseExecutionResult.HeaderEpoch (%d) must be less than or equal to Header.Epoch (%d)", ownBaseExecutionResult.GetHeaderEpoch(), m.Epoch)
	}

	if len(ownBaseExecutionResult.GetRootHash()) == 0 {
		return fmt.Errorf("%w in BaseExecutionResult.RootHash", data.ErrNilValue)
	}

	return nil
}

// SetEpochChangeProposed will set the provided value EpochStartProposed field
func (m *MetaBlockV3) SetEpochChangeProposed(value bool) {
	m.EpochChangeProposed = value
}

// SetEpochStartHandler sets the epoch start handler
func (m *MetaBlockV3) SetEpochStartHandler(epochStartHandler data.EpochStartHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	if epochStartHandler == nil {
		return nil
	}

	es, ok := epochStartHandler.(*EpochStart)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	if es == nil {
		return data.ErrNilPointerDereference
	}
	m.EpochStart = *es

	return nil
}

// IsEpochChangeProposed returns true if the current meta block v3 proposes an epoch change event
func (m *MetaBlockV3) IsEpochChangeProposed() bool {
	return m.EpochChangeProposed
}

// IsHeaderV3 checks if the header is of type MetaBlockV3
func (m *MetaBlockV3) IsHeaderV3() bool {
	return m != nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (m *MetaBlockV3) IsInterfaceNil() bool {
	return m == nil
}
