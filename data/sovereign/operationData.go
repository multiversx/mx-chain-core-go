package sovereign

import "math/big"

type Operation struct {
	Address      []byte
	Tokens       []EsdtToken
	TransferData *TransferData
}

type EsdtTokenPayment struct {
	TokenIdentifier []byte
	Nonce           uint64
	Amount          *big.Int
}

type EsdtToken struct {
	Identifier []byte
	Nonce      uint64
	Data       EsdtTokenData
}

type EsdtTokenType uint32

const (
	Fungible EsdtTokenType = iota
	NonFungible
	SemiFungible
	Meta
	Invalid
)

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
