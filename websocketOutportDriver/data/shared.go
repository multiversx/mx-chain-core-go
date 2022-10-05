package data

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/outport"
)

// HeaderType defines the type to be used for the header that is sent
type HeaderType string

const (
	// MetaHeader defines the type of *block.MetaBlock
	MetaHeader HeaderType = "*block.MetaBlock"
	// ShardHeaderV1 defines the type of *block.Header
	ShardHeaderV1 HeaderType = "*block.Header"
	// ShardHeaderV2 defines the type of *block.HeaderV2
	ShardHeaderV2 HeaderType = "*block.HeaderV2"
)

// WsSendArgs holds the arguments needed for performing a web socket request
type WsSendArgs struct {
	Payload []byte
}

// ArgsRevertIndexedBlock holds the driver's arguments needed for reverting an indexed block
type ArgsRevertIndexedBlock struct {
	HeaderType HeaderType
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
	HeaderType HeaderType
	outport.ArgsSaveBlockData
}
