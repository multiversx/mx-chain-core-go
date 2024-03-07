package sovereign

import "math/big"

// Operation holds the data needed for a bridge transfer
type Operation struct {
	Address      []byte
	Tokens       []EsdtToken
	TransferData *TransferData
}

// EsdtToken holds the token data
type EsdtToken struct {
	Identifier []byte
	Nonce      uint64
	Data       EsdtTokenData
}

// EsdtTokenType holds the type of the esdt token
type EsdtTokenType uint32

const (
	Fungible EsdtTokenType = iota
	NonFungible
	SemiFungible
	Meta
	Invalid
)

// EsdtTokenData holds the properties for a token
type EsdtTokenData struct {
	TokenType  EsdtTokenType
	Amount     *big.Int
	Frozen     bool
	Hash       []byte
	Name       []byte
	Attributes []byte
	Creator    []byte
	Royalties  *big.Int
	Uris       [][]byte
}
