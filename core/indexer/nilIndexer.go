package indexer

import (
	"github.com/ElrondNetwork/elrond-go/core/statistics"
	"github.com/ElrondNetwork/elrond-go/data"
)

type NilIndexer struct {
}

// NewNilIndexer will return a Nil indexer
func NewNilIndexer() *NilIndexer {
	return new(NilIndexer)
}

// SaveBlock will do nothing
func (ni *NilIndexer) SaveBlock(body data.BodyHandler, header data.HeaderHandler, txPool map[string]data.TransactionHandler, signersIndexes []uint64) {
	return
}

// SaveRoundInfo will do nothing
func (ni *NilIndexer) SaveRoundInfo(round int64, shardId uint32, signersIndexes []uint64, blockWasProposed bool) {
	return
}

// UpdateTPS will do nothing
func (ni *NilIndexer) UpdateTPS(tpsBenchmark statistics.TPSBenchmark) {
	return
}

// SaveValidatorsPubKeys will do nothing
func (ni *NilIndexer) SaveValidatorsPubKeys(validatorsPubKeys map[uint32][][]byte) {
	return
}

// IsInterfaceNil returns true if there is no value under the interface
func (ni *NilIndexer) IsInterfaceNil() bool {
	if ni == nil {
		return true
	}
	return false
}

// IsNilIndexer return if implementation of indexer is a nil implementation
func (ni *NilIndexer) IsNilIndexer() bool {
	return true
}
