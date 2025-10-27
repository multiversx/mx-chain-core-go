package block

import "github.com/multiversx/mx-chain-core-go/data"

// EmptyBlockCreator is able to create empty block instances
type EmptyBlockCreator interface {
	CreateNewHeader() data.HeaderHandler
	IsInterfaceNil() bool
}

// ShardDataProposalHandler defines the behavior of a shard data proposal
type ShardDataProposalHandler interface {
	GetHeaderHash() []byte
	SetHeaderHash(headerHash []byte)
	GetRound() uint64
	SetRound(round uint64)
	GetNonce() uint64
	SetNonce(nonce uint64)
	GetShardID() uint32
	SetShardID(shardID uint32)
	GetEpoch() uint32
	SetEpoch(epoch uint32)
	IsInterfaceNil() bool
}
