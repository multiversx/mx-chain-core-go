package data

import (
	"github.com/multiversx/mx-chain-core-go/data/alteredAccount"
	"github.com/multiversx/mx-chain-core-go/data/outport"
)

// WsSendArgs holds the arguments needed for performing a web socket request
type WsSendArgs struct {
	Payload []byte
}

// ArgsSaveValidatorsRating holds the driver's arguments needed for indexing validators' rating
type ArgsSaveValidatorsRating struct {
	IndexID    string
	InfoRating []*outport.ValidatorRatingInfo
}

// ArgsSaveAccounts holds the driver's arguments needed for indexing accounts
type ArgsSaveAccounts struct {
	ShardID        uint32
	BlockTimestamp uint64
	Acc            map[string]*alteredAccount.AlteredAccount
}

// ArgsFinalizedBlock holds the driver's arguments needed for handling a finalized block
type ArgsFinalizedBlock struct {
	HeaderHash []byte
}
