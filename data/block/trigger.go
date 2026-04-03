//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. trigger.proto
package block

import (
	"fmt"

	"github.com/multiversx/mx-chain-core-go/data"
)

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (str *ShardTriggerRegistry) GetEpochStartHeaderHandler() data.HeaderHandler {
	if str == nil {
		return nil
	}
	return str.GetEpochStartShardHeader()
}

// SetIsEpochStart sets the isEpochStart flag
func (str *ShardTriggerRegistry) SetIsEpochStart(isEpochStart bool) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}

	str.IsEpochStart = isEpochStart
	return nil
}

// SetNewEpochHeaderReceived sets the newEpochHeaderReceived flag
func (str *ShardTriggerRegistry) SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.NewEpochHeaderReceived = newEpochHeaderReceived
	return nil
}

// SetEpoch sets the epoch
func (str *ShardTriggerRegistry) SetEpoch(epoch uint32) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.Epoch = epoch
	return nil
}

// SetMetaEpoch Sets the metaChain epoch
func (str *ShardTriggerRegistry) SetMetaEpoch(metaEpoch uint32) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.MetaEpoch = metaEpoch
	return nil
}

// SetCurrentRoundIndex sets the current round index
func (str *ShardTriggerRegistry) SetCurrentRoundIndex(roundIndex int64) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.CurrentRoundIndex = roundIndex
	return nil
}

// SetEpochStartRound sets the epoch start round
func (str *ShardTriggerRegistry) SetEpochStartRound(startRound uint64) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.EpochStartRound = startRound
	return nil
}

// SetEpochFinalityAttestingRound sets the epoch finality attesting round
func (str *ShardTriggerRegistry) SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.EpochFinalityAttestingRound = finalityAttestingRound
	return nil
}

// SetEpochMetaBlockHash sets the epoch metaChain block hash
func (str *ShardTriggerRegistry) SetEpochMetaBlockHash(epochMetaBlockHash []byte) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}
	str.EpochMetaBlockHash = epochMetaBlockHash
	return nil
}

// SetEpochStartHeaderHandler sets the epoch start header
func (str *ShardTriggerRegistry) SetEpochStartHeaderHandler(epochStartHeaderHandler data.HeaderHandler) error {
	if str == nil {
		return data.ErrNilPointerReceiver
	}

	var ok bool
	str.EpochStartShardHeader, ok = epochStartHeaderHandler.(*Header)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	return nil
}

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (strV2 *ShardTriggerRegistryV2) GetEpochStartHeaderHandler() data.HeaderHandler {
	if strV2 == nil {
		return nil
	}
	return strV2.GetEpochStartShardHeader()
}

// SetIsEpochStart sets the isEpochStart flag
func (strV2 *ShardTriggerRegistryV2) SetIsEpochStart(isEpochStart bool) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.IsEpochStart = isEpochStart
	return nil
}

// SetNewEpochHeaderReceived sets the neeEpochHeaderReceived flag
func (strV2 *ShardTriggerRegistryV2) SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.NewEpochHeaderReceived = newEpochHeaderReceived
	return nil
}

// SetEpoch sets the epoch
func (strV2 *ShardTriggerRegistryV2) SetEpoch(epoch uint32) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.Epoch = epoch
	return nil
}

// SetMetaEpoch sets the metaChain epoch
func (strV2 *ShardTriggerRegistryV2) SetMetaEpoch(metaEpoch uint32) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.MetaEpoch = metaEpoch
	return nil
}

// SetCurrentRoundIndex sets the current round index
func (strV2 *ShardTriggerRegistryV2) SetCurrentRoundIndex(roundIndex int64) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.CurrentRoundIndex = roundIndex
	return nil
}

// SetEpochStartRound sets the epoch start round
func (strV2 *ShardTriggerRegistryV2) SetEpochStartRound(startRound uint64) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.EpochStartRound = startRound
	return nil
}

// SetEpochFinalityAttestingRound sets the epoch finality attesting round
func (strV2 *ShardTriggerRegistryV2) SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.EpochFinalityAttestingRound = finalityAttestingRound
	return nil
}

// SetEpochMetaBlockHash sets the epoch metaChain block hash
func (strV2 *ShardTriggerRegistryV2) SetEpochMetaBlockHash(epochMetaBlockHash []byte) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}
	strV2.EpochMetaBlockHash = epochMetaBlockHash
	return nil
}

// SetEpochStartHeaderHandler sets the epoch start header
func (strV2 *ShardTriggerRegistryV2) SetEpochStartHeaderHandler(epochStartHeaderHandler data.HeaderHandler) error {
	if strV2 == nil {
		return data.ErrNilPointerReceiver
	}

	var ok bool
	strV2.EpochStartShardHeader, ok = epochStartHeaderHandler.(*HeaderV2)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	return nil
}

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (strV3 *ShardTriggerRegistryV3) GetEpochStartHeaderHandler() data.HeaderHandler {
	if strV3 == nil {
		return nil
	}
	return strV3.GetEpochStartShardHeader()
}

// SetIsEpochStart sets the isEpochStart flag
func (strV3 *ShardTriggerRegistryV3) SetIsEpochStart(isEpochStart bool) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.IsEpochStart = isEpochStart
	return nil
}

// SetNewEpochHeaderReceived sets the neeEpochHeaderReceived flag
func (strV3 *ShardTriggerRegistryV3) SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.NewEpochHeaderReceived = newEpochHeaderReceived
	return nil
}

// SetEpoch sets the epoch
func (strV3 *ShardTriggerRegistryV3) SetEpoch(epoch uint32) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.Epoch = epoch
	return nil
}

// SetMetaEpoch sets the metaChain epoch
func (strV3 *ShardTriggerRegistryV3) SetMetaEpoch(metaEpoch uint32) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.MetaEpoch = metaEpoch
	return nil
}

// SetCurrentRoundIndex sets the current round index
func (strV3 *ShardTriggerRegistryV3) SetCurrentRoundIndex(roundIndex int64) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.CurrentRoundIndex = roundIndex
	return nil
}

// SetEpochStartRound sets the epoch start round
func (strV3 *ShardTriggerRegistryV3) SetEpochStartRound(startRound uint64) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.EpochStartRound = startRound
	return nil
}

// SetEpochFinalityAttestingRound sets the epoch finality attesting round
func (strV3 *ShardTriggerRegistryV3) SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.EpochFinalityAttestingRound = finalityAttestingRound
	return nil
}

// SetEpochMetaBlockHash sets the epoch metaChain block hash
func (strV3 *ShardTriggerRegistryV3) SetEpochMetaBlockHash(epochMetaBlockHash []byte) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}
	strV3.EpochMetaBlockHash = epochMetaBlockHash
	return nil
}

// SetEpochStartHeaderHandler sets the epoch start header
func (strV3 *ShardTriggerRegistryV3) SetEpochStartHeaderHandler(epochStartHeaderHandler data.HeaderHandler) error {
	if strV3 == nil {
		return data.ErrNilPointerReceiver
	}

	var ok bool
	strV3.EpochStartShardHeader, ok = epochStartHeaderHandler.(*HeaderV3)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	return nil
}

// GetEpochChangeProposed returns false for legacy meta registry
func (m *MetaTriggerRegistry) GetEpochChangeProposed() bool {
	return false
}

// GetEpochStartMetaHeaderHandler returns internal meta v1 header
func (m *MetaTriggerRegistry) GetEpochStartMetaHeaderHandler() data.MetaHeaderHandler {
	return m.GetEpochStartMeta()
}

// SetEpoch sets the epoch
func (m *MetaTriggerRegistry) SetEpoch(epoch uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.Epoch = epoch
	return nil
}

// SetEpochChangeProposed does nothing for legacy meta registry
func (m *MetaTriggerRegistry) SetEpochChangeProposed(_ bool) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	return nil
}

// SetCurrentRound sets the current round
func (m *MetaTriggerRegistry) SetCurrentRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.CurrentRound = round
	return nil
}

// SetEpochFinalityAttestingRound sets epoch finality attesting round
func (m *MetaTriggerRegistry) SetEpochFinalityAttestingRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.EpochFinalityAttestingRound = round
	return nil
}

// SetCurrEpochStartRound sets current epoch start round
func (m *MetaTriggerRegistry) SetCurrEpochStartRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.CurrEpochStartRound = round
	return nil
}

// SetPrevEpochStartRound sets previous epoch start round
func (m *MetaTriggerRegistry) SetPrevEpochStartRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.PrevEpochStartRound = round
	return nil
}

// SetEpochStartMetaHash sets epoch start meta hash
func (m *MetaTriggerRegistry) SetEpochStartMetaHash(hash []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.EpochStartMetaHash = hash
	return nil
}

// SetEpochStartMetaHeaderHandler sets internal meta block v1
func (m *MetaTriggerRegistry) SetEpochStartMetaHeaderHandler(header data.MetaHeaderHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	headerV1, castOk := header.(*MetaBlock)
	if !castOk {
		return fmt.Errorf("%w in MetaTriggerRegistry.SetEpochStartMetaHeaderHandler", data.ErrInvalidTypeAssertion)
	}

	m.EpochStartMeta = headerV1
	return nil
}

// GetEpochStartMetaHeaderHandler returns internal meta v3 header
func (m *MetaTriggerRegistryV3) GetEpochStartMetaHeaderHandler() data.MetaHeaderHandler {
	return m.GetEpochStartMeta()
}

// SetEpoch sets the epoch
func (m *MetaTriggerRegistryV3) SetEpoch(epoch uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.Epoch = epoch
	return nil
}

// SetEpochChangeProposed sets epoch change proposed bool flag
func (m *MetaTriggerRegistryV3) SetEpochChangeProposed(flag bool) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.EpochChangeProposed = flag
	return nil
}

// SetCurrentRound sets the current round
func (m *MetaTriggerRegistryV3) SetCurrentRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.CurrentRound = round
	return nil
}

// SetEpochFinalityAttestingRound sets epoch finality attesting round
func (m *MetaTriggerRegistryV3) SetEpochFinalityAttestingRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.EpochFinalityAttestingRound = round
	return nil
}

// SetCurrEpochStartRound sets current epoch start round
func (m *MetaTriggerRegistryV3) SetCurrEpochStartRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.CurrEpochStartRound = round
	return nil
}

// SetPrevEpochStartRound sets previous epoch start round
func (m *MetaTriggerRegistryV3) SetPrevEpochStartRound(round uint64) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.PrevEpochStartRound = round
	return nil
}

// SetEpochStartMetaHash sets epoch start meta hash
func (m *MetaTriggerRegistryV3) SetEpochStartMetaHash(hash []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}
	m.EpochStartMetaHash = hash
	return nil
}

// SetEpochStartMetaHeaderHandler sets internal meta block v3
func (m *MetaTriggerRegistryV3) SetEpochStartMetaHeaderHandler(header data.MetaHeaderHandler) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	headerV3, castOk := header.(*MetaBlockV3)
	if !castOk {
		return fmt.Errorf("%w in MetaTriggerRegistryV3.SetEpochStartMetaHeaderHandler", data.ErrInvalidTypeAssertion)
	}

	m.EpochStartMeta = headerV3
	return nil
}
