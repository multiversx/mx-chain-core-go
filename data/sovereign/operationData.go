package sovereign

import (
	"math/big"

	"github.com/multiversx/mx-chain-core-go/core"
)

// Operation holds the data needed for an outgoing operation
type Operation struct {
	Address []byte
	Tokens  []EsdtToken
	Data    *EventData
}

// EsdtToken holds the token data
type EsdtToken struct {
	Identifier []byte
	Nonce      uint64
	Data       EsdtTokenData
}

// EsdtTokenData holds the properties for a token
type EsdtTokenData struct {
	TokenType  core.ESDTType
	Amount     *big.Int
	Frozen     bool
	Hash       []byte
	Name       []byte
	Attributes []byte
	Creator    []byte
	Royalties  *big.Int
	Uris       [][]byte
}
