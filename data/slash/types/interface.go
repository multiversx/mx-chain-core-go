package types

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
)

// SlashingProofHandler - contains a proof for a slashing event and can be wrapped in a transaction
type SlashingProofHandler interface {
	//GetType - contains the type of slashing detection
	GetType() slash.SlashingType
}

// MultipleProposalProofHandler contains proof data for a multiple header proposal slashing event
type MultipleProposalProofHandler interface {
	SlashingProofHandler
	// GetLevel - contains the slashing level for the current slashing type
	// multiple colluding parties should have a higher level
	GetLevel() slash.ThreatLevel
	//GetHeaders - returns the slashable proposed headers
	GetHeaders() []data.HeaderInfoHandler
}

// MultipleSigningProofHandler contains proof data for a multiple header signing slashing event
type MultipleSigningProofHandler interface {
	SlashingProofHandler
	// GetPubKeys - returns all validator's public keys which have signed multiple headers
	GetPubKeys() [][]byte
	// GetLevel - returns the slashing level for a given validator
	GetLevel(pubKey []byte) slash.ThreatLevel
	// GetHeaders - returns the slashable signed headers proposed by a given validator
	GetHeaders(pubKey []byte) []data.HeaderInfoHandler
}
