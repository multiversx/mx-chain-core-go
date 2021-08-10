//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf  --gogoslick_out=. trigger.proto
package block

import "github.com/ElrondNetwork/elrond-go-core/data"

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (tr *TriggerRegistry) GetEpochStartHeaderHandler() data.HeaderHandler {
	if tr == nil {
		return nil
	}
	return tr.GetEpochStartShardHeader()
}

// SetIsEpochStart sets the isEpochStart flag
func (tr *TriggerRegistry) SetIsEpochStart(isEpochStart bool) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}

	tr.IsEpochStart = isEpochStart
	return nil
}

// SetNewEpochHeaderReceived sets the newEpochHeaderReceived flag
func (tr *TriggerRegistry) SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.NewEpochHeaderReceived = newEpochHeaderReceived
	return nil
}

// SetEpoch sets the epoch
func (tr *TriggerRegistry) SetEpoch(epoch uint32) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.Epoch = epoch
	return nil
}

// SetMetaEpoch Sets the metaChain epoch
func (tr *TriggerRegistry) SetMetaEpoch(metaEpoch uint32) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.MetaEpoch = metaEpoch
	return nil
}

// SetCurrentRoundIndex sets the current round index
func (tr *TriggerRegistry) SetCurrentRoundIndex(roundIndex int64) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.CurrentRoundIndex = roundIndex
	return nil
}

// SetEpochStartRound sets the epoch start round
func (tr *TriggerRegistry) SetEpochStartRound(startRound uint64) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.EpochStartRound = startRound
	return nil
}

// SetEpochFinalityAttestingRound sets the epoch finality attesting round
func (tr *TriggerRegistry) SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.EpochFinalityAttestingRound = finalityAttestingRound
	return nil
}

// SetEpochMetaBlockHash sets the epoch metaChain block hash
func (tr *TriggerRegistry) SetEpochMetaBlockHash(epochMetaBlockHash []byte) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}
	tr.EpochMetaBlockHash = epochMetaBlockHash
	return nil
}

// SetEpochStartHeaderHandler sets the epoch start header
func (tr *TriggerRegistry) SetEpochStartHeaderHandler(epochStartHeaderHandler data.HeaderHandler) error {
	if tr == nil {
		return data.ErrNilPointerReceiver
	}

	var ok bool
	tr.EpochStartShardHeader, ok = epochStartHeaderHandler.(*Header)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	return nil
}

// GetEpochStartHeaderHandler returns the epoch start headerHandler
func (trV2 *TriggerRegistryV2) GetEpochStartHeaderHandler() data.HeaderHandler {
	if trV2 == nil {
		return nil
	}
	return trV2.GetEpochStartShardHeader()
}

// SetIsEpochStart sets the isEpochStart flag
func (trV2 *TriggerRegistryV2) SetIsEpochStart(isEpochStart bool) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.IsEpochStart = isEpochStart
	return nil
}

// SetNewEpochHeaderReceived sets the neeEpochHeaderReceived flag
func (trV2 *TriggerRegistryV2) SetNewEpochHeaderReceived(newEpochHeaderReceived bool) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.NewEpochHeaderReceived = newEpochHeaderReceived
	return nil
}

// SetEpoch sets the epoch
func (trV2 *TriggerRegistryV2) SetEpoch(epoch uint32) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.Epoch = epoch
	return nil
}

// SetMetaEpoch sets the metaChain epoch
func (trV2 *TriggerRegistryV2) SetMetaEpoch(metaEpoch uint32) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.MetaEpoch = metaEpoch
	return nil
}

// SetCurrentRoundIndex sets the current round index
func (trV2 *TriggerRegistryV2) SetCurrentRoundIndex(roundIndex int64) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.CurrentRoundIndex = roundIndex
	return nil
}

// SetEpochStartRound sets the epoch start round
func (trV2 *TriggerRegistryV2) SetEpochStartRound(startRound uint64) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.EpochStartRound = startRound
	return nil
}

// SetEpochFinalityAttestingRound sets the epoch finality attesting round
func (trV2 *TriggerRegistryV2) SetEpochFinalityAttestingRound(finalityAttestingRound uint64) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.EpochFinalityAttestingRound = finalityAttestingRound
	return nil
}

// SetEpochMetaBlockHash sets the epoch metaChain block hash
func (trV2 *TriggerRegistryV2) SetEpochMetaBlockHash(epochMetaBlockHash []byte) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}
	trV2.EpochMetaBlockHash = epochMetaBlockHash
	return nil
}

// SetEpochStartHeaderHandler sets the epoch start header
func (trV2 *TriggerRegistryV2) SetEpochStartHeaderHandler(epochStartHeaderHandler data.HeaderHandler) error {
	if trV2 == nil {
		return data.ErrNilPointerReceiver
	}

	var ok bool
	trV2.EpochStartShardHeader, ok = epochStartHeaderHandler.(*HeaderV2)
	if !ok {
		return data.ErrInvalidTypeAssertion
	}
	return nil
}
