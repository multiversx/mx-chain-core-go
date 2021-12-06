package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// SlashingProofHandler - contains a proof for a slashing event and can be wrapped in a transaction
type SlashingProofHandler interface {
	// GetProofTxData extracts proof tx data(see ProofTxData) from a slashing proof
	GetProofTxData() (*ProofTxData, error)
}

// MultipleProposalProofHandler contains proof data for a multiple header proposal slashing event
type MultipleProposalProofHandler interface {
	SlashingProofHandler
	// GetLevel - contains the slashing level for the current slashing type
	// multiple colluding parties should have a higher level
	GetLevel() ThreatLevel
	//GetHeaders - returns the slashable proposed headers
	GetHeaders() []data.HeaderHandler
}

// MultipleSigningProofHandler contains proof data for a multiple header signing slashing event
type MultipleSigningProofHandler interface {
	SlashingProofHandler
	// GetPubKeys - returns all validator's public keys which have signed multiple headers
	GetPubKeys() [][]byte
	// GetAllHeaders returns all signed headers
	GetAllHeaders() []data.HeaderHandler
	// GetLevel - returns the slashing level for a given validator
	GetLevel(pubKey []byte) ThreatLevel
	// GetHeaders - returns the slashable signed headers proposed by a given validator
	GetHeaders(pubKey []byte) []data.HeaderHandler
}
