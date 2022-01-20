package slash

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// SlashingProofHandler contains a proof for a slashing event and can be wrapped in a transaction
type SlashingProofHandler interface {
	// GetProofTxData extracts proof tx data(see ProofTxData) from a slashing proof
	GetProofTxData() (*ProofTxData, error)
}

// MultipleProposalProofHandler contains proof data for a multiple header proposal slashing event
type MultipleProposalProofHandler interface {
	SlashingProofHandler
	// GetLevel returns the threat level of the proposer
	GetLevel() ThreatLevel
	//GetHeaders returns all proposed headers in a certain round
	GetHeaders() []data.HeaderHandler
}

// MultipleSigningProofHandler contains proof data for a multiple header signing slashing event
type MultipleSigningProofHandler interface {
	SlashingProofHandler
	// GetPubKeys returns all validator's public keys which have signed multiple headers
	GetPubKeys() [][]byte
	// GetAllHeaders returns all signed headers by all validators in a certain round
	GetAllHeaders() []data.HeaderHandler
	// GetLevel returns the threat level of a given validator
	GetLevel(pubKey []byte) ThreatLevel
	// GetHeaders returns the signed headers by a given validator
	GetHeaders(pubKey []byte) []data.HeaderHandler
}
