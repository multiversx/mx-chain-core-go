package core

import "github.com/mr-tron/base58/base58"

// PeerID is a p2p peer identity.
type PeerID string

// NewPeerID creates a new peer id or returns an error
func NewPeerID(input string) (PeerID, error) {
	pidBytes, err := base58.Decode(input)
	return PeerID(pidBytes), err
}

// Bytes returns the peer ID as byte slice
func (pid PeerID) Bytes() []byte {
	return []byte(pid)
}

// Pretty returns a b58-encoded string of the peer id
func (pid PeerID) Pretty() string {
	return base58.Encode(pid.Bytes())
}
