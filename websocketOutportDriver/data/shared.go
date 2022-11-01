package data

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/outport"
)

// WsSendArgs holds the arguments needed for performing a web socket request
type WsSendArgs struct {
	Payload []byte
}

// ArgsRevertIndexedBlock holds the driver's arguments needed for reverting an indexed block
type ArgsRevertIndexedBlock struct {
	HeaderType core.HeaderType
	Header     data.HeaderHandler
	Body       data.BodyHandler
}

// ArgsSaveRoundsInfo holds the driver's arguments needed for indexing rounds info
type ArgsSaveRoundsInfo struct {
	RoundsInfos []*outport.RoundInfo
}

// ArgsSaveValidatorsPubKeys holds the driver's arguments needed for indexing validator public keys
type ArgsSaveValidatorsPubKeys struct {
	ValidatorsPubKeys map[uint32][][]byte
	Epoch             uint32
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
	Acc            map[string]*outport.AlteredAccount
}

// ArgsFinalizedBlock holds the driver's arguments needed for handling a finalized block
type ArgsFinalizedBlock struct {
	HeaderHash []byte
}

// ArgsSaveBlock holds the driver's arguments needed for handling a save block
type ArgsSaveBlock struct {
	HeaderType core.HeaderType
	outport.ArgsSaveBlockData
}
