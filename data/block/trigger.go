//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. trigger.proto
package block

import "github.com/multiversx/mx-chain-core-go/data"

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
