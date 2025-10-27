package block

// SetHeaderHash sets the header hash
func (sdp *ShardDataProposal) SetHeaderHash(headerHash []byte) {
	sdp.HeaderHash = headerHash
}

// SetRound sets the round
func (sdp *ShardDataProposal) SetRound(round uint64) {
	sdp.Round = round
}

// SetNonce sets the nonce
func (sdp *ShardDataProposal) SetNonce(nonce uint64) {
	sdp.Nonce = nonce
}

// SetShardID sets the shard ID
func (sdp *ShardDataProposal) SetShardID(shardID uint32) {
	sdp.ShardID = shardID
}

// SetEpoch sets the epoch
func (sdp *ShardDataProposal) SetEpoch(epoch uint32) {
	sdp.Epoch = epoch
}

// IsInterfaceNil returns true if there is no value under the interface
func (sdp *ShardDataProposal) IsInterfaceNil() bool {
	return sdp == nil
}
