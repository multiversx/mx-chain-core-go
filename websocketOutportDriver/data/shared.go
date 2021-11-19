package data

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/indexer"
)

// WsSendArgs holds the arguments needed for performing a web socket request
type WsSendArgs struct {
	Payload []byte
	Route   string
}

// ArgsRevertIndexedBlock holds the driver's arguments needed for reverting an indexed block
type ArgsRevertIndexedBlock struct {
	Header data.HeaderHandler
	Body   data.BodyHandler
}

// ArgsSaveRoundsInfo holds the driver's arguments needed for indexing rounds info
type ArgsSaveRoundsInfo struct {
	RoundsInfos []*indexer.RoundInfo
}

// ArgsSaveValidatorsPubKeys holds the driver's arguments needed for indexing validator public keys
type ArgsSaveValidatorsPubKeys struct {
	ValidatorsPubKeys map[uint32][][]byte
	Epoch             uint32
}

// ArgsSaveValidatorsRating holds the driver's arguments needed for indexing validators' rating
type ArgsSaveValidatorsRating struct {
	IndexID    string
	InfoRating []*indexer.ValidatorRatingInfo
}

// ArgsSaveAccounts holds the driver's arguments needed for indexing accounts
type ArgsSaveAccounts struct {
	BlockTimestamp uint64
	Acc            []data.UserAccountHandler
}

// ArgsFinalizedBlock holds the driver's arguments needed for handling a finalized block
type ArgsFinalizedBlock struct {
	HeaderHash []byte
}
