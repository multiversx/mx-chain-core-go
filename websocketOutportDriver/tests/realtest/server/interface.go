package main

import (
	"github.com/multiversx/mx-chain-core-go/data/alteredAccount"
	"github.com/multiversx/mx-chain-core-go/data/outport"
)

// Driver is an interface for saving node specific data to other storage.
// This could be an elastic search index, a MySql database or any other external services.
type Driver interface {
	SaveBlock(outportBlock *outport.OutportBlock) error
	RevertIndexedBlock(headerBody *outport.HeaderBody) error
	SaveRoundsInfo(roundsInfos *outport.RoundsInfo) error
	SaveValidatorsPubKeys(validatorsPubKeys *outport.ValidatorsPubKeys) error
	SaveValidatorsRating(indexID string, infoRating []*outport.ValidatorRatingInfo) error
	SaveAccounts(blockTimestamp uint64, acc map[string]*alteredAccount.AlteredAccount, shardID uint32) error
	FinalizedBlock(headerHash []byte) error
	Close() error
	IsInterfaceNil() bool
}
