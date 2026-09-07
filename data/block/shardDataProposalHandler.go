package block

import "github.com/multiversx/mx-chain-core-go/data"

// SetHeaderHash sets the header hash
func (sdp *ShardDataProposal) SetHeaderHash(headerHash []byte) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.HeaderHash = headerHash

	return nil
}

// SetRound sets the round
func (sdp *ShardDataProposal) SetRound(round uint64) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.Round = round

	return nil
}

// SetNonce sets the nonce
func (sdp *ShardDataProposal) SetNonce(nonce uint64) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.Nonce = nonce

	return nil
}

// SetShardID sets the shard ID
func (sdp *ShardDataProposal) SetShardID(shardID uint32) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.ShardID = shardID

	return nil
}

// SetEpoch sets the epoch
func (sdp *ShardDataProposal) SetEpoch(epoch uint32) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.Epoch = epoch

	return nil
}

// SetNumPendingMiniBlocks sets the number of pending miniBlocks
func (sdp *ShardDataProposal) SetNumPendingMiniBlocks(num uint32) error {
	if sdp == nil {
		return data.ErrNilPointerReceiver
	}

	sdp.NumPendingMiniBlocks = num

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (sdp *ShardDataProposal) IsInterfaceNil() bool {
	return sdp == nil
}
