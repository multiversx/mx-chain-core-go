package main

import (
	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/multiversx/mx-chain-core-go/data/alteredAccount"
	"github.com/multiversx/mx-chain-core-go/data/outport"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(args *outport.OutportBlock) error
	RevertIndexedBlock(header data.HeaderHandler, body data.BodyHandler) error
	SaveRoundsInfo(roundsInfos []*outport.RoundInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys map[uint32][][]byte, epoch uint32) error
	SaveValidatorsRating(indexID string, infoRating []*outport.ValidatorRatingInfo) error
	SaveAccounts(blockTimestamp uint64, acc map[string]*alteredAccount.AlteredAccount, shardID uint32) error
	FinalizedBlock(headerHash []byte) error
	Close() error
	IsInterfaceNil() bool
}
